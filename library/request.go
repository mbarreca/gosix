package library

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/mbarreca/gosix/client"
)

// APISix Error Object
type APISixError struct {
	ErrorMsg    string `json:"error_msg"`
	Description string `json:"description"`
}

// REQUEST
// body - any - This can be any object that will go in the request's body
// headers - []gosix.library.Header - This will be an array of the type header and will include all the headers you want to send
// endpoint - string - This will contain the endpoint address (we use an environment variable to set the hostname)
// requestType - string - This will contain the type of request you're making
// client - gosix.Client - This will contain the context and HTTP Client we need to make the request
// RETURN
// []byte - A byte array with the response
// error - An error if one has occured
func DoRequest(body any, headers []Header, endpoint, requestType string, client *client.Client) ([]byte, error) {
	// Create the request object
	var req *http.Request
	var err error
	if body != nil {
		// Check to see if this is an SSL request, if so then we need to remove the escape new lines
		buf := new(bytes.Buffer)
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		// This will prevent JSON marhsalling from breaking the SSL certificate and key
		if reflect.TypeOf(body).String() == "models.SSL" {
			bodyJSON = bytes.Replace(bodyJSON, []byte(`\\n`), []byte(`\n`), -1)
		}
		buf = bytes.NewBuffer(bodyJSON)
		// Form the request
		req, err = http.NewRequest(requestType, os.Getenv("GOSIX_APISIX_ADDRESS")+endpoint, buf)
		if err != nil {
			return nil, err
		}
	} else {
		// Otherwise set it to nil and create the request
		req, err = http.NewRequest(requestType, os.Getenv("GOSIX_APISIX_ADDRESS")+endpoint, nil)
		if err != nil {
			return nil, err
		}
	}
	// Loop through and apply headers to the request
	for _, header := range headers {
		req.Header.Set(header.Key, header.Value)
	}

	// Add Ubiqitious Headers - JSON Type Definition and API Key
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", os.Getenv("GOSIX_APISIX_ADMIN_API_KEY"))

	// Add Context
	req.WithContext(client.Ctx)
	// Form and make the request
	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	// Read the body, check for errors
	responseBody, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		// This means there's been an error, marshal in the typical APISix format
		var r APISixError
		err = json.Unmarshal([]byte(responseBody), &r)
		if err != nil {
			return nil, err
		}

		// If it returns something unexpected, provide a catch all just in case
		if r.ErrorMsg == "" {
			return nil, errors.New(string(responseBody))
		}
		return nil, errors.New(r.ErrorMsg)
	}
	// Close the body
	resp.Body.Close()
	// Return
	return responseBody, nil
}

func RESTCall[T any, V any](url, method string, obj T, client *client.Client) (V, error) {
	var empty V
	response, err := DoRequest(obj, nil, url, method, client)
	if err != nil {
		return empty, err
	}
	// Unmarshal
	var r V
	err = json.Unmarshal([]byte(response), &r)
	if err != nil {
		return empty, err
	}
	// Validate Response
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(r)
	if err != nil {
		fmt.Println("R: ", r)
		return empty, err
	}
	return r, nil
}

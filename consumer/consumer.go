package consumer

import (
	"encoding/json"

	"github.com/go-playground/validator/v10" // MIT
	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer/models"
	"github.com/mbarreca/gosix/library"
)

// Fetches a list of all Consumers.
func Get(client *gosix.Client) (models.GetResponse, error) {
	response, err := library.DoRequest(nil, nil, "/apisix/admin/consumers", "GET", client)
	if err != nil {
		return models.GetResponse{}, err
	}
	// Unmarshal
	var r models.GetResponse
	err = json.Unmarshal([]byte(response), &r)
	if err != nil {
		return models.GetResponse{}, err
	}
	// Validate Response
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(r)
	if err != nil {
		return models.GetResponse{}, err
	}
	return r, nil
}

// Fetches specified Consumer by username.
func GetByUsername(username string, client *gosix.Client) (models.GetByUsernameResponse, error) {
	response, err := library.DoRequest(nil, nil, "/apisix/admin/consumers/"+username, "GET", client)
	if err != nil {
		return models.GetByUsernameResponse{}, err
	}
	// Unmarshal
	var r models.GetByUsernameResponse
	err = json.Unmarshal([]byte(response), &r)
	if err != nil {
		return models.GetByUsernameResponse{}, err
	}
	// Validate Response
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(r)
	if err != nil {
		return models.GetByUsernameResponse{}, err
	}
	return r, nil
}

// Create new Consumer. Ensure to put initial auth in here upon creation
func Put(consumer models.ConsumerRequest, client *gosix.Client) (models.PutResponse, error) {
	response, err := library.DoRequest(consumer, nil, "/apisix/admin/consumers/", "PUT", client)
	if err != nil {
		return models.PutResponse{}, err
	}
	// Unmarshal
	var r models.PutResponse
	err = json.Unmarshal([]byte(response), &r)
	if err != nil {
		return models.PutResponse{}, err
	}
	// Validate Response
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(r)
	if err != nil {
		return models.PutResponse{}, err
	}
	return r, nil
}

// Removes the Consumer with the specified username.
func Delete(username string, client *gosix.Client) (models.DeleteResponse, error) {
	response, err := library.DoRequest(nil, nil, "/apisix/admin/consumers/"+username, "DELETE", client)
	if err != nil {
		return models.DeleteResponse{}, err
	}
	// Unmarshal
	var r models.DeleteResponse
	err = json.Unmarshal([]byte(response), &r)
	if err != nil {
		return models.DeleteResponse{}, err
	}
	// Validate Response
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(r)
	if err != nil {
		return models.DeleteResponse{}, err
	}
	return r, nil
}

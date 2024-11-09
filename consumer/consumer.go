package consumer

import (
	"encoding/json"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/library"
)

// Fetches a list of all Consumers.
func Get(client *gosix.Client) (GetResponse, error) {
	response, err := library.DoRequest(nil, nil, "/apisix/admin/consumers", "GET", client)
	if err != nil {
		return GetResponse{}, err
	}
	// Unmarshal
	var r GetResponse
	err = json.Unmarshal([]byte(response), &r)
	if err != nil {
		return GetResponse{}, err
	}
	return r, nil
}

// Fetches specified Consumer by username.
func GetByUsername(username string, client *gosix.Client) (GetByUsernameResponse, error) {
	response, err := library.DoRequest(nil, nil, "/apisix/admin/consumers/"+username, "GET", client)
	if err != nil {
		return GetByUsernameResponse{}, err
	}
	// Unmarshal
	var r GetByUsernameResponse
	err = json.Unmarshal([]byte(response), &r)
	if err != nil {
		return GetByUsernameResponse{}, err
	}
	return r, nil
}

// Create new Consumer.
func Put(consumer ConsumerRequest, client *gosix.Client) (PutResponse, error) {
	response, err := library.DoRequest(consumer, nil, "/apisix/admin/consumers/", "PUT", client)
	if err != nil {
		return PutResponse{}, err
	}
	// Unmarshal
	var r PutResponse
	err = json.Unmarshal([]byte(response), &r)
	if err != nil {
		return PutResponse{}, err
	}
	return r, nil
}

// Removes the Consumer with the specified username.
func Delete(username string, client *gosix.Client) (DeleteResponse, error) {
	response, err := library.DoRequest(nil, nil, "/apisix/admin/consumers/"+username, "DELETE", client)
	if err != nil {
		return DeleteResponse{}, err
	}
	// Unmarshal
	var r DeleteResponse
	err = json.Unmarshal([]byte(response), &r)
	if err != nil {
		return DeleteResponse{}, err
	}
	return r, nil
}

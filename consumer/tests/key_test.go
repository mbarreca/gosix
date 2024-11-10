package tests

import (
	"testing"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer"
	"github.com/mbarreca/gosix/consumer/models"
)

/*
Testing Functions
*/

// TestGet tests the get method
func TestKeyAuth(t *testing.T) {
	// Define the request object
	var c models.ConsumerRequest
	c.Username = "DELETETHISISATESTUSER"
	c.Desc = "DELETETHISISATESTUSERDESCRIPTION"

	client, err := gosix.New()
	if err != nil {
		t.Fatalf("Error in Client Creation: %v", err)
	}
	// Create a Consumer
	put, err := consumer.Put(c, client)
	if err != nil {
		t.Fatalf("Error in PUT Request: %v", err)
	}
	// Check the key and the username
	if put.Key != ("/apisix/consumers/" + c.Username) {
		t.Fatalf("Failed PUT Request Assertion: Key Field: %v", err)
	}
	if put.Value.Username != c.Username {
		t.Fatalf("Failed PUT Request Assertion: Username Field: %v", err)
	}
	// Add the Key to the User
	key, err := consumer.KeyAuthAdd(c.Username, client)
	if err != nil {
		t.Fatalf("Error in Key Auth add: %v", err)
	}
	// Check to see if the Key is present
	err = checkKey(c.Username, key, client)
	if err != nil {
		t.Fatalf("Error in Key Check: %v", err)
	}
	// Cycle the Key
	key, err = consumer.KeyAuthCycle(c.Username, client)
	if err != nil {
		t.Fatalf("Error in Key Cycle: %v", err)
	}
	// Check to see if the Key is present
	err = checkKey(c.Username, key, client)
	if err != nil {
		t.Fatalf("Error in Key Check: %v", err)
	}
	// Disable the Key
	err = consumer.KeyAuthEnabled(false, c.Username, client)
	if err != nil {
		t.Fatalf("Error in Key Cycle: %v", err)
	}
	status, err := getKeyStatus(c.Username, client)
	if err != nil {
		t.Fatalf("Error in Key Status Get Request: %v", err)
	}
	if status {
		t.Fatalf("Key Disable Failed")
	}
	// Enable the Key
	err = consumer.KeyAuthEnabled(true, c.Username, client)
	if err != nil {
		t.Fatalf("Error in Key Cycle: %v", err)
	}
	status, err = getKeyStatus(c.Username, client)
	if err != nil {
		t.Fatalf("Error in Key Status Get Request: %v", err)
	}
	if !status {
		t.Fatalf("Key Disable Failed")
	}
	// Delete the Username
	delete, err := consumer.Delete(c.Username, client)
	if err != nil {
		t.Fatalf("Error in DELETE By Username Request: %v", err)
	}
	// Check the key and the username
	if delete.Key != ("/apisix/consumers/" + c.Username) {
		t.Fatalf("Failed DELETE By Username Request Assertion: Key Field: %v", err)
	}
	// Check to see if the key is still there
	_, err = consumer.GetByUsername(c.Username, client)
	if err == nil {
		t.Fatalf("Error in DELETE By Username Request, Found After Request: %v", err)
	}
}

func checkKey(username, key string, client *gosix.Client) error {
	// Check to see if the Key is present
	getUser, err := consumer.GetByUsername(username, client)
	if err != nil {
		return err
	}
	if getUser.Key != ("/apisix/consumers/" + username) {
		return err
	}
	if getUser.Value.Username != username {
		return err
	}
	if getUser.Value.Plugins.KeyAuth.Key != key {
		return err
	}
	return nil
}
func getKeyStatus(username string, client *gosix.Client) (bool, error) {
	// Check to see if the Key is present
	getUser, err := consumer.GetByUsername(username, client)
	if err != nil {
		return false, err
	}
	return !getUser.Value.Plugins.KeyAuth.Meta.Disable, nil
}

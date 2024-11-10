package tests

import (
	"errors"
	"testing"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer"
	"github.com/mbarreca/gosix/consumer/models"
	"github.com/mbarreca/gosix/consumer/tests/lib"
)

/*
Testing Functions
*/

// TestGet tests the get method
func TestKeyAuth(t *testing.T) {
	// Create Consumer Request
	var c models.ConsumerRequest
	c.Username = "GosixTestUsername"
	c.Desc = "GosixTestDescription"
	// Create Client
	client, err := gosix.New()
	if err != nil {
		t.Fatalf("Error in Client Creation: %v", err)
	}
	// Add Consumer
	if err := lib.AddConsumer(c, client); err != nil {
		t.Fatalf("Error in Add Consumer: %v", err)
	}
	// Add the Key to the User
	key, err := consumer.KeyAuthAdd(c.Username, client)
	if err != nil {
		t.Fatalf("Error in Key Auth add: %v", err)
	}
	// Check to see if the Key is present
	if err := checkKey(c.Username, key, client); err != nil {
		t.Fatalf("Error in Key Check: %v", err)
	}
	// Cycle the Key
	key, err = consumer.KeyAuthCycle(c.Username, client)
	if err != nil {
		t.Fatalf("Error in Key Cycle: %v", err)
	}
	// Check to see if the Key is present and matches
	err = checkKey(c.Username, key, client)
	if err != nil {
		t.Fatalf("Error in Key Check: %v", err)
	}
	// Disable the Key
	if err := changeKeyStatus(false, c.Username, client); err != nil {
		t.Fatalf("Error in Key Disable: %v", err)
	}
	// Enable the Key
	if err := changeKeyStatus(true, c.Username, client); err != nil {
		t.Fatalf("Error in Key Enable: %v", err)
	}
	// Delete the Username
	if err := lib.DeleteConsumer(c.Username, client); err != nil {
		t.Fatalf("Error in Delete Consumer: %v", err)
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
func changeKeyStatus(status bool, username string, client *gosix.Client) error {
	// Change status
	if err := consumer.KeyAuthEnabled(status, username, client); err != nil {
		return err
	}
	// Get User
	getUser, err := consumer.GetByUsername(username, client)
	if err != nil {
		return err
	}
	// See if the status matches it
	if status != !getUser.Value.Plugins.KeyAuth.Meta.Disable {
		return errors.New("Key Status Mismatch")
	}
	return nil
}

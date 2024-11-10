package tests

import (
	"errors"
	"testing"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer"
	"github.com/mbarreca/gosix/consumer/models"
	"github.com/mbarreca/gosix/test/lib"
)

/*
Testing Functions
*/

// TestGet tests the get method
func TestBasicAuth(t *testing.T) {
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
	// Add the Basic to the User
	basicAuth := new(models.BasicAuth)
	basicAuth.Username = "testUser@testUserGosix.com"
	basicAuth.Password = "testPassword"
	if err := consumer.BasicAuthAdd(c.Username, basicAuth, client); err != nil {
		t.Fatalf("Error in Basic Auth add: %v", err)
	}
	// Check to see if the basic user has been added
	if err := checkBasic(c.Username, basicAuth, client); err != nil {
		t.Fatalf("Error in Basic Check: %v", err)
	}
	// Change the password
	basicAuth.Password = "testNewPassword"
	if err := consumer.BasicAuthChangePassword(c.Username, basicAuth.Password, client); err != nil {
		t.Fatalf("Error in Basic Change Password: %v", err)
	}
	// Check to see if the basic user has been added
	if err := checkBasic(c.Username, basicAuth, client); err != nil {
		t.Fatalf("Error in Basic Check: %v", err)
	}

	// Disable Basic Auth
	if err := changeBasicStatus(false, c.Username, client); err != nil {
		t.Fatalf("Error in Basic Disable: %v", err)
	}
	// Enable Basic Auth
	if err := changeBasicStatus(true, c.Username, client); err != nil {
		t.Fatalf("Error in Basic Enable: %v", err)
	}
	// Delete the User
	if err := lib.DeleteConsumer(c.Username, client); err != nil {
		t.Fatalf("Error in Delete Consumer: %v", err)
	}
}

// Checks if the Basic Auth User/Password matches APISix
func checkBasic(username string, basicAuth *models.BasicAuth, client *gosix.Client) error {
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
	if getUser.Value.Plugins.BasicAuth.Username != basicAuth.Username && getUser.Value.Plugins.BasicAuth.Password != basicAuth.Password {
		return errors.New("Basic Auth Username or Password does not match")
	}
	return nil
}

// Disables/Enables the plugin and checks APISix
func changeBasicStatus(status bool, username string, client *gosix.Client) error {
	// Change status
	if err := consumer.AuthEnabled(models.BasicAuth{}, status, username, client); err != nil {
		return err
	}
	// Get User
	getUser, err := consumer.GetByUsername(username, client)
	if err != nil {
		return err
	}
	// See if the status matches it
	if status != !getUser.Value.Plugins.BasicAuth.Meta.Disable {
		return errors.New("Key Status Mismatch")
	}
	return nil
}

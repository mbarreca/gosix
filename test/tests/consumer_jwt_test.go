package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer"
	"github.com/mbarreca/gosix/consumer/models"
	"github.com/mbarreca/gosix/test/lib"
)

/*
Testing Functions
*/

// TestGet tests the get method
func TestJwtAuth(t *testing.T) {
	// Create Consumer Request
	var c models.ConsumerRequest
	c.Username = "GosixTestUsername"
	c.Desc = "GosixTestDescription"
	key := "GosixTestKeyMinLength"
	// Create Client
	client, err := gosix.New()
	if err != nil {
		t.Fatalf("Error in Client Creation: %v", err)
	}
	// Add Consumer
	if err := lib.AddConsumer(c, client); err != nil {
		t.Fatalf("Error in Add Consumer: %v", err)
	}
	// Add the JWT Parameters to the User
	token, err := consumer.JWTAuthGetKey(c.Username, key, client)
	if err != nil {
		t.Fatalf("Error in Jwt Auth add: %v", err)
	}
	// Check to see if the Key is present
	if err := checkJwtKey(c.Username, key, client); err != nil {
		t.Fatalf("Error in Jwt Token Check: %v", err)
	}
	time.Sleep(time.Second * 2)
	// Cycle the Key
	newToken, err := consumer.JWTAuthGetKey(c.Username, key, client)
	if err != nil {
		t.Fatalf("Error in Jwt Token Cycle: %v", err)
	}
	if token == newToken {
		t.Fatalf("Error cycling Jwt Token")
	}
	// Disable the JwtToken
	if err := changeJwtStatus(false, c.Username, client); err != nil {
		t.Fatalf("Error in Jwt Token Disable: %v", err)
	}
	// Enable the JwtToken
	if err := changeJwtStatus(true, c.Username, client); err != nil {
		t.Fatalf("Error in Jwt Token Enable: %v", err)
	}
	// Delete the User
	if err := lib.DeleteConsumer(c.Username, client); err != nil {
		t.Fatalf("Error in Delete Consumer: %v", err)
	}
}

func checkJwtKey(username, key string, client *gosix.Client) error {
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
	if getUser.Value.Plugins.JwtAuth.Key != key {
		return err
	}
	return nil
}
func changeJwtStatus(status bool, username string, client *gosix.Client) error {
	// Change status
	if err := consumer.AuthEnabled(models.JwtAuth{}, status, username, client); err != nil {
		return err
	}
	// Get User
	getUser, err := consumer.GetByUsername(username, client)
	if err != nil {
		return err
	}
	// See if the status matches it
	if status != !getUser.Value.Plugins.JwtAuth.Meta.Disable {
		return errors.New("Key Status Mismatch")
	}
	return nil
}

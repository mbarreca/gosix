package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/test/lib"
)

/*
Testing Functions
*/

// Test JWT
func TestJwtAuth(t *testing.T) {
	key := "GosixTestKeyMinLength"
	// Create Consumer
	a, consumer, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	// Add the JWT Parameters to the User
	token, err := a.JWT.Get(consumer.Username, key)
	if err != nil {
		t.Fatal(err)
	}
	// Check to see if the Key is present
	if err := checkJwtKey(consumer.Username, key, a); err != nil {
		t.Fatalf("Error in Jwt Token Check: %v", err)
	}
	// Since its timestamp based we need to slow Go down
	time.Sleep(time.Second * 2)
	// Cycle the Key
	newToken, err := a.JWT.Get(consumer.Username, key)
	if err != nil {
		t.Fatal(err)
	}
	if token == newToken {
		t.Fatalf("Error cycling Jwt Token")
	}
	// Disable the JwtToken
	if err := changeJwtStatus(false, consumer.Username, a); err != nil {
		t.Fatalf("Error in Jwt Token Disable: %v", err)
	}
	// Enable the JwtToken
	if err := changeJwtStatus(true, consumer.Username, a); err != nil {
		t.Fatalf("Error in Jwt Token Enable: %v", err)
	}
	// Delete the JwtToken
	if err := a.JWT.Delete(consumer.Username); err != nil {
		t.Fatal(err)
	}
	// Delete the User
	if err := a.Consumer.Delete(consumer.Username); err != nil {
		t.Fatal(err)
	}
}

func checkJwtKey(username, key string, a *gosix.APISix) error {
	// Check to see if the Key is present
	consumer, err := a.Consumer.Get(username)
	if err != nil {
		return err
	}
	if consumer.Plugins.JwtAuth.Key != key {
		return errors.New("JWT Key check failed")
	}
	return nil
}
func changeJwtStatus(status bool, username string, a *gosix.APISix) error {
	// Change status
	if err := a.JWT.Enabled(status, username); err != nil {
		return err
	}
	consumer, err := a.Consumer.Get(username)
	if err != nil {
		return err
	}
	if consumer.Plugins == nil || consumer.Plugins.JwtAuth == nil || consumer.Plugins.JwtAuth.Meta == nil || consumer.Plugins.JwtAuth.Meta.Disable == status {
		return errors.New("JWT Auth Status Failure")
	}
	return nil
}

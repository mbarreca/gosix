package tests

import (
	"errors"
	"testing"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/test/lib"
)

/*
Testing Functions
*/

// Test Key
func TestKeyAuth(t *testing.T) {
	// Create Consumer
	a, consumer, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	// Add the Key Parameters to the User
	token, err := a.Key.Get(consumer.Username)
	if err != nil {
		t.Fatal(err)
	}
	// Check to see if the Key is present
	if err := checkKey(consumer.Username, token, a); err != nil {
		t.Fatalf("Error in Jwt Token Check: %v", err)
	}
	// Cycle the Key
	newToken, err := a.Key.Get(consumer.Username)
	if err != nil {
		t.Fatal(err)
	}
	if token == newToken {
		t.Fatalf("Error cycling Key Token")
	}
	// Disable
	if err := changeKeyStatus(false, consumer.Username, a); err != nil {
		t.Fatalf("Error in Key Token Disable: %v", err)
	}
	// Enable
	if err := changeKeyStatus(true, consumer.Username, a); err != nil {
		t.Fatalf("Error in Key Token Enable: %v", err)
	}
	// Delete the Key Token
	if err := a.Key.Delete(consumer.Username); err != nil {
		t.Fatal(err)
	}
	// Delete the User
	if err := a.Consumer.Delete(consumer.Username); err != nil {
		t.Fatal(err)
	}
}

func checkKey(username, key string, a *gosix.APISix) error {
	// Check to see if the Key is present
	consumer, err := a.Consumer.Get(username)
	if err != nil {
		return err
	}
	if consumer.Plugins.KeyAuth.Key != key {
		return errors.New("Key check failed")
	}
	return nil
}
func changeKeyStatus(status bool, username string, a *gosix.APISix) error {
	// Change status
	if err := a.Key.Enabled(status, username); err != nil {
		return err
	}
	consumer, err := a.Consumer.Get(username)
	if err != nil {
		return err
	}
	if consumer.Plugins == nil || consumer.Plugins.KeyAuth == nil || consumer.Plugins.KeyAuth.Meta == nil || consumer.Plugins.KeyAuth.Meta.Disable == status {
		return errors.New("Key Auth Status Failure")
	}
	return nil
}

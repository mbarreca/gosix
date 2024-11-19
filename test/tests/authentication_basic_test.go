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

// Test Basic Auth
func TestBasicAuth(t *testing.T) {
	username := "GoSixUsername"
	password := "GoSixTestPassword"
	// Create Consumer
	a, consumer, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	a.Basic.Add(consumer.Username, username, password)
	consumer, err = a.Consumer.Get(consumer.Username)
	if err != nil {
		t.Fatal(err)
	}
	if consumer.Plugins == nil || consumer.Plugins.BasicAuth == nil || consumer.Plugins.BasicAuth.Password != password {
		t.Fatal("Basic Auth Failed Add")
	}
	// Change the password
	password = "GoSixTestPasswordNew"
	a.Basic.ChangePassword(consumer.Username, password)
	if err != nil {
		t.Fatal(err)
	}
	consumer, err = a.Consumer.Get(consumer.Username)
	if err != nil {
		t.Fatal(err)
	}
	if consumer.Plugins == nil || consumer.Plugins.BasicAuth == nil || consumer.Plugins.BasicAuth.Password != password {
		t.Fatal("Basic Auth Failed Password Change")
	}
	// Disable Basic Auth
	if err := changeBasicStatus(false, consumer.Username, a); err != nil {
		t.Fatal(err)
	}
	// Enable Basic Auth
	if err := changeBasicStatus(true, consumer.Username, a); err != nil {
		t.Fatal(err)
	}
	// Delete the Basic Auth
	if err := a.Basic.Delete(consumer.Username); err != nil {
		t.Fatal(err)
	}
	// Delete the User
	if err := a.Consumer.Delete(consumer.Username); err != nil {
		t.Fatal(err)
	}
}

// Disables/Enables the plugin and checks APISix
func changeBasicStatus(status bool, username string, a *gosix.APISix) error {
	// Change status
	if err := a.Basic.Enabled(status, username); err != nil {
		return err
	}
	consumer, err := a.Consumer.Get(username)
	if err != nil {
		return err
	}
	if consumer.Plugins == nil || consumer.Plugins.BasicAuth == nil || consumer.Plugins.BasicAuth.Meta == nil || consumer.Plugins.BasicAuth.Meta.Disable == status {
		return errors.New("Basic Auth Status Failure")
	}
	return nil
}

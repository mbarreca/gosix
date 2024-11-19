package tests

import (
	"testing"

	"github.com/mbarreca/gosix/models"
	"github.com/mbarreca/gosix/test/lib"
)

/*
Testing Functions
*/

// TestGet tests the get method
func TestConsumer(t *testing.T) {
	a, consumer, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	// Check to see if it exists in all users
	users, err := a.Consumer.GetAll()
	found := false
	for _, user := range users {
		if user.Username == consumer.Username {
			found = true
		}
	}
	if !found {
		t.Fatal("Error finding Consumer from Get All")
	}
	if err := a.Consumer.Delete(consumer.Username); err != nil {
		t.Fatal(err)
	}
}

func TestConsumerPlugins(t *testing.T) {
	a, consumer, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	ba := new(models.BasicAuth)
	ba.Username = "A"
	ba.Password = "B"
	if err := a.Consumer.AddPlugin(consumer.Username, ba); err != nil {
		t.Fatal(err)
	}
	baG, err := a.Consumer.GetPlugin(consumer.Username, models.BasicAuth{})
	if err != nil {
		t.Fatal(err)
	}
	baGet := baG.(*models.BasicAuth)
	if baGet.Username != ba.Username || baGet.Password != ba.Password {
		t.Fatal("Add/Get Plugin Failed")
	}
	// Cleanup
	if err := a.Consumer.DeletePlugin(consumer.Username, models.BasicAuth{}); err != nil {
		t.Fatal(err)
	}
}

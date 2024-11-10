package tests

import (
	"testing"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer/models"
	"github.com/mbarreca/gosix/test/lib"
)

/*
Testing Functions
*/

// TestGet tests the get method
func TestConsumer(t *testing.T) {
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
	// Get Consumer
	if err := lib.GetConsumer(c.Username, client); err != nil {
		t.Fatalf("Error in Get Consumer: %v", err)
	}
	// Get Consumer from All
	if err := lib.GetConsumerFromAll(c.Username, client); err != nil {
		t.Fatalf("Error in Get Consumer From All: %v", err)
	}
	// Delete the Username
	if err := lib.DeleteConsumer(c.Username, client); err != nil {
		t.Fatalf("Error in Delete Consumer: %v", err)
	}
}

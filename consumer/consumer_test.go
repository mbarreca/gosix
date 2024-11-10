package consumer

import (
	"testing"

	"github.com/mbarreca/gosix"
)

/*
Testing Functions
*/

// TestGet tests the get method
func TestConsumer(t *testing.T) {
	// Define the request object
	var c ConsumerRequest
	c.Username = "DELETETHISISATESTUSER"
	c.Desc = "DELETETHISISATESTUSERDESCRIPTION"

	client, err := gosix.New()
	if err != nil {
		t.Fatalf("Error in Client Creation: %v", err)
	}
	// Make Create Request (PUT)
	put, err := Put(c, client)
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
	// Check Get By Username (GET by Username)
	getUser, err := GetByUsername(c.Username, client)
	if err != nil {
		t.Fatalf("Error in GET By Username Request: %v", err)
	}
	// Check the key and the username
	if getUser.Key != ("/apisix/consumers/" + c.Username) {
		t.Fatalf("Failed GET By Username Request Assertion: Key Field: %v", err)
	}
	if getUser.Value.Username != c.Username {
		t.Fatalf("Failed GET By Username Request Assertion: Username Field: %v", err)
	}
	// Check Get Request
	getUsers, err := Get(client)
	if err != nil {
		t.Fatalf("Error in GET: %v", err)
	}
	// Find the Key
	found := false
	for _, user := range getUsers.List {
		if (user.Key == ("/apisix/consumers/" + c.Username)) && (user.Value.Username == c.Username) {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Failed GET Request Assertion: Username or Key Field: %v", err)
	}
	// Delete the Username
	delete, err := Delete(c.Username, client)
	if err != nil {
		t.Fatalf("Error in DELETE By Username Request: %v", err)
	}
	// Check the key and the username
	if delete.Key != ("/apisix/consumers/" + c.Username) {
		t.Fatalf("Failed DELETE By Username Request Assertion: Key Field: %v", err)
	}
	// Check to see if the key is still there
	_, err = GetByUsername(c.Username, client)
	if err == nil {
		t.Fatalf("Error in DELETE By Username Request, Found After Request: %v", err)
	}
}

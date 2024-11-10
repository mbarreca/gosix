package tests

import (
	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer"
	"github.com/mbarreca/gosix/consumer/models"
)

func AddConsumer(c models.ConsumerRequest, client *gosix.Client) error {
	// Make Create Request (PUT)
	put, err := consumer.Put(c, client)
	if err != nil {
		return err
	}
	// Check the key and the username
	if put.Key != ("/apisix/consumers/" + c.Username) {
		return err
	}
	if put.Value.Username != c.Username {
		return err
	}
	return nil
}
func GetConsumer(username string, client *gosix.Client) error {
	// Check Get By Username (GET by Username)
	getUser, err := consumer.GetByUsername(username, client)
	if err != nil {
		return err
	}
	// Check the key and the username
	if getUser.Key != ("/apisix/consumers/" + username) {
		return err
	}
	if getUser.Value.Username != username {
		return err
	}
	return nil
}
func GetConsumerFromAll(username string, client *gosix.Client) error {
	// Check Get Request
	getUsers, err := consumer.Get(client)
	if err != nil {
		return err
	}
	// Find the Key
	found := false
	for _, user := range *getUsers.List {
		if (user.Key == ("/apisix/consumers/" + username)) && (user.Value.Username == username) {
			found = true
			break
		}
	}
	if !found {
		return err
	}
	return nil
}
func DeleteConsumer(username string, client *gosix.Client) error {
	// Delete the Username
	delete, err := consumer.Delete(username, client)
	if err != nil {
		return err
	}
	// Check the key and the username
	if delete.Key != ("/apisix/consumers/" + username) {
		return err
	}
	// Check to see if the key is still there
	_, err = consumer.GetByUsername(username, client)
	if err == nil {
		return err
	}
	return nil
}

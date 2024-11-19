package authentication

import (
	"errors"
	"os"
	"strconv"

	"github.com/mbarreca/gosix/client"
	"github.com/mbarreca/gosix/consumer"
	"github.com/mbarreca/gosix/library"
	"github.com/mbarreca/gosix/models"
)

type Key struct {
	client *client.Client
	c      *consumer.Consumer
}

// Constructor - You *shouldn't* be using this
func NewKey(client *client.Client, consumer *consumer.Consumer) *Key {
	return &Key{client: client, c: consumer}
}

// Add key authentication to the selected consumer, we will auto-generate a key - 100 characters long
// If Key auth exists, this will cycle the key and return it to you
// username -> Consumers username
func (k *Key) Get(username string) (string, error) {
	user, err := k.c.Get(username)
	if err != nil {
		return "", err
	}
	if user.Plugins != nil && user.Plugins.JwtAuth != nil {
		return "", errors.New("You can't have JWT and Key Auth on the same consumer")
	}
	// Get an new key object with a new key
	keyAuth, err := createKeyObject()
	if err != nil {
		return "", err
	}
	if user.Plugins == nil {
		user.Plugins = new(models.Plugins)
	}
	// Modify the consumer
	user.Plugins.KeyAuth = keyAuth
	if err := k.c.Update(user); err != nil {
		return "", err
	}
	return keyAuth.Key, nil
}

// Delete the Key Plugin from the Consumer
// username -> Consumers username
func (k *Key) Delete(username string) error {
	user, err := k.c.Get(username)
	if err != nil {
		return err
	}
	if user.Plugins == nil || user.Plugins.KeyAuth == nil {
		return errors.New("No Key Auth on this consumer")
	}
	user.Plugins.KeyAuth = nil
	if err := k.c.Update(user); err != nil {
		return err
	}
	return nil
}

// Set the enabled/disabled sate of Basic Auth for this consumer
// enabled - True if enabled, false if disabled - default state is Enabled
// username -> Consumers username

func (k *Key) Enabled(enabled bool, username string) error {
	// Get the consumer
	user, err := k.c.Get(username)
	if err != nil {
		return err
	}
	if user.Plugins != nil && user.Plugins.KeyAuth != nil {
		// Check to see if Meta Exists
		if user.Plugins.KeyAuth.Meta == nil {
			user.Plugins.KeyAuth.Meta = new(models.Meta)
		}
		// Disable the key
		user.Plugins.KeyAuth.Meta.Disable = !enabled
	} else {
		return errors.New("User doesn't have a Key Auth plugin")
	}
	// Update the consumer
	if err := k.c.Update(user); err != nil {
		return err
	}
	return nil
}

func createKeyObject() (*models.KeyAuth, error) {
	length := 100
	var err error
	// Check if we've overriden the default key length
	if os.Getenv("GOSIX_APISIX_PLUGIN_KEY_LENGTH") != "" {
		length, err = strconv.Atoi(os.Getenv("GOSIX_APISIX_PLUGIN_KEY_LENGTH"))
		if err != nil {
			return nil, err
		}
	}
	// Build the Key Auth Object, make a new key and update
	var keyAuth models.KeyAuth
	keyAuth.Key = library.RandomString(length)
	return &keyAuth, nil
}

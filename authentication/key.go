package authentication

import (
	"errors"
	"os"
	"strconv"
	"time"

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

// Add key authentication to the selected consumer, we will auto-generate a key - 100 characters long
// If Key auth exists, this will cycle the key and return it to you
// username -> Consumers username
// exp -> The key's expiration offset in seconds -> Note that this is not a field built into APISIX so you will need to verify it yourself
func (k *Key) GetWithExp(username string, exp int) (string, error) {
	user, err := k.c.Get(username)
	if err != nil {
		return "", err
	}
	// Get an new key object with a new key
	keyAuth, err := createKeyObject()
	if err != nil {
		return "", err
	}
	if user.Plugins == nil {
		user.Plugins = new(models.Plugins)
	}
	keyAuth.Exp = time.Now().Add(time.Second * time.Duration(exp)).UTC().Format("01-02-2006 15:04:05.000000")
	// Modify the consumer
	user.Plugins.KeyAuth = keyAuth
	if err := k.c.Update(user); err != nil {
		return "", err
	}
	return keyAuth.Key, nil
}

// Validate will compare the key provided plus the expiry time and validate it
// username -> Consumers username
// key -> The key provided to validate
func (k *Key) Validate(username, key string) (bool, error) {
	user, err := k.c.Get(username)
	if err != nil {
		return false, err
	}
	if user.Plugins == nil || user.Plugins.KeyAuth == nil {
		return false, errors.New("Key Auth doesn't exist")
	}
	keyAuth := user.Plugins.KeyAuth
	if keyAuth.Exp != "" {
		t, err := time.Parse("01-02-2006 15:04:05.000000", keyAuth.Exp)
		if err != nil {
			return false, errors.New("Issue with parsing Expiry")
		}
		if time.Now().UTC().Unix() > t.Unix() {
			// Key is expired
			return false, nil
		}
	}
	if keyAuth.Key != key {
		return false, nil
	}
	return true, nil
}

// Validate will only validate if a key is expired
// username -> Consumers username
func (k *Key) ValidateExp(username string) (bool, error) {
	user, err := k.c.Get(username)
	if err != nil {
		return false, err
	}
	if user.Plugins == nil || user.Plugins.KeyAuth == nil {
		return false, errors.New("Key Auth doesn't exist")
	}
	keyAuth := user.Plugins.KeyAuth
	if keyAuth.Exp != "" {
		t, err := time.Parse("01-02-2006 15:04:05.000000", keyAuth.Exp)
		if err != nil {
			return false, errors.New("Issue with parsing Expiry")
		}
		if time.Now().UTC().Unix() > t.Unix() {
			// Key is expired
			return false, nil
		}
	}
	return true, nil
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
	keyAuth.Key = library.RandomStringAlphaNum(length)
	return &keyAuth, nil
}

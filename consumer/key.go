package consumer

import (
	"errors"
	"os"
	"strconv"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer/models"
	"github.com/mbarreca/gosix/library"
)

// Add key authentication to the selected consumer, we will auto-generate a key - 100 characters long
// If Key auth exists, this will cycle the key and return it to you
func KeyAuthGetKey(username string, client *gosix.Client) (string, error) {
	// Get the consumer
	origConsumer, err := GetByUsername(username, client)
	if err != nil {
		return "", err
	}
	// Pull relevant fields
	var modConsumer models.ConsumerRequest
	// This is safe because Value needs to be included otherwise the validator will throw an error
	plugins := origConsumer.Value.Plugins
	// You can't have Key and JWT Auth in APISix, prevent this
	if plugins != nil && plugins.JwtAuth != nil {
		return "", errors.New("You can't have JWT and Key Auth on the same consumer")
	}
	// Get an new key object with a new key
	keyAuth, err := createKeyObject()
	if err != nil {
		return "", err
	}
	// If there are no plugins, add them
	if plugins == nil {
		plugins = new(models.Plugins)
	}
	// Re-create the modified consumer
	plugins.KeyAuth = keyAuth
	modConsumer.Username = origConsumer.Value.Username
	modConsumer.Desc = origConsumer.Value.Desc
	modConsumer.Plugins = plugins
	_, err = Put(modConsumer, client)
	if err != nil {
		return "", err
	}
	return keyAuth.Key, nil
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

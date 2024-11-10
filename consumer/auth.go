package consumer

import (
	"errors"
	"reflect"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer/models"
)

// This package will use reflection to retreive the auth type and enable/disable
func AuthEnabled(authObj any, enabled bool, username string, client *gosix.Client) error {
	// Get the consumer
	origConsumer, err := GetByUsername(username, client)
	if err != nil {
		return err
	}
	// Pull relevant fields
	var modConsumer models.ConsumerRequest
	plugins := origConsumer.Value.Plugins
	// This is safe because Value needs to be included otherwise the validator will throw an error
	if plugins == nil {
		return errors.New("There is no auth on this consumer")
	}
	// Find the data type
	authType := reflect.TypeOf(authObj)
	if authType.Name() == "JwtAuth" {
		if plugins.JwtAuth == nil {
			return errors.New("Jwt Auth not added to the consumer")
		}
		if plugins.JwtAuth.Meta == nil {
			plugins.JwtAuth.Meta = new(models.Meta)
		}
		// Change the key state
		plugins.JwtAuth.Meta.Disable = !enabled

	} else if authType.Name() == "BasicAuth" {
		if plugins.BasicAuth == nil {
			return errors.New("Basic Auth not added to the consumer")
		}
		if plugins.BasicAuth.Meta == nil {
			plugins.BasicAuth.Meta = new(models.Meta)
		}
		// Change the key state
		plugins.BasicAuth.Meta.Disable = !enabled

	} else if authType.Name() == "KeyAuth" {
		if plugins.KeyAuth == nil {
			return errors.New("Key Auth not added to the consumer")
		}
		if plugins.KeyAuth.Meta == nil {
			plugins.KeyAuth.Meta = new(models.Meta)
		}
		// Change the key state
		plugins.KeyAuth.Meta.Disable = !enabled
	} else {
		return errors.New("Unsupported or Invalid Type Sent")
	}
	modConsumer.Username = origConsumer.Value.Username
	modConsumer.Desc = origConsumer.Value.Desc
	modConsumer.Plugins = plugins
	_, err = Put(modConsumer, client)
	if err != nil {
		return err
	}
	return nil
}

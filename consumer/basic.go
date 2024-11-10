package consumer

import (
	"errors"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer/models"
)

// Add the Basic Auth Authentication Method to a consumer
func BasicAuthAdd(username string, basicAuth *models.BasicAuth, client *gosix.Client) error {
	// Get the consumer
	origConsumer, err := GetByUsername(username, client)
	if err != nil {
		return err
	}
	// Pull relevant fields
	var modConsumer models.ConsumerRequest
	// This is safe because Value needs to be included otherwise the validator will throw an error
	plugins := origConsumer.Value.Plugins
	if plugins != nil && plugins.BasicAuth != nil {
		// This means Basic Auth is already added, error out
		return errors.New("Basic Auth Already added to consumer")
	}
	// If there are no plugins, add them
	if plugins == nil {
		plugins = new(models.Plugins)
	}
	// Re-create the modified consumer
	plugins.BasicAuth = basicAuth
	modConsumer.Username = origConsumer.Value.Username
	modConsumer.Desc = origConsumer.Value.Desc
	modConsumer.Plugins = plugins
	_, err = Put(modConsumer, client)
	if err != nil {
		return err
	}
	return nil
}

// Reset the users basic auth password to a random string and return it
func BasicAuthChangePassword(username, password string, client *gosix.Client) error {
	// Get the consumer
	origConsumer, err := GetByUsername(username, client)
	if err != nil {
		return err
	}
	// Pull relevant fields
	var modConsumer models.ConsumerRequest
	// This is safe because Value needs to be included otherwise the validator will throw an error
	plugins := origConsumer.Value.Plugins
	if plugins == nil || plugins.BasicAuth == nil {
		// This means Basic Auth is already added, error out
		return errors.New("Basic Auth Already added to consumer")
	}
	// Update the password and re-create the modified consumer
	plugins.BasicAuth.Password = password
	modConsumer.Username = origConsumer.Value.Username
	modConsumer.Desc = origConsumer.Value.Desc
	modConsumer.Plugins = plugins
	_, err = Put(modConsumer, client)
	if err != nil {
		return err
	}
	return nil
}

// Disable Basic Auth for this consumer
func BasicAuthEnabled(enabled bool, username string, client *gosix.Client) error {
	// Get the consumer
	origConsumer, err := GetByUsername(username, client)
	if err != nil {
		return err
	}
	// Pull relevant fields
	var modConsumer models.ConsumerRequest
	plugins := origConsumer.Value.Plugins
	// This is safe because Value needs to be included otherwise the validator will throw an error
	if plugins == nil || plugins.BasicAuth == nil {
		// This means Key Auth is already added, error out
		return errors.New("Basic Auth not added to the consumer")
	}
	// Check to see if Meta Exists
	if plugins.BasicAuth.Meta == nil {
		plugins.BasicAuth.Meta = new(models.Meta)
	}
	// Disable the key
	plugins.BasicAuth.Meta.Disable = !enabled
	modConsumer.Username = origConsumer.Value.Username
	modConsumer.Desc = origConsumer.Value.Desc
	modConsumer.Plugins = plugins
	_, err = Put(modConsumer, client)
	if err != nil {
		return err
	}
	return nil
}

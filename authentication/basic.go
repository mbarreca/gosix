package authentication

import (
	"errors"

	"github.com/mbarreca/gosix/client"
	"github.com/mbarreca/gosix/consumer"
	"github.com/mbarreca/gosix/models"
)

type Basic struct {
	client *client.Client
	c      *consumer.Consumer
}

// Constructor - You *shouldn't* be using this
func NewBasic(client *client.Client, consumer *consumer.Consumer) *Basic {
	return &Basic{client: client, c: consumer}
}

// Add will either add basic auth to a consumer (if it doesn't exist) or update the password if it does
// consumerUsername -> The username of *consumer* you'd like to add to
// basicUsername -> The username you'd like the user to login with
// basicPassword -> The password you'd like the user to have
func (b *Basic) Add(consumerUsername, basicUsername, basicPassword string) error {
	user, err := b.c.Get(consumerUsername)
	if err != nil {
		return err
	}
	if user.Plugins == nil {
		user.Plugins = new(models.Plugins)
		user.Plugins.BasicAuth = new(models.BasicAuth)
	} else if user.Plugins.BasicAuth == nil {
		user.Plugins.BasicAuth = new(models.BasicAuth)
	}
	user.Plugins.BasicAuth.Username = basicUsername
	user.Plugins.BasicAuth.Password = basicPassword
	if err := b.c.Update(user); err != nil {
		return err
	}
	return nil
}

// Reset the users basic auth password to a random string and return it
// consumerUsername -> The username of *consumer* you'd like to access
// basicPassword -> The password you'd like the user to have
func (b *Basic) ChangePassword(consumerUsername, basicPassword string) error {
	// Get the consumer
	user, err := b.c.Get(consumerUsername)
	if err != nil {
		return err
	}
	if user.Plugins != nil && user.Plugins.BasicAuth != nil {
		user.Plugins.BasicAuth.Password = basicPassword
	} else {
		return errors.New("User doesn't have a Basic Auth plugin")
	}
	// Update the consumer
	if err := b.c.Update(user); err != nil {
		return err
	}
	return nil
}

// Delete the entire Basic Plugin from the Consumer using the consumers username
// consumerUsername -> The username of *consumer* you'd like to delete
func (b *Basic) Delete(consumerUsername string) error {
	user, err := b.c.Get(consumerUsername)
	if err != nil {
		return err
	}
	if user.Plugins == nil || user.Plugins.BasicAuth == nil {
		return errors.New("No Key Auth on this consumer")
	}
	user.Plugins.BasicAuth = nil
	if err := b.c.Update(user); err != nil {
		return err
	}
	return nil
}

// Set the enabled/disabled sate of Basic Auth for this consumer
// enabled - True if enabled, false if disabled - default state is Enabled
// consumerUsername -> The username of *consumer* you'd like to change
func (b *Basic) Enabled(enabled bool, consumerUsername string) error {
	// Get the consumer
	user, err := b.c.Get(consumerUsername)
	if err != nil {
		return err
	}
	if user.Plugins != nil && user.Plugins.BasicAuth != nil {
		// Check to see if Meta Exists
		if user.Plugins.BasicAuth.Meta == nil {
			user.Plugins.BasicAuth.Meta = new(models.Meta)
		}
		// Disable the key
		user.Plugins.BasicAuth.Meta.Disable = !enabled
	} else {
		return errors.New("User doesn't have a Basic Auth plugin")
	}
	// Update the consumer
	if err := b.c.Update(user); err != nil {
		return err
	}
	return nil
}

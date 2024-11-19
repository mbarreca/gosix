package consumer

import (
	"errors"

	"github.com/mbarreca/gosix/api"
	"github.com/mbarreca/gosix/client"
	"github.com/mbarreca/gosix/library"
	"github.com/mbarreca/gosix/models"
)

type Consumer struct {
	client *client.Client
}

// Constructor - You *shouldn't* be using this
func New(c *client.Client) *Consumer {
	return &Consumer{client: c}
}

// This will create a consumer using the data within the consumer object
// username -> The username you want the user to have
// description -> OPTIONAL - a description for the consumer
// plugins -> OPTIONAL - a plugins object to attach to the user
func (c *Consumer) Create(username, description string, plugins *models.Plugins) error {
	var user models.Consumer
	if username == "" {
		return errors.New("Must provide a username to create a consumer")
	}
	user.Username = username
	user.Desc = description
	user.Plugins = plugins
	resp, err := api.API[models.Consumer, models.Create]("", "", user, c.client)
	if err != nil {
		return err
	}
	u := resp.(models.Object[models.Consumer])
	if u.Value.Username != username {
		return errors.New("Error creating the consumer, usernames don't match, there might be something wrong with your APISIX Instance")
	}
	return nil
}

// This will see if the user in the consumer.User object exists, if so it will set and return the object, if not you'll get an error
// username -> The username of the user to get
func (c *Consumer) Get(username string) (models.Consumer, error) {
	resp, err := api.API[models.Consumer, models.Get]("", username, models.Consumer{}, c.client)
	if err != nil {
		return models.Consumer{}, err
	}
	u := resp.(models.Object[models.Consumer])
	if u.Value.Username != username {
		return models.Consumer{}, errors.New("The user doesn't exist")
	}
	return *u.Value, nil
}

// This will get all consumers
func (c *Consumer) GetAll() ([]models.Consumer, error) {
	resp, err := api.API[models.Consumer, models.Get]("", "", models.Consumer{}, c.client)
	if err != nil {
		return nil, err
	}
	var r []models.Consumer
	users := resp.(models.All[models.Consumer])
	for _, user := range *users.Objects {
		r = append(r, *user.Value)
	}
	return r, nil
}

// This will Update a consumer
// consumer -> The consumer object to update
func (c *Consumer) Update(consumer models.Consumer) error {
	// Strip out create and update time, APISix can't handle it
	var user models.Consumer
	user.Username = consumer.Username
	user.Desc = consumer.Desc
	user.Plugins = consumer.Plugins
	if _, err := api.API[models.Consumer, models.Create]("", "", user, c.client); err != nil {
		return err
	}
	return nil
}

// This will delete a consumer
// username -> The username of the user you want to delete
func (c *Consumer) Delete(username string) error {
	if _, err := api.API[models.Consumer, models.Delete]("", username, models.Consumer{}, c.client); err != nil {
		return err
	}
	return nil
}

/*
	Plugin Specific Functions
*/

// This will get the plugin attached to a consumer and write it into plugin
// username -> The username of the consumer
// plugin -> The plugin object (used only for type)
func (c *Consumer) GetPlugin(username string, plugin any) (any, error) {
	user, err := c.Get(username)
	if err != nil {
		return nil, err
	}
	return library.GetPlugin(user.Plugins, plugin)
}

// This will add the plugin if it doesn't exist or update it if it does
// username -> The username of the consumer
// plugin -> The plugin object
func (c *Consumer) AddPlugin(username string, plugin any) error {
	user, err := c.Get(username)
	if err != nil {
		return err
	}
	p, err := library.AddPlugin(user.Plugins, plugin)
	if err != nil {
		return err
	}
	user.Plugins = p
	if err := c.Update(user); err != nil {
		return err
	}
	return nil
}

// This will delete the plugin if it exists, if it doesn't it will still return no error
// username -> The username of the consumer
// plugin -> The plugin object (used only for type)
func (c *Consumer) DeletePlugin(username string, plugin any) error {
	user, err := c.Get(username)
	if err != nil {
		return err
	}
	p, err := library.DeletePlugin(user.Plugins, plugin)
	if err != nil {
		return err
	}
	user.Plugins = p
	if err := c.Update(user); err != nil {
		return err
	}
	return nil
}

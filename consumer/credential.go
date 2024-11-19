package consumer

import (
	"github.com/mbarreca/gosix/api"
	"github.com/mbarreca/gosix/client"
	"github.com/mbarreca/gosix/models"
)

type Credential struct {
	client *client.Client
	c      *Consumer
}

// Constructor - You *shouldn't* be using this
func NewCredential(c *client.Client, consumer *Consumer) *Credential {
	return &Credential{client: c, c: consumer}
}

// Gets all credentials attached to a username
// username -> The username to get credentials from
// id -> A specific credential to get
func (c *Credential) Get(username, id string) (models.Credential, error) {
	resp, err := api.API[models.Credential, models.Get](id, username, models.Credential{}, c.client)
	if err != nil {
		return models.Credential{}, err
	}
	return *resp.(models.Object[models.Credential]).Value, nil
}

// Get all credentials for a username
// username -> The username to get credentials from
func (c *Credential) GetAll(username string) ([]models.Credential, error) {
	resp, err := api.API[models.Credential, models.Get]("", username, models.Credential{}, c.client)
	if err != nil {
		return nil, err
	}
	var ret []models.Credential
	credentials := resp.(models.All[models.Credential])
	for _, credential := range *credentials.Objects {
		ret = append(ret, *credential.Value)
	}
	return ret, nil
}

// Create a credential
// id -> A specific Service to get
// username -> The username of the credential
// credential -> Credential Object
func (c *Credential) Create(username, id string, credential models.Credential) error {
	_, err := api.API[models.Credential, models.Create](id, username, credential, c.client)
	if err != nil {
		return err
	}
	return nil
}

// Delete a Credential
// id -> A specific Service to get
// username -> The username of the credential
func (c *Credential) Delete(username, id string) error {
	_, err := api.API[models.Credential, models.Delete](id, username, models.Credential{}, c.client)
	if err != nil {
		return err
	}
	return nil
}

package consumergroup

import (
	"github.com/mbarreca/gosix/api"
	"github.com/mbarreca/gosix/client"
	"github.com/mbarreca/gosix/consumer"
	"github.com/mbarreca/gosix/models"
)

type ConsumerGroup struct {
	client *client.Client
	c      *consumer.Consumer
}

// Constructor - You *shouldn't* be using this
func New(c *client.Client, consumer *consumer.Consumer) *ConsumerGroup {
	return &ConsumerGroup{client: c, c: consumer}
}

// Gets the ConsumerGroups attached to an ID
// id -> A specific ConsumerGroup to get
func (c *ConsumerGroup) Get(id string) (models.ConsumerGroup, error) {
	resp, err := api.API[models.ConsumerGroup, models.Get](id, "", models.ConsumerGroup{}, c.client)
	if err != nil {
		return models.ConsumerGroup{}, err
	}
	return *resp.(models.Object[models.ConsumerGroup]).Value, nil
}

// Gets all ConsumerGroups
func (c *ConsumerGroup) GetAll() ([]models.ConsumerGroup, error) {
	resp, err := api.API[models.ConsumerGroup, models.Get]("", "", models.ConsumerGroup{}, c.client)
	if err != nil {
		return nil, err
	}
	var ret []models.ConsumerGroup
	vals := resp.(models.All[models.ConsumerGroup])
	for _, val := range *vals.Objects {
		ret = append(ret, *val.Value)
	}
	return ret, nil
}

// Create an ConsumerGroup, leave the ID blank for an autogenerated ID
// id -> ID of the ConsumerGroup
// value -> ConsumerGroup Object
func (c *ConsumerGroup) Create(id string, value models.ConsumerGroup) error {
	_, err := api.API[models.ConsumerGroup, models.Create](id, "", value, c.client)
	if err != nil {
		return err
	}
	return nil
}

// Update a ConsumerGroup
// id -> ID of the ConsumerGroup
// value -> ConsumerGroup Object
func (c *ConsumerGroup) Update(id string, value models.ConsumerGroup) error {
	_, err := api.API[models.ConsumerGroup, models.Update](id, "", value, c.client)
	if err != nil {
		return err
	}
	return nil
}

// Delete a ConsumerGroup
// id -> ID of the ConsumerGroup
func (c *ConsumerGroup) Delete(id string) error {
	_, err := api.API[models.ConsumerGroup, models.Delete](id, "", models.ConsumerGroup{}, c.client)
	if err != nil {
		return err
	}
	return nil
}

package lib

import (
	"context"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/models"
)

func CreateAndGetConsumer() (*gosix.APISix, models.Consumer, error) {
	username := "GosixTestUsername"
	// Create APISix Client
	a, err := gosix.New(context.Background(), false)
	if err != nil {
		return nil, models.Consumer{}, err
	}
	// Add Consumer
	if err := a.Consumer.Create(username, "GosixTestDescription", nil); err != nil {
		return nil, models.Consumer{}, err
	}
	// Check to see it was created properly
	user, err := a.Consumer.Get(username)
	if err != nil {
		return nil, models.Consumer{}, err
	}
	return a, user, nil
}

package lib

import (
	"errors"
	"reflect"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/models"
)

// This function will do all the common functions for the basic API groups
// Create, Get, GetAll, Update, Delete
// Returns the original created consumer and the updated one for you to compare
func Test[A APIs[V], V models.Value](id string, value V, updateValue V, api A, a *gosix.APISix) (V, V, error) {
	var e V
	vType := reflect.TypeOf(updateValue).String()
	// Create item
	if err := api.Create(id, value); err != nil {
		return e, e, err
	}
	// Get item
	item, err := api.Get(id)
	if err != nil {
		return e, e, err
	}
	if item.GetID() != id {
		return e, e, errors.New("GET: ID doesn't match")
	}
	// Get all items
	items, err := api.GetAll()
	if err != nil {
		return e, e, err
	}
	found := false
	for _, valueItems := range items {
		if valueItems.GetID() == id {
			found = true
			break
		}
	}
	if !found {
		return e, e, errors.New("GetAll: Can't find ID")

	}
	// Skip update for SSL since its not supported
	if vType != "models.SSL" {
		// Update item
		if err := api.Update(id, updateValue); err != nil {
			return e, e, err
		}
	}
	// Get item
	itemUpdated, err := api.Get(id)
	if err != nil {
		return e, e, err
	}
	// Delete item
	if err := api.Delete(id); err != nil {
		return e, e, err
	}
	// Done
	return item, itemUpdated, nil
}

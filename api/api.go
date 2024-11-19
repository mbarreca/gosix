package api

import (
	"errors"
	"reflect"

	"github.com/mbarreca/gosix/client"
	"github.com/mbarreca/gosix/library"
	"github.com/mbarreca/gosix/models"
)

// This is the generic API Call that breaks into the following operations
// Get -> If you provide the requisite ID/Username combo it will return one object, otherwise it will get all
// Create -> This will create a new entry, if it exists, this will perform an update
// Update -> This will update an entry, if it doesn't exist it will fail
// Delete -> This will delete an entry
func API[Value models.Value, Kind models.Kind](id, username string, value Value, client *client.Client) (any, error) {
	var k Kind
	var v Value
	var err error
	url := ""
	method := ""
	apiValue := reflect.TypeOf(v).String()
	apiKind := reflect.TypeOf(k).String()
	// Set the URLs and do any last minute checks
	if apiValue == "models.Consumer" {
		method, url, err = validateConsumer(username, apiKind, value)
	} else if apiValue == "models.ConsumerGroup" {
		method, url, err = validateConsumerGroup(id, apiKind, value)
	} else if apiValue == "models.Credential" {
		method, url, err = validateCredential(id, username, apiKind, value)
	} else if apiValue == "models.GlobalRule" {
		method, url, err = validateGlobalRule(id, apiKind, value)
	} else if apiValue == "models.PluginConfig" {
		method, url, err = validatePluginConfig(id, apiKind, value)
	} else if apiValue == "models.Proto" {
		method, url, err = validateProto(id, apiKind, value)
	} else if apiValue == "models.Route" {
		method, url, err = validateRoute(id, apiKind, value)
	} else if apiValue == "models.Secret" {
		method, url, err = validateSecret(id, apiKind, value)
	} else if apiValue == "models.Service" {
		method, url, err = validateService(id, apiKind, value)
	} else if apiValue == "models.SSL" {
		method, url, err = validateSSL(id, apiKind, value)
	} else if apiValue == "models.Stream" {
		method, url, err = validateStream(id, apiKind, value)
	} else if apiValue == "models.Upstream" {
		method, url, err = validateUpstream(id, apiKind, value)
	} else {
		err = errors.New("Somehow a type got passed that wasn't registered")
	}
	// Check to see if we have an error
	if err != nil {
		return nil, err
	}
	// Run call
	if apiKind == "models.Get" && (id == "" && username == "") {
		return library.RESTCall[Value, models.All[Value]](url, method, value, client)
	} else if apiKind == "models.Delete" {
		return library.RESTCall[Value, models.DeleteResponse](url, method, value, client)
	}
	return library.RESTCall[Value, models.Object[Value]](url, method, value, client)
}

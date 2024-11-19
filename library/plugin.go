package library

import (
	"errors"
	"reflect"

	"github.com/mbarreca/gosix/models"
)

// Universal Add Plugin Function - Consumer/Route/Service
func AddPlugin(userPlugins *models.Plugins, plugin any) (*models.Plugins, error) {
	if userPlugins == nil {
		userPlugins = new(models.Plugins)
	}
	pluginType := reflect.TypeOf(plugin).String()
	if pluginType == "*models.BasicAuth" {
		p := plugin.(*models.BasicAuth)
		userPlugins.BasicAuth = p
	} else if pluginType == "*models.KeyAuth" {
		p := plugin.(*models.KeyAuth)
		userPlugins.KeyAuth = p
	} else if pluginType == "*models.JwtAuth" {
		p := plugin.(*models.JwtAuth)
		userPlugins.JwtAuth = p
	} else if pluginType == "*models.ConsumerRestriction" {
		p := plugin.(*models.ConsumerRestriction)
		userPlugins.ConsumerRestriction = p
	} else {
		return nil, errors.New("No match for this plugin type found")
	}
	return userPlugins, nil
}

// Universal Get Plugin Function - Consumer/Route/Service
func GetPlugin(userPlugins *models.Plugins, plugin any) (any, error) {
	if userPlugins == nil {
		return nil, errors.New("No plugins on this consumer")
	}
	pluginType := reflect.TypeOf(plugin).String()
	if pluginType == "models.BasicAuth" {
		return getPlugin(userPlugins.BasicAuth)
	} else if pluginType == "models.KeyAuth" {
		return getPlugin(userPlugins.KeyAuth)
	} else if pluginType == "models.JwtAuth" {
		return getPlugin(userPlugins.JwtAuth)
	} else if pluginType == "models.ConsumerRestriction" {
		return getPlugin(userPlugins.ConsumerRestriction)
	}
	return nil, errors.New("No match for this plugin type found")

}

// Internal Get Plugin Function
func getPlugin(a any) (any, error) {
	if a == nil {
		return nil, errors.New("No Restriction on this consumer")
	}
	return a, nil
}

// Universal Delete Plugin Function - Consumer/Route/Service
func DeletePlugin(userPlugins *models.Plugins, plugin any) (*models.Plugins, error) {
	if userPlugins == nil {
		return nil, nil
	}
	pluginType := reflect.TypeOf(plugin).String()
	if pluginType == "models.BasicAuth" {
		userPlugins.BasicAuth = nil
	} else if pluginType == "models.KeyAuth" {
		userPlugins.KeyAuth = nil
	} else if pluginType == "models.JwtAuth" {
		userPlugins.JwtAuth = nil
	} else if pluginType == "models.ConsumerRestriction" {
		userPlugins.ConsumerRestriction = nil
	} else {
		return nil, errors.New("No match for this plugin type found")
	}
	return userPlugins, nil
}

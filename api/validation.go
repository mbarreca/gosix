package api

import (
	"errors"
	"strings"

	"github.com/mbarreca/gosix/models"
)

func validateConsumer[Value models.Value](username, kind string, value Value) (string, string, error) {
	// 	Get | Create | Update | Delete
	// Get the operation type
	if kind == "models.Get" {
		if username == "" {
			return "GET", "/apisix/admin/consumers", nil
		}
		return "GET", "/apisix/admin/consumers/" + username, nil
	} else if kind == "models.Create" || kind == "models.Update" {
		// API Six doesn't provide an update function, it uses create
		if value.GetUsername() == "" && username == "" {
			return "", "", errors.New("You must provide a username to create a credential")
		} else if username == "" {
			username = value.GetUsername()
		}
		return "PUT", "/apisix/admin/consumers/" + username, nil
	} else if kind == "models.Delete" {
		if username == "" {
			return "", "", errors.New("You must provide a username so I know who to delete")
		}
		return "DELETE", "/apisix/admin/consumers/" + username, nil
	}
	return "", "", errors.New("Somehow a type got passed that wasn't registered")
}
func validateConsumerGroup[Value models.Value](id, kind string, value Value) (string, string, error) {
	// 	Get | Create | Update | Delete
	// Get the operation type
	if kind == "models.Get" {
		if id == "" {
			return "GET", "/apisix/admin/consumer_groups", nil
		}
		return "GET", "/apisix/admin/consumer_groups/" + id, nil
	} else if kind == "models.Create" {
		if id == "" {
			return "", "", errors.New("You must provide an ID")
		}
		return "PUT", "/apisix/admin/consumer_groups/" + id, nil
	} else if kind == "models.Update" {
		if id == "" {
			return "", "", errors.New("You must provide an ID")
		}
		return "PATCH", "/apisix/admin/consumer_groups/" + id, nil
	} else if kind == "models.Delete" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to delete")
		}
		return "DELETE", "/apisix/admin/consumer_groups/" + id, nil
	}
	return "", "", errors.New("Somehow a type got passed that wasn't registered")
}

func validateCredential[Value models.Value](id, username, kind string, value Value) (string, string, error) {
	// 	Get | Create | Update | Delete
	// Get the operation type
	if kind == "models.Get" {
		if username == "" {
			return "", "", errors.New("You must provide a username to get credential from")
		} else if id == "" {
			return "GET", "/apisix/admin/consumers/" + username + "/credentials", nil
		}
		return "GET", "/apisix/admin/consumers/" + username + "/credentials/" + id, nil
	} else if kind == "models.Create" || kind == "models.Update" {
		// API Six doesn't provide an update function, it uses create
		// Check to see if a username is passed anywhere
		if value.GetUsername() == "" && username == "" {
			return "", "", errors.New("You must provide a username and ID to create a credential")
		} else if username == "" {
			username = value.GetUsername()
		}
		// Check to see if an ID is provided anywhere
		if id == "" {
			return "", "", errors.New("You must provide a username and ID to create a credential")
		}
		return "PUT", "/apisix/admin/consumers/" + username + "/credentials/" + id, nil
	} else if kind == "models.Delete" {
		if username == "" || id == "" {
			return "", "", errors.New("You must provide a username and ID to delete a credential")
		}
		return "DELETE", "/apisix/admin/consumers/" + username + "/credentials/" + id, nil
	}
	return "", "", errors.New("Somehow a type got passed that wasn't registered")
}
func validateGlobalRule[Value models.Value](id, kind string, value Value) (string, string, error) {
	// 	Get | Create | Update | Delete
	// Get the operation type
	if kind == "models.Get" {
		if id == "" {
			return "GET", "/apisix/admin/global_rules", nil
		}
		return "GET", "/apisix/admin/global_rules/" + id, nil
	} else if kind == "models.Create" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to create a Global Rule")
		}
		return "PUT", "/apisix/admin/global_rules/" + id, nil
	} else if kind == "models.Update" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to update")
		}
		return "PATCH", "/apisix/admin/global_rules/" + id, nil
	} else if kind == "models.Delete" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to delete")
		}
		return "DELETE", "/apisix/admin/global_rules/" + id, nil
	}
	return "", "", errors.New("Somehow a type got passed that wasn't registered")
}
func validatePluginConfig[Value models.Value](id, kind string, value Value) (string, string, error) {
	// 	Get | Create | Update | Delete
	// Get the operation type
	if kind == "models.Get" {
		if id == "" {
			return "GET", "/apisix/admin/plugin_configs", nil
		}
		return "GET", "/apisix/admin/plugin_configs/" + id, nil
	} else if kind == "models.Create" {
		if id == "" {
			return "", "", errors.New("You must provide an ID")
		}
		return "PUT", "/apisix/admin/plugin_configs/" + id, nil
	} else if kind == "models.Update" {
		if id == "" {
			return "", "", errors.New("You must provide an ID")
		}
		return "PATCH", "/apisix/admin/plugin_configs/" + id, nil
	} else if kind == "models.Delete" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to delete")
		}
		return "DELETE", "/apisix/admin/plugin_configs/" + id, nil
	}
	return "", "", errors.New("Somehow a type got passed that wasn't registered")
}
func validateProto[Value models.Value](id, kind string, value Value) (string, string, error) {
	// 	Get | Create | Update | Delete
	// Get the operation type
	if kind == "models.Get" {
		if id == "" {
			return "GET", "/apisix/admin/protos", nil
		}
		return "GET", "/apisix/admin/protos/" + id, nil
	} else if kind == "models.Create" {
		if id == "" {
			return "POST", "/apisix/admin/protos", nil
		}
		return "PUT", "/apisix/admin/protos/" + id, nil
	} else if kind == "models.Update" {
		if id == "" {
			return "", "", errors.New("You must provide an ID")
		}
		return "PUT", "/apisix/admin/protos/" + id, nil
	} else if kind == "models.Delete" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to delete")
		}
		return "DELETE", "/apisix/admin/protos/" + id, nil
	}
	return "", "", errors.New("Somehow a type got passed that wasn't registered")
}
func validateRoute[Value models.Value](id, kind string, value Value) (string, string, error) {
	// 	Get | Create | Update | Delete
	// Get the operation type
	if kind == "models.Get" {
		if id == "" {
			return "GET", "/apisix/admin/routes", nil
		}
		return "GET", "/apisix/admin/routes/" + id, nil
	} else if kind == "models.Create" {
		ttl := ""
		if value.GetTTL() != "" {
			ttl = "?ttl=" + ttl
		}
		if value.GetURI() == "" && value.GetURIs() == nil {
			return "", "", errors.New("Only one of URI or URIs needs to be provided")
		} else if id == "" {
			return "POST", "/apisix/admin/routes" + ttl, nil
		}
		return "PUT", "/apisix/admin/routes/" + id + ttl, nil
	} else if kind == "models.Update" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to update a route")
		}
		return "PATCH", "/apisix/admin/routes/" + id, nil
	} else if kind == "models.Delete" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to delete")
		}
		return "DELETE", "/apisix/admin/routes/" + id, nil
	}
	return "", "", errors.New("Somehow a type got passed that wasn't registered")
}

func validateService[Value models.Value](id, kind string, value Value) (string, string, error) {
	// 	Get | Create | Update | Delete
	// Get the operation type
	if kind == "models.Get" {
		if id == "" {
			return "GET", "/apisix/admin/services", nil
		}
		return "GET", "/apisix/admin/services/" + id, nil
	} else if kind == "models.Create" {
		if id == "" {
			return "POST", "/apisix/admin/services", nil
		}
		return "PUT", "/apisix/admin/services/" + id, nil
	} else if kind == "models.Update" {
		if id == "" {
			return "", "", errors.New("You must provide an ID")
		}
		return "PATCH", "/apisix/admin/services/" + id, nil
	} else if kind == "models.Delete" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to delete")
		}
		return "DELETE", "/apisix/admin/services/" + id, nil
	}
	return "", "", errors.New("Somehow a type got passed that wasn't registered")
}
func validateSSL[Value models.Value](id, kind string, value Value) (string, string, error) {
	// 	Get | Create | Update | Delete
	// Get the operation type
	if kind == "models.Get" {
		if id == "" {
			return "GET", "/apisix/admin/ssls", nil
		}
		return "GET", "/apisix/admin/ssls/" + id, nil
	} else if kind == "models.Create" {
		if id == "" {
			return "POST", "/apisix/admin/ssls", nil
		}
		return "PUT", "/apisix/admin/ssls/" + id, nil
	} else if kind == "models.Update" {
		return "", "", errors.New("You cannot update an SSL credential")
	} else if kind == "models.Delete" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to delete")
		}
		return "DELETE", "/apisix/admin/ssls/" + id, nil
	}
	return "", "", errors.New("Somehow a type got passed that wasn't registered")
}

func validateSecret[Value models.Value](id, kind string, value Value) (string, string, error) {
	// 	Get | Create | Update | Delete
	// Get the operation type
	if kind == "models.Get" {
		if id == "" {
			return "GET", "/apisix/admin/secrets", nil
		} else if strings.Contains(id, "/") {
			return "GET", "/apisix/admin/secrets/" + id, nil
		} else {
			return "", "", errors.New("You need to include the manager and the ID in the format {manager}/{id}")
		}
	} else if kind == "models.Create" {
		if id == "" || !strings.Contains(id, "/") {
			return "", "", errors.New("You need to include the manager and the ID in the format {manager}/{id}")
		}
		return "PUT", "/apisix/admin/secrets/" + id, nil
	} else if kind == "models.Update" {
		if id == "" || !strings.Contains(id, "/") {
			return "", "", errors.New("You need to include the manager and the ID in the format {manager}/{id}")
		}
		return "PATCH", "/apisix/admin/secrets/" + id, nil
	} else if kind == "models.Delete" {
		if id == "" || !strings.Contains(id, "/") {
			return "", "", errors.New("You need to include the manager and the ID in the format {manager}/{id}")
		}
		return "DELETE", "/apisix/admin/secrets/" + id, nil
	}
	return "", "", errors.New("Somehow a type got passed that wasn't registered")
}

func validateStream[Value models.Value](id, kind string, value Value) (string, string, error) {
	// 	Get | Create | Update | Delete
	// Get the operation type
	if kind == "models.Get" {
		if id == "" {
			return "GET", "/apisix/admin/stream_routes", nil
		}
		return "GET", "/apisix/admin/stream_routes/" + id, nil
	} else if kind == "models.Create" {
		if id == "" {
			return "POST", "/apisix/admin/stream_routes", nil
		}
		return "PUT", "/apisix/admin/stream_routes/" + id, nil
	} else if kind == "models.Update" {
		return "", "", errors.New("You cannot update a Stream Route")
	} else if kind == "models.Delete" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to delete")
		}
		return "DELETE", "/apisix/admin/stream_routes/" + id, nil
	}
	return "", "", errors.New("Somehow a type got passed that wasn't registered")
}

func validateUpstream[Value models.Value](id, kind string, value Value) (string, string, error) {
	// 	Get | Create | Update | Delete
	// Get the operation type
	if kind == "models.Get" {
		if id == "" {
			return "GET", "/apisix/admin/upstreams", nil
		}
		return "GET", "/apisix/admin/upstreams/" + id, nil
	} else if kind == "models.Create" {
		if id == "" {
			return "POST", "/apisix/admin/upstreams", nil
		}
		return "PUT", "/apisix/admin/upstreams/" + id, nil
	} else if kind == "models.Update" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to update")
		}
		return "PATCH", "/apisix/admin/upstreams/" + id, nil
	} else if kind == "models.Delete" {
		if id == "" {
			return "", "", errors.New("You must provide an ID to delete")
		}
		return "DELETE", "/apisix/admin/upstreams/" + id, nil
	}
	return "", "", errors.New("Somehow a type got passed that wasn't registered")
}

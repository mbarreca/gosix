package consumer

import (
	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer/models"
)

// Add the Basic Auth Authentication Method to a consumer
func BasicAuthAdd(basic models.BasicAuth, client *gosix.Client) {
}

// Reset the users basic auth password to a random string and return it
func BasicAuthResetPassword(username, client *gosix.Client) {
}

// Reset the users basic auth password to a random string and return it
func BasicAuthChangePassword(username, password string, client *gosix.Client) {
}

// Disable Basic Auth for this consumer
func BasicAuthDisable(username, client *gosix.Client) {
}

// Enable Basic Auth for this consumer
func BasicAuthEnable(username, client *gosix.Client) {
}

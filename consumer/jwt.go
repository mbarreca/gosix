package consumer

import (
	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/consumer/models"
)

// Add key authentication to the selected consumer
func JWTAuthAdd(username string, key models.JWTAuth, client *gosix.Client) {

}

// Cycle the key and return
func JWTAuthCycle(username string, client *gosix.Client) {

}

// Disable JWT Auth for this consumer
func JWTAuthDisable(username, client *gosix.Client) {
}

// Enable JWT Auth for this consumer
func JWTAuthEnable(username, client *gosix.Client) {
}

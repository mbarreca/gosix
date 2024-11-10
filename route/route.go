package route

import "github.com/mbarreca/gosix"

// Fetches a list of all configured Routes.
func Get(client *gosix.Client) (GetResponse, error) {
}

// Fetches specified Route by id.
func GetByID(routeID string, client *gosix.Client) (GetResponse, error) {
}

// WARNING - Conflicts can occur. Creates a Route with the specified id.
func Put() {

}

// Creates a Route and assigns a random id.
func Post() {

}

// Removes the Route with the specified id.
func Delete() {

}

// Updates the selected attributes of the specified, existing Route. To delete an attribute, set value of attribute set to null.
func PatchByID() {

}

// Updates the attribute specified in the path. The values of other attributes remain unchanged.
func PatchByIDAndRoute() {

}

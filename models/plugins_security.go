package models

type ConsumerRestriction struct {
	Type           string                       `json:"type,omitempty" validate:"omitempty,alphanum,min=1,max=100"`         // Default - consumer_name - Valid Values - consumer_name, consumer_group_id, service_id, route_id - Type of object to base the restriction on.
	Whitelist      []string                     `json:"whitelist,omitempty" validate:"omitempty,dive"`                      // List of objects to whitelist. Has a higher priority than allowed_by_methods.
	Blacklist      []string                     `json:"blacklist,omitempty" validate:"omitempty,dive"`                      // List of objects to blacklist. Has a higher priority than whitelist.
	RejectedCode   int                          `json:"rejected_code,omitempty" validate:"omitempty,number"`                // Default - 403 - Valid Values - HTTP status code returned when the request is rejected.
	RejectedMsg    string                       `json:"rejected_msg,omitempty" validate:"omitempty,alphanum,min=1,max=250"` // Message returned when the request is rejected.
	AllowByMethods []*RestrictionAllowByMethods `json:"allowed_by_methods,omitempty" validate:"omitempty"`                  // List of allowed configurations for Consumer settings, including a username of the Consumer and a list of allowed HTTP methods.

}

type RestrictionAllowByMethods struct {
	User    string   `json:"user,omitempty" validate:"omitempty,dive"`    // A username for a Consumer.
	Methods []string `json:"methods,omitempty" validate:"omitempty,dive"` // ["GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "CONNECT", "TRACE", "PURGE"]	List of allowed HTTP methods for a Consumer.
}

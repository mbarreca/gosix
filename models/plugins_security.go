package models

type ConsumerRestriction struct {
	Type           string                       `json:"type,omitempty" validate:"omitempty,ascii,min=1,max=100"`            // Default - consumer_name - Valid Values - consumer_name, consumer_group_id, service_id, route_id - Type of object to base the restriction on.
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

type CORS struct {
	AllowOrigins           string   `json:"allow_origins,omitempty" validate:"omitempty,ascii"`          //	Origins to allow CORS. Use the scheme://host:port format. For example, https://somedomain.com:8081. If you have multiple origins, use a , to list them. If allow_credential is set to false, you can enable CORS for all origins by using *. If allow_credential is set to true, you can forcefully allow CORS on all origins by using ** but it will pose some security issues.
	AllowMethods           string   `json:"allow_methods,omitempty" validate:"omitempty,ascii"`          // Request methods to enable CORS on. For example GET, POST. Use , to add multiple methods. If allow_credential is set to false, you can enable CORS for all methods by using *. If allow_credential is set to true, you can forcefully allow CORS on all methods by using ** but it will pose some security issues.
	AllowHeaders           string   `json:"allow_headers,omitempty" validate:"omitempty,ascii"`          // Headers in the request allowed when accessing a cross-origin resource. Use , to add multiple headers. If allow_credential is set to false, you can enable CORS for all request headers by using *. If allow_credential is set to true, you can forcefully allow CORS on all request headers by using ** but it will pose some security issues.
	ExposeHeaders          string   `json:"expose_headers,omitempty" validate:"omitempty,ascii"`         // Headers in the response allowed when accessing a cross-origin resource. Use , to add multiple headers. If allow_credential is set to false, you can enable CORS for all response headers by using *. If not specified, the plugin will not modify the Access-Control-Expose-Headers header. See Access-Control-Expose-Headers - MDN for more details.
	MaxAge                 int      `json:"max_age,omitempty" validate:"omitempty,number"`               // Maximum time in seconds the result is cached. If the time is within this limit, the browser will check the cached result. Set to -1 to disable caching. Note that the maximum value is browser dependent. See Access-Control-Max-Age for more details.
	AllowCredential        bool     `json:"allow_credential,omitempty" validate:"omitempty,boolean"`     // When set to true, allows requests to include credentials like cookies. According to CORS specification, if you set this to true, you cannot use '*' to allow all for the other attributes.
	AllowOriginsByRegex    []string `json:"allow_origins_by_regex,omitempty" validate:"omitempty,ascii"` // Regex to match origins that allow CORS. For example, [".*\.test.com$"] can match all subdomains of test.com. When set to specified range, only domains in this range will be allowed, no matter what allow_origins is.
	AllowOriginsByMetadata []string `json:"allow_origins_by_metadata,omitempty" validate:"omitempty,ascii"`
}

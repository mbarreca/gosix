package models

type Checks struct {
	Active  Active  `json:"active,omitempty"`
	Passive Passive `json:"checks,omitempty"`
}

type Active struct {
	Timeout                int       `json:"timeout,omitempty" validate:"omitempty,number"`               // The timeout period of the active check (unit: second).
	HTTPPath               string    `json:"http_path,omitempty" validate:"omitempty,uri,min=1,max=1000"` // The HTTP request path that is actively checked.
	Host                   string    `json:"host,omitempty" validate:"omitempty,url,min=1,max=1000"`      // The hostname of the HTTP request actively checked.
	Port                   int       `json:"port,omitempty" validate:"omitempty,number"`                  // The host port of the HTTP request that is actively checked.
	Concurrency            int       `json:"concurrency,omitempty" validate:"omitempty,number"`           // The number of targets to be checked at the same time during the active check.
	Type                   string    `json:"type,omitempty" validate:"omitempty,ascii,min=1,max=1000"`    // The type of passive check.
	Healthy                Healthy   `json:"healthy,omitempty"`
	Unhealthy              Unhealthy `json:"unhealthy,omitempty"`
	ReqHeaders             []string  `json:"req_headers,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	HTTPSVerifyCertificate bool      `json:"https_verify_certificate,omitempty" validate:"omitempty,boolean"` // Active check whether to check the SSL certificate of the remote host when HTTPS type checking is used.
}
type Passive struct {
	Type      string    `json:"type,omitempty" validate:"omitempty,uri,min=1,max=1000"` // The type of passive check.
	Healthy   Healthy   `json:"healthy,omitempty"`
	Unhealthy Unhealthy `json:"unhealthy,omitempty"`
}
type Healthy struct {
	HTTPStatuses []int `json:"http_statuses,omitempty" validate:"omitempty,number"` // Active check (healthy node) HTTP or HTTPS type check, the HTTP status code of the healthy node.
	Interval     int   `json:"interval,omitempty" validate:"omitempty,number"`      // Active check (healthy node) check interval (unit: second)
	Successes    int   `json:"successes,omitempty" validate:"omitempty,number"`     // Active check (healthy node) determine the number of times a node is healthy.
}
type Unhealthy struct {
	Interval     int   `json:"interval,omitempty" validate:"omitempty,number"`      // Active check (unhealthy node) check interval (unit: second)
	HTTPStatuses []int `json:"http_statuses,omitempty" validate:"omitempty,number"` // Active check (unhealthy node) HTTP or HTTPS type check, the HTTP status code of the non-healthy node.
	HTTPFailures int   `json:"http_failures,omitempty" validate:"omitempty,number"` // Active check (unhealthy node) HTTP or HTTPS type check, determine the number of times that the node is not healthy.
	TCPFailures  int   `json:"tcp_failures,omitempty" validate:"omitempty,number"`  // Active check (unhealthy node) TCP type check, determine the number of times that the node is not healthy.
	Timeouts     int   `json:"timeouts,omitempty" validate:"omitempty,number"`      // Active check (unhealthy node) to determine the number of timeouts for unhealthy nodes.
}

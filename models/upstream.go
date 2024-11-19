package models

type Upstream struct {
	ID            string         `json:"id,omitempty" validate:"omitempty,ascii,min=1,max=64"`
	Name          string         `json:"name" validate:"required,ascii,min=1,max=500"`
	Desc          string         `json:"desc,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	Retries       int            `json:"retries,omitempty" validate:"omitempty,number"`
	DiscoveryType string         `json:"discovery_type,omitempty" validate:"omitempty,ascii,min=1,max=1000"` // True, if service_name is used
	ServiceName   string         `json:"service_name,omitempty" validate:"omitempty,ascii,min=1,max=1000"`   // True, can't be used with nodes - Service name used for service discovery.
	Nodes         map[string]int `json:"nodes,omitempty"`                                                    // IP addresses (with optional ports) of the Upstream nodes represented as a hash table or an array. In the hash table, the key is the IP address and the value is the weight of the node for the load balancing algorithm. For hash table case, if the key is IPv6 address with port, then the IPv6 address must be quoted with square brackets. In the array, each item is a hash table with keys host, weight, and the optional port and priority (defaults to 0). Nodes with lower priority are used only when all nodes with a higher priority are tried and are unavailable. Empty nodes are treated as placeholders and clients trying to access this Upstream will receive a 502 response.
	Type          string         `json:"type,omitempty" validate:"omitempty,ascii,min=1,max=1000"`           // Load balancing algorithm to be used, and the default value is roundrobin.
	KeepAlivePool KeepAlivePool  `json:"keepalive_pool,omitempty"`
	Checks        *Checks        `json:"checks,omitempty"`                                            // Configures the parameters for the health check.
	HashOn        string         `json:"hash_on,omitempty" validate:"omitempty,ascii,min=1,max=1000"` // Only valid if the type is chash. Supports Nginx variables (vars), custom headers (header), cookie and consumer. Defaults to vars.
	Key           string         `json:"key,omitempty" validate:"omitempty,ascii,min=1,max=1000"`     // Only valid if the type is chash. Finds the corresponding node id according to hash_on and key values. When hash_on is set to vars, key is a required parameter and it supports Nginx variables. When hash_on is set as header, key is a required parameter, and header name can be customized. When hash_on is set to cookie, key is also a required parameter, and cookie name can be customized. When hash_on is set to consumer, key need not be set and the key used by the hash algorithm would be the authenticated consumer_name.
	Scheme        string         `json:"scheme,omitempty" validate:"omitempty,ascii,min=1,max=1000"`
	Timeout       *Timeout       `json:"timeout,omitempty"`
	PassHost      string         `json:"pass_host,omitempty" validate:"omitempty,ascii,min=1,max=1000"`
	CreateTime    int            `json:"create_time,omitempty" validate:"omitempty,number"`
	UpdateTime    int            `json:"update_time,omitempty" validate:"omitempty,number"`
}

type KeepAlivePool struct {
	IdleTimeout int `json:"idle_timeout,omitempty" validate:"omitempty,number"`
	Requests    int `json:"requests,omitempty" validate:"omitempty,number"`
	Size        int `json:"size,omitempty" validate:"omitempty,number"`
}

func (c Upstream) GetID() string {
	return c.ID
}
func (c Upstream) GetName() string {
	return c.Name
}
func (c Upstream) GetTTL() string {
	return ""
}
func (c Upstream) GetURI() string {
	return ""
}
func (c Upstream) GetURIs() []string {
	return nil
}
func (c Upstream) GetUsername() string {
	return ""
}

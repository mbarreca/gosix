package models

type Route struct {
	TTL             string     `json:"-"`
	ID              string     `json:"id,omitempty" validate:"omitempty,ascii,min=1,max=64"`
	Desc            string     `json:"desc,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	Name            string     `json:"name" validate:"required,ascii,min=1,max=500"`
	RemoteAddr      string     `json:"remote_addr,omitempty" validate:"omitempty,ascii,min=1,max=1000"`
	RemoteAddrs     []string   `json:"remote_addrs,omitempty" validate:"omitempty,ascii,min=1,max=1000"`
	Methods         []string   `json:"methods" validate:"required,min=1,max=1000"`
	Vars            [][]string `json:"vars,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	FilterFunc      string     `json:"filter_func,omitempty" validate:"omitempty,ascii,min=1,max=1000"`
	PluginConfigID  string     `json:"plugin_config_id,omitempty" validate:"omitempty,ascii,min=1,max=1000"`
	Plugins         *Plugins   `json:"plugins,omitempty" validate:"omitempty"`
	Priority        int        `json:"priority,omitempty" validate:"omitempty,number"`
	Script          string     `json:"script,omitempty" validate:"omitempty,ascii,min=1,max=1000"`
	Upstream        *Upstream  `json:"upstream,omitempty" validate:"omitempty"`
	UpstreamID      string     `json:"upstream_id,omitempty" validate:"omitempty,ascii,min=1,max=200"`
	ServiceID       string     `json:"service_id,omitempty" validate:"omitempty,ascii,min=1,max=1000"`
	Timestamp       *Timeout   `json:"timestamp,omitempty" validate:"omitempty"`
	EnableWebsocket bool       `json:"enable_websocket,omitempty" validate:"omitempty,boolean"`
	Status          int        `json:"status,omitempty" validate:"omitempty,number"`
	Hosts           []string   `json:"hosts,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	URI             string     `json:"uri,omitempty" validate:"omitempty,uri,min=1,max=1000"`
	URIs            []string   `json:"uris,omitempty" validate:"omitempty,min=1,max=2000"`
	CreateTime      int        `json:"create_time,omitempty" validate:"omitempty,number"`
	UpdateTime      int        `json:"update_time,omitempty" validate:"omitempty,number"`
}

func (c Route) GetID() string {
	return c.ID
}
func (c Route) GetName() string {
	return c.Name
}
func (c Route) GetTTL() string {
	return c.TTL
}
func (c Route) GetURI() string {
	return c.URI
}
func (c Route) GetURIs() []string {
	return c.URIs
}
func (c Route) GetUsername() string {
	return ""
}

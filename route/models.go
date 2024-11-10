package route

type RouteCreate struct {
	ID      string   `json:"id"`
	Uris    []string `json:"uris"`
	Methods []string `json:"methods"`
	Hosts   []string `json:"hosts"`
	Plugins struct {
	} `json:"plugins"`
	Priority    int        `json:"priority"`
	Name        string     `json:"name"`
	Desc        string     `json:"desc"`
	RemoteAddrs []string   `json:"remote_addrs"`
	Vars        [][]string `json:"vars"`
	UpstreamID  string     `json:"upstream_id"`
	Upstream    Upstream   `json:"upstream"`
	Timeout     Timeout    `json:"timeout"`
	FilterFunc  string     `json:"filter_func"`
}

type Timeout struct {
	Connect int `json:"connect"`
	Send    int `json:"send"`
	Read    int `json:"read"`
}
type Upstream struct {
	ID      string `json:"id"`
	Retries int    `json:"retries"`
	Timeout struct {
		Connect int `json:"connect"`
		Send    int `json:"send"`
		Read    int `json:"read"`
	} `json:"timeout"`
	Nodes  map[string]string `json:"nodes"`
	Type   string            `json:"type"`
	Checks struct {
	} `json:"checks"`
	HashOn string `json:"hash_on"`
	Key    string `json:"key"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Scheme string `json:"scheme"`
}

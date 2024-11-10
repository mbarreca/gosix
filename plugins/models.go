package plugins

type Plugins struct {
	BasicAuth BasicAuth `json:"basic-auth"`
	JwtAuth   JwtAuth   `json:"jwt-auth"`
	KeyAuth   KeyAuth   `json:"key-auth"`
}

type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Meta     struct {
		Disable bool `json:"disable"`
	} `json:"_meta"`
}

type JwtAuth struct {
	Key  string `json:"key"`
	Exp  int    `json:"exp"`
	Meta struct {
		Disable bool `json:"disable"`
	} `json:"_meta"`
}
type KeyAuth struct {
	Key  string `json:"key"`
	Exp  int    `json:"exp"`
	Meta struct {
		Disable bool `json:"disable"`
	} `json:"_meta"`
}

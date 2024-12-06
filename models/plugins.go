package models

type Plugins struct {
	// Authentication Plugins
	BasicAuth *BasicAuth `json:"basic-auth,omitempty" validate:"omitempty"`
	JwtAuth   *JwtAuth   `json:"jwt-auth,omitempty" validate:"omitempty"`
	KeyAuth   *KeyAuth   `json:"key-auth,omitempty" validate:"omitempty"`
	// Security Plugins
	ConsumerRestriction *ConsumerRestriction `json:"consumer-restriction,omitempty" validate:"omitempty"`
	CORS                *CORS                `json:"cors,omitempty" validate:"omitempty"`
}

type PluginConfig struct {
	ID         string            `json:"id,omitempty" validate:"omitempty,ascii,min=1,max=64"`
	Desc       string            `json:"desc,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	Labels     map[string]string `json:"labels,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	Plugins    *Plugins          `json:"plugins,omitempty" validate:"omitempty"`
	CreateTime int               `json:"create_time,omitempty" validate:"omitempty,number"`
	UpdateTime int               `json:"update_time,omitempty" validate:"omitempty,number"`
}

func (c PluginConfig) GetID() string {
	return c.ID
}
func (c PluginConfig) GetName() string {
	return ""
}
func (c PluginConfig) GetTTL() string {
	return ""
}
func (c PluginConfig) GetURI() string {
	return ""
}
func (c PluginConfig) GetURIs() []string {
	return nil
}
func (c PluginConfig) GetUsername() string {
	return ""
}

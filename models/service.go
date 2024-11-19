package models

type Service struct {
	ID              string            `json:"id,omitempty" validate:"omitempty,ascii,min=1,max=64"`
	Name            string            `json:"name" validate:"required,ascii,min=1,max=500"`
	Desc            string            `json:"desc,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	UpstreamID      string            `json:"upstream_id,omitempty" validate:"omitempty,ascii,min=1,max=200"`
	Upstream        Upstream          `json:"upstream,required" validate:"omitempty"`
	EnableWebsocket bool              `json:"enable_websocket,omitempty" validate:"omitempty,boolean"`
	Labels          map[string]string `json:"labels,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	Plugins         *Plugins          `json:"plugins,omitempty" validate:"omitempty"`
	Hosts           []string          `json:"hosts,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	CreateTime      int               `json:"create_time,omitempty" validate:"omitempty,number"`
	UpdateTime      int               `json:"update_time,omitempty" validate:"omitempty,number"`
}

func (c Service) GetID() string {
	return c.ID
}
func (c Service) GetName() string {
	return c.Name
}
func (c Service) GetTTL() string {
	return ""
}
func (c Service) GetURI() string {
	return ""
}
func (c Service) GetURIs() []string {
	return nil
}
func (c Service) GetUsername() string {
	return ""
}

package models

type Stream struct {
	ID         string    `json:"id,omitempty" validate:"omitempty,ascii,min=1,max=64"`
	Desc       string    `json:"desc,omitempty" validate:"omitempty,ascii,min=1,max=2000"`
	Protocol   *Protocol `json:"protocol,omitempty" validate:"omitempty"`
	RemoteAddr string    `json:"remote_addr,omitempty" validate:"omitempty,ascii,min=1,max=1000"`
	ServerAddr string    `json:"server_addr,omitempty" validate:"omitempty,ascii,min=1,max=1000"`
	ServerPort string    `json:"server_port,omitempty" validate:"omitempty,alphanum,min=1,max=7"`
	ServiceID  string    `json:"service_id,omitempty" validate:"omitempty,ascii,min=1,max=1000"`
	SNIs       []string  `json:"snis,omitempty" validate:"omitempty,ascii"`
	Upstream   *Upstream `json:"upstream,required" validate:"omitempty"`
	UpstreamID string    `json:"upstream_id,omitempty" validate:"omitempty,ascii,min=1,max=200"`
	CreateTime int       `json:"create_time,omitempty" validate:"omitempty,number"`
	UpdateTime int       `json:"update_time,omitempty" validate:"omitempty,number"`
}

type Protocol struct {
	Name string `json:"name,omitempty" validate:"omitempty,ascii,min=1,max=500"`
	Conf string `json:"conf,omitempty" validate:"omitempty,ascii,min=1,max=10000"`
}

func (c Stream) GetID() string {
	return c.ID
}
func (c Stream) GetName() string {
	return ""
}
func (c Stream) GetTTL() string {
	return ""
}
func (c Stream) GetURI() string {
	return ""
}
func (c Stream) GetURIs() []string {
	return nil
}
func (c Stream) GetUsername() string {
	return ""
}

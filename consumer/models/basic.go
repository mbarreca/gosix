package models

type BasicAuth struct {
	Username string `json:"username,omitempty" validate:"required,alphanum,min=1,max=200"` // Unique username for a Consumer. If multiple Consumers use the same username, a request matching exception is raised.
	Password string `json:"password,omitempty" validate:"required,alphanum,min=1,max=200"` // Password of the user. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	Meta     *Meta  `json:"_meta,omitempty" validate:"omitempty"`
}

type Meta struct {
	Disable bool `json:"disable,omitempty" validate:"omitempty,boolean"`
}

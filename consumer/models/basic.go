package models

type BasicAuth struct {
	Username string `json:"username,omitempty" validate:"required,email"`               // This library uses emails as usernames - Must be Unique
	Password string `json:"password,omitempty" validate:"required,ascii,min=1,max=200"` // Password of the user. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	Meta     *Meta  `json:"_meta,omitempty" validate:"omitempty"`
}

type Meta struct {
	Disable bool `json:"disable,omitempty" validate:"omitempty,boolean"`
}

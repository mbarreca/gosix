package models

type SSL struct {
	ID           string            `json:"id,omitempty" validate:"omitempty,ascii,min=1,max=64"`
	Key          string            `json:"key,omitempty" validate:"omitempty"`         // HTTPS private key. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	Keys         []string          `json:"keys,omitempty" validate:"omitempty,ascii"`  // Private keys to pair with the certs. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	Cert         string            `json:"cert,omitempty" validate:"omitempty"`        // HTTPS certificate. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	Certs        []string          `json:"certs,omitempty" validate:"omitempty,ascii"` // Used for configuring multiple certificates for the same domain excluding the one provided in the cert field. This field supports saving the value in Secret Manager using the APISIX Secret resource.
	Client       *Client           `json:"client,omitempty"`
	SNIs         []string          `json:"snis,omitempty" validate:"omitempty,ascii"`
	Labels       map[string]string `json:"labels,omitempty" validate:"omitempty,ascii"`
	SSLProtocols []string          `json:"ssl_protocols,omitempty" validate:"omitempty,ascii"`
	Type         string            `json:"type,omitempty" validate:"omitempty,ascii,min=1,max=1000"`
	Status       int               `json:"status,omitempty" validate:"omitempty,number"`
	CreateTime   int               `json:"create_time,omitempty" validate:"omitempty,number"`
	UpdateTime   int               `json:"update_time,omitempty" validate:"omitempty,number"`
}

type Client struct {
	CA               string   `json:"ca,omitempty" validate:"omitempty,ascii,min=1,max=4096"`    // Sets the CA certificate that verifies the client. Requires OpenResty 1.19+.
	Depth            string   `json:"depth,omitempty" validate:"omitempty,ascii,min=1,max=4096"` // Sets the verification depth in client certificate chains. Defaults to 1. Requires OpenResty 1.19+.
	SkipMTLSUTIRegex []string `json:"skip_mtls_uri_regex,omitempty" validate:"omitempty,ascii"`  // Used to match URI, if matched, this request bypasses the client certificate checking, i.e. skip the MTLS.
}

func (c SSL) GetID() string {
	return c.ID
}
func (c SSL) GetName() string {
	return ""
}
func (c SSL) GetTTL() string {
	return ""
}
func (c SSL) GetURI() string {
	return ""
}
func (c SSL) GetURIs() []string {
	return nil
}
func (c SSL) GetUsername() string {
	return ""
}

package gosix

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"time"

	"golang.org/x/net/publicsuffix"
)

// The client object holds the request context and the HTTP Client.
// This is designed to increase telemetry visibility.
type Client struct {
	Client *http.Client
	Ctx    context.Context
}

// Constructor
func New() (*Client, error) {
	// Create new cookiejar for holding cookies
	// This *may not* be necessary, however in my experience Apache
	// Projects do use cookies and it doesn't hurt to hold them for purposes
	// of accessing them at a certain point in time.
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		return nil, err
	}
	// Create new http client with predefined options
	client := &http.Client{
		Jar:     jar,
		Timeout: time.Second * 5,
	}
	ctx := context.Background()
	return &Client{client: client, ctx: ctx}, nil
}

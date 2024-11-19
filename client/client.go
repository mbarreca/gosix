package client

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/mbarreca/godistcache"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"golang.org/x/net/publicsuffix"
)

// The client object holds the request context and the HTTP Client.
// This is designed to increase telemetry visibility.
type Client struct {
	Client *http.Client
	Ctx    context.Context
	Cache  *godistcache.Cache
}

// Constructor - You *shouldn't* be using this
func New(ctx context.Context, otel bool) (*Client, error) {
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
	// Add telemetry to the endpoint requests
	if otel {
		client = &http.Client{
			Jar:       jar,
			Timeout:   time.Second * 5,
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		}
	}
	return &Client{Client: client, Ctx: ctx}, nil
}

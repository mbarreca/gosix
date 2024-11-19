package gosix

import (
	"context"

	"github.com/mbarreca/gosix/authentication"
	"github.com/mbarreca/gosix/client"
	"github.com/mbarreca/gosix/consumer"
	"github.com/mbarreca/gosix/consumergroup"
	"github.com/mbarreca/gosix/globalrule"
	"github.com/mbarreca/gosix/plugin"
	"github.com/mbarreca/gosix/proto"
	"github.com/mbarreca/gosix/route"
	"github.com/mbarreca/gosix/secret"
	"github.com/mbarreca/gosix/service"
	"github.com/mbarreca/gosix/ssl"
	"github.com/mbarreca/gosix/stream"
	"github.com/mbarreca/gosix/upstream"
)

// The client object holds the request context and the HTTP Client.
// This is designed to increase telemetry visibility.

// | Proto | Secret | StreamRoute
type APISix struct {
	Client        *client.Client
	Consumer      *consumer.Consumer
	ConsumerGroup *consumergroup.ConsumerGroup
	Credential    *consumer.Credential
	GlobalRule    *globalrule.GlobalRule
	PluginConfig  *plugin.PluginConfig
	Proto         *proto.Proto
	Route         *route.Route
	Secret        *secret.Secret
	Service       *service.Service
	SSL           *ssl.SSL
	Stream        *stream.Stream
	Upstream      *upstream.Upstream
	Basic         *authentication.Basic
	JWT           *authentication.JWT
	Key           *authentication.Key
}

// Constructor
func New(ctx context.Context, otel bool) (*APISix, error) {
	// Create Client
	c, err := client.New(ctx, otel)
	if err != nil {
		return nil, err
	}
	// Create Consumer
	consumerObject := consumer.New(c)
	consumergroup := consumergroup.New(c, consumerObject)
	credential := consumer.NewCredential(c, consumerObject)
	globalrule := globalrule.New(c, consumerObject)
	pluginconfig := plugin.New(c, consumerObject)
	proto := proto.New(c, consumerObject)
	route := route.New(c, consumerObject)
	secret := secret.New(c, consumerObject)
	service := service.New(c, consumerObject)
	ssl := ssl.New(c, consumerObject)
	stream := stream.New(c, consumerObject)
	upstream := upstream.New(c, consumerObject)
	basic := authentication.NewBasic(c, consumerObject)
	jwt := authentication.NewJWT(c, consumerObject)
	key := authentication.NewKey(c, consumerObject)
	return &APISix{Client: c, Consumer: consumerObject, ConsumerGroup: consumergroup,
		Credential: credential, GlobalRule: globalrule, PluginConfig: pluginconfig,
		Proto: proto, Route: route, Secret: secret, Service: service, SSL: ssl,
		Stream: stream, Upstream: upstream, Basic: basic, JWT: jwt, Key: key}, nil
}

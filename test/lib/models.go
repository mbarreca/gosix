package lib

import (
	"github.com/mbarreca/gosix/consumergroup"
	"github.com/mbarreca/gosix/globalrule"
	"github.com/mbarreca/gosix/models"
	"github.com/mbarreca/gosix/plugin"
	"github.com/mbarreca/gosix/proto"
	"github.com/mbarreca/gosix/route"
	"github.com/mbarreca/gosix/secret"
	"github.com/mbarreca/gosix/service"
	"github.com/mbarreca/gosix/ssl"
	"github.com/mbarreca/gosix/stream"
	"github.com/mbarreca/gosix/upstream"
)

type APIs[V models.Value] interface {
	*consumergroup.ConsumerGroup | *globalrule.GlobalRule | *proto.Proto | *plugin.PluginConfig | *route.Route | *secret.Secret | *service.Service | *ssl.SSL | *stream.Stream | *upstream.Upstream
	Create(string, V) error
	Get(string) (V, error)
	GetAll() ([]V, error)
	Update(string, V) error
	Delete(string) error
}

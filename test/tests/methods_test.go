package tests

import (
	"os"
	"testing"

	"github.com/mbarreca/gosix/models"
	"github.com/mbarreca/gosix/test/lib"
)

// Test Consumer Group
func TestConsumerGroup(t *testing.T) {
	// Create APISix Object
	a, _, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	// Create objects
	id := "GoSixTestID"
	var value models.ConsumerGroup
	var updated models.ConsumerGroup
	value.Desc = "GoSix Library Testing"
	value.Plugins = new(models.Plugins)
	value.Plugins.BasicAuth = new(models.BasicAuth)
	value.Plugins.BasicAuth.Username = "A"
	value.Plugins.BasicAuth.Password = "B"
	updated.Plugins = new(models.Plugins)
	updated.Plugins.BasicAuth = new(models.BasicAuth)
	updated.Plugins.BasicAuth.Username = "UpdatedA"
	updated.Plugins.BasicAuth.Password = "UpdatedB"

	// Call Test
	v, e, err := lib.Test(id, value, updated, a.ConsumerGroup, a)
	if err != nil {
		t.Fatal(err)
	}

	// Confirm Data matches
	b1 := v.Plugins.BasicAuth
	b2 := e.Plugins.BasicAuth
	if b1.Username != "A" && b1.Password != "B" && b2.Username != "UpdatedA" && b2.Password != "UpdatedB" {
		t.Fatal("Updating Failed")
	}
}

// Test GlobalRule
func TestGlobalRule(t *testing.T) {
	// Create APISix Object
	a, _, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	// Create objects
	id := "GoSixTestID"
	var value models.GlobalRule
	var updated models.GlobalRule
	value.Plugins = new(models.Plugins)
	value.Plugins.BasicAuth = new(models.BasicAuth)
	value.Plugins.BasicAuth.Username = "A"
	value.Plugins.BasicAuth.Password = "B"
	updated.Plugins = new(models.Plugins)
	updated.Plugins.BasicAuth = new(models.BasicAuth)
	updated.Plugins.BasicAuth.Username = "UpdatedA"
	updated.Plugins.BasicAuth.Password = "UpdatedB"

	// Call Test
	v, e, err := lib.Test(id, value, updated, a.GlobalRule, a)
	if err != nil {
		t.Fatal(err)
	}
	// Confirm Data matches
	b1 := v.Plugins.BasicAuth
	b2 := e.Plugins.BasicAuth
	if b1.Username != "A" && b1.Password != "B" && b2.Username != "UpdatedA" && b2.Password != "UpdatedB" {
		t.Fatal("Updating Global Rule Failed")
	}
}

// Test Plugin Config
func TestPluginConfig(t *testing.T) {
	// Create APISix Object
	a, _, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	// Create objects
	id := "GoSixTestID"
	var value models.PluginConfig
	var updated models.PluginConfig
	value.Desc = "GoSix Library Testing"
	value.Plugins = new(models.Plugins)
	value.Plugins.BasicAuth = new(models.BasicAuth)
	value.Plugins.BasicAuth.Username = "A"
	value.Plugins.BasicAuth.Password = "B"
	updated.Plugins = new(models.Plugins)
	updated.Plugins.BasicAuth = new(models.BasicAuth)
	updated.Plugins.BasicAuth.Username = "UpdatedA"
	updated.Plugins.BasicAuth.Password = "UpdatedB"

	// Call Test
	v, e, err := lib.Test(id, value, updated, a.PluginConfig, a)
	if err != nil {
		t.Fatal(err)
	}

	// Confirm Data matches
	b1 := v.Plugins.BasicAuth
	b2 := e.Plugins.BasicAuth
	if b1.Username != "A" && b1.Password != "B" && b2.Username != "UpdatedA" && b2.Password != "UpdatedB" {
		t.Fatal("Updating Failed")
	}
}

// Test Proto
func TestProto(t *testing.T) {
	// Create APISix Object
	a, _, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	// Create objects
	id := "GoSixTestID"
	var value models.Proto
	var updated models.Proto
	content := "syntax = \"proto3\";package helloworld;service Greeter {rpc SayHello (HelloRequest) returns (HelloReply) {}}message HelloRequest {string name = 1;}message HelloReply {string message = 1;}"
	updatedContent := "syntax = \"proto3\";package helloworld;service Greeter {rpc SayHello (HelloRequest) returns (HelloReply) {}}message HelloRequest {string name = 2;}message HelloReply {string message = 1;}"
	value.Content = content
	updated.Content = updatedContent

	// Call Test
	v, e, err := lib.Test(id, value, updated, a.Proto, a)
	if err != nil {
		t.Fatal(err)
	}

	// Confirm Data matches
	if v.Content != content && e.Content != updatedContent {
		t.Fatal("Updating Failed")
	}
}

// Test Route
func TestRoute(t *testing.T) {
	// Create APISix Object
	a, _, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	// Create objects
	id := "GoSixTestID"
	var r models.Route
	var rU models.Route
	r.Desc = "GoSix Library Testing Route"
	r.Name = "Testing Route"
	methods := make([]string, 0)
	methods = append(methods, "GET", "POST")
	r.Methods = methods
	addrs := make([]string, 0)
	addrs = append(addrs, "127.0.0.0/8")
	r.RemoteAddrs = addrs
	uris := make([]string, 0)
	uris = append(uris, "/index.html")
	r.URIs = uris
	var u models.Upstream
	m := make(map[string]int)
	m["127.0.0.1:1980"] = 1
	u.Nodes = m
	u.Name = "Test Upstream"
	u.Type = "roundrobin"
	r.Upstream = &u
	rU = r
	rU.Desc = "GoSix Library Testing Route Updated"
	rU.Name = "Testing Route Updated"

	// Call Test
	v, e, err := lib.Test(id, r, rU, a.Route, a)
	if err != nil {
		t.Fatal(err)
	}
	// Confirm Data matches
	if v.Name != "Testing Route" && v.Desc != "GoSix Library Testing Route" && e.Name != "Testing Route Updated" && e.Desc != "GoSix Library Testing Route Updated" {
		t.Fatal("Updating Route Rule Failed")
	}
	// Add Route
	if err := a.Route.Create(id, r); err != nil {
		t.Fatal(err)
	}
	// Add Plugin
	ba := new(models.BasicAuth)
	ba.Username = "A"
	ba.Password = "B"
	if err := a.Route.AddPlugin(id, ba); err != nil {
		t.Fatal(err)
	}
	baG, err := a.Route.GetPlugin(id, models.BasicAuth{})
	if err != nil {
		t.Fatal(err)
	}
	baGet := baG.(*models.BasicAuth)
	if baGet.Username != ba.Username || baGet.Password != ba.Password {
		t.Fatal("Add/Get Plugin Failed")
	}

	// Cleanup
	if err := a.Route.Delete(id); err != nil {
		t.Fatal(err)
	}
}

// Test Secret
func TestSecret(t *testing.T) {
	// Create APISix Object
	a, _, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}

	// Create objects
	id := "vault/1"
	var value models.Secret
	var updated models.Secret
	uri := "http://www/get"
	uriU := "http://www/updated"
	pre := "Apisix"
	preU := "Updated"
	tok := "Apisix"
	tokU := "Updated"
	value.URI = uri
	value.Prefix = pre
	value.Token = tok
	updated.URI = uriU
	updated.Prefix = preU
	updated.Token = tokU

	// Call Test
	v, e, err := lib.Test(id, value, updated, a.Secret, a)
	if err != nil {
		t.Fatal(err)
	}

	// Confirm Data matches
	if v.URI != uri && v.Prefix != pre && v.Token != tok && e.URI != uriU && e.Prefix != preU && e.Token != tokU {
		t.Fatal("Updating Failed")
	}
}

// Test Service
func TestService(t *testing.T) {
	// Create APISix Object
	a, _, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	// Create objects
	id := "GoSixTestID"
	var s models.Service
	var sU models.Service
	s.Desc = "GoSix Library Testing Service"
	s.Name = "Testing Service"
	var u models.Upstream
	m := make(map[string]int)
	m["127.0.0.1:1980"] = 1
	u.Nodes = m
	u.Name = "Test Upstream"
	u.Type = "roundrobin"
	s.Upstream = u
	sU = s
	sU.Desc = "GoSix Library Testing Service Updated"
	sU.Name = "Testing Service Updated"

	// Call Test
	v, e, err := lib.Test(id, s, sU, a.Service, a)
	if err != nil {
		t.Fatal(err)
	}
	// Confirm Data matches
	if v.Name != "Testing Service" && v.Desc != "GoSix Library Testing Service" && e.Name != "Testing Service Updated" && e.Desc != "GoSix Library Testing Service Updated" {
		t.Fatal("Updating Service Rule Failed")
	}
	// Add Route
	if err := a.Service.Create(id, s); err != nil {
		t.Fatal(err)
	}
	// Add Plugin
	ba := new(models.BasicAuth)
	ba.Username = "A"
	ba.Password = "B"
	if err := a.Service.AddPlugin(id, ba); err != nil {
		t.Fatal(err)
	}
	baG, err := a.Service.GetPlugin(id, models.BasicAuth{})
	if err != nil {
		t.Fatal(err)
	}
	baGet := baG.(*models.BasicAuth)
	if baGet.Username != ba.Username || baGet.Password != ba.Password {
		t.Fatal("Add/Get Plugin Failed")
	}
	// Cleanup
	if err := a.Service.Delete(id); err != nil {
		t.Fatal(err)
	}
}

// Test SSL
func TestSSL(t *testing.T) {
	// Create APISix Object
	a, _, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	// Create objects
	id := "GoSixTestID"
	var s models.SSL
	s.Cert = os.Getenv("GOSIX_APISIX_SSL_CERT_1")
	s.Key = os.Getenv("GOSIX_APISIX_SSL_KEY_1")
	var SNIs []string
	SNIs = append(SNIs, "test.com")
	s.SNIs = SNIs

	// Call Test
	v, _, err := lib.Test(id, s, models.SSL{}, a.SSL, a)
	if err != nil {
		t.Fatal(err)
	}
	// Confirm Data matches
	if v.SNIs[0] != "test.com" {
		t.Fatal("Updating Service Rule Failed")
	}
}

/*
// Enable this if you use Streams on your APISIX Server, its not enabled by default so this is commented out
//
// Test Stream
func TestStream(t *testing.T) {
	// Create APISix Object
	a, _, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	// Create objects
	id := "GoSixTestID"
	desc := "GoSix Library Testing Route"
	descU := "GoSix Library Testing Route Updated"
	addr := "127.0.0.0/8"
	addrU := "127.0.0.1/8"
	var value models.Stream
	var updated models.Stream
	value.Desc = desc
	value.RemoteAddr = addr
	var u models.Upstream
	m := make(map[string]int)
	m["127.0.0.1:1980"] = 1
	u.Nodes = m
	u.Name = "Test Upstream"
	u.Type = "roundrobin"
	value.Upstream = &u
	updated = value
	updated.Desc = descU
	updated.RemoteAddr = addrU

	// Call Test
	v, e, err := lib.Test(id, value, updated, a.Stream, a)
	if err != nil {
		t.Fatal(err)
	}
	// Confirm Data matches
	if v.Desc != desc && e.Desc != descU && v.RemoteAddr != addr && e.RemoteAddr != addrU {
		t.Fatal("Updating Route Rule Failed")
	}
}
*/
// Test Upstream
func TestUpstream(t *testing.T) {
	// Create APISix Object
	a, _, err := lib.CreateAndGetConsumer()
	if err != nil {
		t.Fatal(err)
	}
	// Create objects
	id := "GoSixTestID"
	var u models.Upstream
	var uU models.Upstream
	u.Desc = "GoSix Library Testing"
	u.Name = "Testing"
	m := make(map[string]int)
	m["127.0.0.1:1980"] = 1
	u.Nodes = m
	u.Name = "Test Upstream"
	u.Type = "roundrobin"
	uU = u
	uU.Desc = "GoSix Library Testing Updated"
	uU.Name = "Testing Updated"

	// Call Test
	v, e, err := lib.Test(id, uU, uU, a.Upstream, a)
	if err != nil {
		t.Fatal(err)
	}
	// Confirm Data matches
	if v.Name != "Testing" && v.Desc != "GoSix Library Testing" && e.Name != "Testing Updated" && e.Desc != "GoSix Library Testing Updated" {
		t.Fatal("Updating Failed")
	}
}

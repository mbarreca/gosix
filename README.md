# gosix

<div align="center">

A client library for APISIX written in Go.

[![GoDoc][doc-img]][doc]

<div align="left">

## Installation

<div align="center">

`go get github.com/mbarreca/gosix`

<div align="left">

Gosix was built and tested with Go 1.23, it may still work with prior versions however it has not been tested so use at your own risk.

## Why would I use this library?

This was built for a project that I'm working on that leverages the wonderful open source project by Apache called APISIX. API Gateways provide a lot of utility, everything from intelligent middleware to offering a comprehensive selection of authentication methods. APISIX stands out as its not paywalled and offers a tremendous amount of functionality out of the box. Support for things like client libraries make the development experience easier and this is my little contribution to the project.

## Scope

This library supports interfacing with the APISIX Admin API, all endpoints are implemented, plugins are still a work in progress. There is strong test coverage. See the Roadmap for more information of whats built and whats coming.

## Core Concepts

This library was built to be as extensible, performant and practical as possible. In a production environment, things like telemetry and performance are critical so it was built in from day one.

### OpenTelemetry

Opentelemetry is supported in this library via *context*. We use embedded structs to keep the same HTTPClient so that context can be passed throughout the process (i.e - Consumer created, added to whitelist, added to blacklist) to allow for rich tracing and performance monitoring. If you want to enable it, when you're instantiating the APISix object, instantiate OTel first, then pass true and the proper context through.

## Testing
`
go test -v ./test/tests
`

## Example
```

package main

import (
	"context"

	"github.com/mbarreca/gosix"
	"github.com/mbarreca/gosix/models"
)

func main() {
	// Create APISix Client
	gosix, err := gosix.New(context.Background(), false)
	if err != nil {
		panic(err)
	}
	// Add Consumer
	username := "MyUsername"
	if err := gosix.Consumer.Create(username, "MyDescription", nil); err != nil {
		panic(err)
	}

	// Check to see it was created properly
	_, err = gosix.Consumer.Get(username)
	if err != nil {
		panic(err)
	}

	// Add Basic Auth to Consumer
	basicUsername := "username@username.com"
	basicPassword := "password"
	if err = gosix.Basic.Add(username, basicUsername, basicPassword); err != nil {
		panic(err)
	}

	// Add a Route
	var route models.Route
	// Add Basic Details
	route.Name = "Test Route"
	route.Desc = "Testing Route"
	// Add Methods Supported
	methods := make([]string, 0)
	methods = append(methods, "GET", "POST")
	route.Methods = methods
	// Add the upstream to route this to
	var upstream models.Upstream
	m := make(map[string]int)
	m["127.0.0.1:1980"] = 1
	upstream.Nodes = m
	upstream.Name = "Test Upstream"
	upstream.Type = "roundrobin"
	route.Upstream = &upstream
	if err = gosix.Route.Create("", route); err != nil {
		panic(err)
	}
}

```

## Roadmap

### Implemented

*Consumer*
*Credential*
*Global Rule*
*Plugins*
- Basic, JWT, Key, Consumer Restriction
*Plugin Config*
*Proto*
*Route*
*Secret*
*Service*
*SSL*
*Stream*
*Upstream*

### Planned

*Additional Plugin Support*
- Additional Plugin Support
- Schema Validation

### Out of Scope

*Control API*
*Plugin Meta*

## How to use this library

You'll need to set the following environment variables in order to provide the correct values to the system. Please do not hard code them in, that's bad practice. Here's a list of what needs to be provided with examples.

```
// The address, inclusive of the port of your APISIX Instance
GOSIX_APISIX_ADDRESS="https://apisix.matteobarreca.com:8080

// Your Admin API KEY
GOSIX_APISIX_ADMIN_API_KEY="supersecretkey"

// Key Auth Plugin - Key Length - Default 100
GOSIX_APISIX_PLUGIN_KEY_LENGTH="100"

// REQUIRED - JWT Auth Plugin - Secret - Minimum 25 Characters
GOSIX_APISIX_PLUGIN_JWT_SECRET="supersecret"

// REQUIRED - JWT Auth Plugin - Algorithm Type - Must be HS256, HS512, RS256 or ES256
GOSIX_APISIX_PLUGIN_JWT_ALGORITHM="HS256"

// JWT Auth Plugin - Expiry Time in Seconds - Default 86400
GOSIX_APISIX_PLUGIN_JWT_EXP="86400"
```

## License

This is licensed under the Apache 2.0 Library to match its partner in crime, APISIX.

[doc]: https://pkg.go.dev/github.com/mbarreca/gosix
[doc-img]: https://pkg.go.dev/badge/github.com/mbarreca/gosix

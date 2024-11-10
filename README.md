# gosix

<div align="center">

A client library for APISIX written in Go.

<div align="left">

## Installation

<div align="center">

`go get github.com/mbarreca/gosix`

<div align="left">

Gosix was built and tested with Go 1.23, it may still work with prior versions however it has not been tested so use at your own risk.

## Why would I use this library?

This was built for a project that I'm working on that leverages the wonderful open source project by Apache called APISIX. API Gateways provide a lot of utility, everything from intelligent middleware to offering a comprehensive selection of authentication methods. APISIX stands out as its not paywalled and offers a tremendous amount of functionality out of the box. Support for things like client libraries make the development experience easier and this is my little contribution to the project.

## Scope

This library supports interfacing with the APISIX Admin API, specifically implementing the Route Endpoints

## Core Concepts

This library was built to be as extensible, performant and practical as possible. In a production environment, things like telemetry and performance are critical so it was built in from day one.

### OpenTelemetry

Opentelemetry is supported in this library via *context*. We use embedded structs to keep the same HTTPClient so that context can be passed throughout the process (i.e - Consumer created, added to whitelist, added to blacklist) to allow for rich tracing and performance monitoring.

## Testing
`
go test ./consumer/tests/*
`

## Roadmap

### Implemented

*Consumer*
- All Endpoint Abrstractions
- Plugin Support for Basic*, Key and JWT*

### Planned

- Additional Auth Plugins
- Routes
- Service
- Credential
- Upstream
- SSL
- Global Rule
- Plugin Config
- Plugin Metadata
- Plugin
- Stream Route
- Secret
- Proto
- Schema Validation

### Out of Scope

*Control API*


## How to use this library

You'll need to set the following environment variables in order to provide the correct values to the system. Please do not hard code them in, that's bad practice. Here's a list of what needs to be provided with examples.

`
// The address, inclusive of the port of your APISIX Instance
GOSIX_APISIX_ADDRESS="https://apisix.matteobarreca.com:8080
// Your Admin API KEY
GOSIX_APISIX_ADMIN_API_KEY="supersecretkey"
// Key Auth Plugin - Key Length - Default 100
GOSIX_APISIX_PLUGIN_KEY_LENGTH="100"
`


## License

This is licensed under the Apache 2.0 Library to match its partner in crime, APISIX.

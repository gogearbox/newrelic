<p align="center">
	<a href="https://gogearbox.com">
    	<img src="https://raw.githubusercontent.com/gogearbox/gearbox/master/assets/gearbox-512.png"/>
	</a>
    <br />
    <a href="https://godoc.org/github.com/gogearbox/netadaptor">
      <img src="https://godoc.org/github.com/gogearbox/netadaptor?status.png" />
    </a>
    <img src="https://github.com/gogearbox/netadaptor/workflows/Test%20&%20Build/badge.svg?branch=master" />
    <a href="https://goreportcard.com/report/github.com/gogearbox/gearbox">
      <img src="https://goreportcard.com/badge/github.com/gogearbox/netadaptor" />
    </a>
	<a href="https://discord.com/invite/CT8my4R">
      <img src="https://img.shields.io/discord/716724372642988064?label=Discord&logo=discord">
  	</a>
    <a href="https://deepsource.io/gh/gogearbox/netadaptor/?ref=repository-badge" target="_blank">
      <img alt="DeepSource" title="DeepSource" src="https://static.deepsource.io/deepsource-badge-light-mini.svg">
    </a>
</p>

**newrelic** middleware


### Supported Go versions & installation

:gear: gearbox requires version `1.14` or higher of Go ([Download Go](https://golang.org/dl/))

Just use [go get](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) to download and install gearbox

```bash
go get -u github.com/gogearbox/gearbox
go get u- github.com/gogearbox/newrelic
```


### Examples

```go
package main

import (
	  "github.com/newrelic/go-agent/v3/newrelic"
	  "github.com/gogearbox/gearbox"
    newrelicmiddleware "github.com/gogearbox/newrelic"
)

func main() {
	// Setup gearbox
	gb := gearbox.New()

	// Initialize newrelic
	nr, _ := newrelic.NewApplication(
		newrelic.ConfigAppName(APP_NAME),
		newrelic.ConfigLicense(LICENSE),
	)

	// Register the newrelic middleware for all requests
	gb.Use(newrelicmiddleware.New(nr))

	// Define your handler
	gb.Post("/hello", func(ctx gearbox.Context) {
		panic("There is an issue")
	})

	// Start service
	gb.Start(":3000")
}

```

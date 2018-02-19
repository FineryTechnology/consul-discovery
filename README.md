# Service discovery for Amaiz project

Will help to store configs and service monitoring.
Based on Consul: https://www.consul.io

## Prerequisites

You need to install local consul.
By default consul if located on `http://localhost:8500`
If you want to use custom URL you can modify env variable `AMAIZ_CONSUL_URL`:

### Set ENV variable:

```
export AMAIZ_CONSUL_URL="http://localhost:32768"
```

### Or change ENV on launch:

```
AMAIZ_CONSUL_URL="http://localhost:32768" go run main.go
```

## Usage

```go

package main

import (
	"fmt"

	discovery "github.com/FineryTechnology/amaiz-discovery"
)

type Config struct {
	ID   string
	Name string
	Host struct {
		Host string `json:"host"`
	}
}

func main() {
  var config Config
  err := discovery.New("/broker_gateway/v1/queue", &config)
  if err != nil {
    ...
  }
  fmt.Printf("Config: %+v", config)
}

```

## Updating KV storage (config)

Every consul has its own UI:

http://localhost:8500/ui/

You can update config on Key/Value section (link on top menu)

### Dev-server

Only by ssh-tunnel

### Production-server

Only by ssh-tunnel
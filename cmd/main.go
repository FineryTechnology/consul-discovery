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
	fmt.Printf("Config error: %+v", err)
	fmt.Printf("Config: %+v", config)
}

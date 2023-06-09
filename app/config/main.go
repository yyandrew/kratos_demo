package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"log"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "conf", "config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	c := config.New(config.WithSource(file.NewSource(flagconf)))
	if err := c.Load(); err != nil {
		panic(err)
	}

	var v struct {
		Service struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:service`
	}

	if err := c.Scan(&v); err != nil {
		panic(err)
	}
	log.Printf("config: %+v", v)

	name, err := c.Value("service.name").String()
	if err != nil {
		panic(err)
	}
	log.Printf("Service: %s", name)

	if err := c.Watch("service.name", func(key string, value config.Value) {
		log.Printf("config changed: %s = %v\n", key, value)
	}); err != nil {
		panic(err)
	}
	<-make(chan struct{})
}

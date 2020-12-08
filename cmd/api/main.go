package main

import (
	"log"
	"os"

	"github.com/sorokinkir/my-service-blueprint/config"
)

func main() {
	cfgPath, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
}

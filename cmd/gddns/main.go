package main

import (
	"time"

	"github.com/tserkov/google-ddns-client/internal/config"
	"github.com/tserkov/google-ddns-client/internal/log"
	"github.com/tserkov/google-ddns-client/pkg/client"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Error.Fatalln(err)
	}

	c := client.Client{
		Verbose: cfg.Verbose,
	}

	// Update all records on start
	for _, r := range cfg.Records {
		if err := c.UpdateRecord(r); err != nil {
			log.Error.Println(err)
		}
	}

	// Update every hour
	t := time.Tick(1 * time.Hour)
	for range t {
		for _, r := range cfg.Records {
			if err := c.UpdateRecord(r); err != nil {
				log.Error.Println(err)
			}
		}
	}
}

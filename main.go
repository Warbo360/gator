package main

import (
	"github.com/Warbo360/gator/internal/config"
	"log"
	"fmt"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
    	log.Fatalf("ERROR: Failed to read config file: %v", err)
	}

	err = cfg.SetUser("Walter")
	if err != nil {
		log.Fatalf("ERROR: Failed to set username: %v", err)
	}
	
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("ERROR: Failed to read config files: %v", err)
	}

    fmt.Printf("%+v\n", cfg)

}

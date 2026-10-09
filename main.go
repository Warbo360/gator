package main

import (
	"github.com/Warbo360/gator/internal/config"
	"log"
	"os"
)

func main() {

	currentCfg, err := config.Read()
	if err != nil {
    	log.Fatalf("ERROR: Failed to read config file: %v", err)
	}

	currentState := state{
		cfg: &currentCfg,
	}

	currentCommands := commands{cmds: make(map[string]func(*state, command) error)}
	currentCommands.register("login", handlerLogin)

	userArgs := os.Args
	if len(userArgs) < 2 {
    	log.Fatalf("ERROR: Not enough arguments passed")
	}

	passedCommand := command{
		name: userArgs[1],
		args: userArgs[2:],
	}

	err = currentCommands.run(&currentState, passedCommand)
	if err != nil {
		log.Fatalf("ERROR: Failed to run passed command: %v", err)
	}
}

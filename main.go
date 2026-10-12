package main

import (
	"github.com/Warbo360/gator/internal/config"
	"github.com/Warbo360/gator/internal/database"
	"log"
	"os"
	"database/sql"
)

import _ "github.com/lib/pq"

func main() {

	currentCfg, err := config.Read()
	if err != nil {
    	log.Fatalf("ERROR: Failed to read config file: %v", err)
	}

	db, err := sql.Open("postgres", currentCfg.DBURL)
	if err != nil {
		log.Fatalf("ERROR: Failed to open database connecition: %v", err)
	}
	dbQueries := database.New(db)

	currentState := state{
		db: dbQueries,
		cfg: &currentCfg,
	}

	currentCommands := commands{cmds: make(map[string]func(*state, command) error)}
	currentCommands.register("login", handlerLogin)
	currentCommands.register("register", handlerRegister)

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

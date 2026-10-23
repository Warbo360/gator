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

	// Gets the config file from users home dir

	currentCfg, err := config.Read()
	if err != nil {
    	log.Fatalf("[ERROR]: Failed to read config file - \"%v\"", err)
	}

	// Connects to the database

	db, err := sql.Open("postgres", currentCfg.DBURL)
	if err != nil {
		log.Fatalf("[ERROR]: Failed to open database connecition - \"%v\"", err)
	}
	dbQueries := database.New(db)

	// Creates a state variable that holds the config and database queries

	currentState := state{
		db: dbQueries,
		cfg: &currentCfg,
	}

	// CLI Command Registrations

	currentCommands := commands{cmds: make(map[string]func(*state, command) error)}
	currentCommands.register("login", handlerLogin)
	currentCommands.register("register", handlerRegister)
	currentCommands.register("reset", handlerReset)
	currentCommands.register("users", handlerUsers)
	currentCommands.register("agg", handlerAgg)
	currentCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	currentCommands.register("feeds", handlerFeeds)
	currentCommands.register("follow", middlewareLoggedIn(handlerFollow))
	currentCommands.register("following", middlewareLoggedIn(handlerFollowing))
	currentCommands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	currentCommands.register("browse", middlewareLoggedIn(handlerBrowse))

	// Checks for passed CLI args for an actual command

	userArgs := os.Args
	if len(userArgs) < 2 {
    	log.Fatalf("[ERROR]: No command given")
	}

	// Splits the CLI args into command name and then args for that command

	passedCommand := command{
		name: userArgs[1],
		args: userArgs[2:],
	}

	// Passes the command struct to the run command which is then handled by whatever handler from the above registered
	// commands

	err = currentCommands.run(&currentState, passedCommand)
	if err != nil {
		log.Fatalf("[ERROR]: Failed to run passed command - \"%v\"", err)
	}
}

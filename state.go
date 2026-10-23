package main

import (
	"github.com/Warbo360/gator/internal/config"
	"github.com/Warbo360/gator/internal/database"
)

// State struct which hold the config pointer for the database connection and queries as well as config functionality
// for reading and writing the json config file which holds the currently set user and database connection string

type state struct {
	db *database.Queries
	cfg *config.Config
}

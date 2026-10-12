package main

import (
	"github.com/Warbo360/gator/internal/config"
	"github.com/Warbo360/gator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

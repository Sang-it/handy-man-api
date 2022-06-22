package main

import (
	"log"
	"os"

	"github.com/Sang-it/handy-man-api/internal-new/adapters/framework/left/server"
	"github.com/Sang-it/handy-man-api/internal-new/adapters/framework/right/db"
	"github.com/Sang-it/handy-man-api/internal-new/ports"
	// "github.com/Sang-it/handy-man-api/internal/adapters/api"
)

func main() {

	var err error

	// PORTS
	var dbaseAdapter ports.DatabasePort
	var apiAdapter ports.ServerPort

	// ENV Variables
	var connection_string string = os.Getenv("CONNECTION_STRING")
	var port string = os.Getenv("PORT")

	// CONNECTING PORTS TO THE RESPECTIVE ADAPTERS
	dbaseAdapter, err = db.NewAdapter(connection_string)
	if err != nil {
		log.Fatalf("Failed to Initiate DB Connection: %v", err)
	}

	apiAdapter = server.NewAdapter(dbaseAdapter, port)

	apiAdapter.Start()

}

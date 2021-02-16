package main

import (
	"fmt"
	"log"

	"github.com/stevenkl/wkr/pkg/config"
	"github.com/stevenkl/wkr/pkg/database"
	"github.com/stevenkl/wkr/pkg/webserver"
)

func main() {
	config := config.New()
	err := config.Parse("wkr.config")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)

	app := webserver.New(&config)
	if dberr := database.Init(&config); dberr != nil {
		log.Fatal(dberr)
	}

	serverAddress := fmt.Sprintf("%s:%d", config.Host, config.Port)
	fmt.Println(serverAddress)

	app.Listen(serverAddress)
}

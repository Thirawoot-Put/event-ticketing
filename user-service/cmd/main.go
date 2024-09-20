package main

import (
	"fmt"

	"github.com/Thirawoot-Put/event-ticketing/user-service/internal/config"
	"github.com/Thirawoot-Put/event-ticketing/user-service/internal/infras/server"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		panic("Error loading .env file, falling back to system environment variables")
	}

	cfg := config.Config{}

	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%v", err)
		panic("error parse rnv file")
	}

	app := server.AppServer()
	app.Start(cfg.Port)
}

package main

import (
	"fmt"
	"go-technical/api/database"
	"go-technical/api/route"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

const (
	ENVIRONMENT_FILE_PATH = ".env"
	SERVER_PORT           = "1323"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to run server")
	}
}

func run() error {
	err := godotenv.Load(ENVIRONMENT_FILE_PATH)
	if err != nil {
		log.Fatal().Msgf("unable to load .env file: %v", err)
	}

	db, err := database.NewPostgresDatabase()
	if err != nil {
		return err
	}
	err = db.Migrate()
	if err != nil {
		return err
	}
	defer db.Close()

	router := route.NewEchoRouter(db)
	router.RegisterRoutes()

	return router.Start(fmt.Sprintf(":%v", SERVER_PORT))
}

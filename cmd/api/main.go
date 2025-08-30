package main

import (
	"log"

	"github.com/AryanMishra09/Social/internal/env"
	"github.com/AryanMishra09/Social/internal/store"
	"github.com/joho/godotenv"
)

func main() {

	//Load env:
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8000"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}

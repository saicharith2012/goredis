package main

import (
	"log"

	"github.com/saicharith2012/goredis/internal/config"
	"github.com/saicharith2012/goredis/internal/server"
	"github.com/saicharith2012/goredis/internal/store"
)

func main() {
	cfg := config.Default()
	store := store.NewSharedState()
	srv := server.New(cfg.Port, store)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/saicharith2012/goredis/internal/config"
	"github.com/saicharith2012/goredis/internal/server"
)

func main() {
	cfg := config.Default()
	srv := server.New(cfg.Port)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

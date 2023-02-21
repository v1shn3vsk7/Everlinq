package main

import (
	"context"
	"github.com/v1shn3vsk7/Everlinq/internal/config"
	"github.com/v1shn3vsk7/Everlinq/internal/server"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	cfg := config.NewConfig(&ctx)

	s := server.NewServer(cfg)

	if err := s.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

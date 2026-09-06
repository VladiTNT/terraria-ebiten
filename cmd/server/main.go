package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/VladiTNT/terraria-ebiten/internal/server"
	"github.com/VladiTNT/terraria-ebiten/internal/server/config"
)

func main() {
	cfg := config.Default()
	srv := server.NewServer(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := srv.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error shutting down server: %v\n", err)
		os.Exit(1)
	}
}

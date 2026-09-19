package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/VladiTNT/terraria-ebiten/internal/server/config"
)

type Server struct {
	Config *config.Config
	Game   *Game

	Ln net.Listener
	Wg sync.WaitGroup
}

func NewServer(cfg *config.Config) *Server {
	ln, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)))
	if err != nil {
		panic(err)
	}

	return &Server{
		Game:   NewGame(cfg.GameSettings),
		Config: cfg,
		Ln:     ln,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// Listener goroutine
	go func() {
		fmt.Printf("Terraria server listening on address: %s\n", s.Ln.Addr().String())
		for {
			conn, err := s.Ln.Accept()
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					fmt.Println("Listener closed.")
					break
				}

				fmt.Printf("Error accetpting new TCP conn: %v\n", err)
				continue
			}

			// Add new player handler to waitgroup
			s.Wg.Go(func() { s.handleConn(conn) })
		}
	}()

	// Start game loop in a separate goroutine
	go s.Game.Main()

	// Waiting for context to initiate shutdown sequence.
	<-ctx.Done()
	fmt.Println("Shutdown signal received, shutting down...")

	err := s.Ln.Close()
	if err != nil {
		return err
	}

	wait := make(chan struct{})

	shutdownCtx, cancel := context.WithTimeout(context.Background(), s.Config.ShutdownTime)
	defer cancel()

	go func() {
		s.Wg.Wait()
		close(wait)
	}()

	// Server blocks, waiting if the players quit, otherwise it shuts down on it's own after a while.
	select {
	case <-shutdownCtx.Done():
		fmt.Println("Server forced to shutdown, players didn't leave.")
		return shutdownCtx.Err()
	case <-wait:
		return nil
	}
}

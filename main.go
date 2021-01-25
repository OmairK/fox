package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/OmairK/fox/internal/server"
)

func main() {
	service := ":8000"
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	srv := server.NewServer(service)

	select {
	case <-sig:
		srv.Stop()
	}

}

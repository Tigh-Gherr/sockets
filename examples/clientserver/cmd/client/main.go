package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/theflyingcodr/sockets"
	"github.com/theflyingcodr/sockets/examples/clientserver"
)

type TestMessage struct {
	When time.Time `json:"when"`
	Test string    `json:"test"`
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	c := clientserver.SetupClient()
	defer c.Close()
	c.WithMiddleware(sockets.PanicHandler, sockets.Timeout)
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info().Msg("closing client")
}

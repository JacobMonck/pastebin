package main

import "github.com/jacobmonck/pastebin/services/api/internal/api"

func main() {
	app := api.New()

	go app.Start()

	// This gets deferred until an interrupt signal is sent.
	app.Shutdown()
}

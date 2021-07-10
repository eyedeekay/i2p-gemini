// This example illustrates a Gemini server using I2P

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"git.sr.ht/~adnano/go-gemini"
	"git.sr.ht/~adnano/go-gemini/certificate"

	"github.com/eyedeekay/sam3/helper"
)

func main() {
	certificates := &certificate.Store{}
	certificates.Register("localhost")
	if err := certificates.Load("var/lib/gemini/certs"); err != nil {
		log.Fatal(err)
	}

	mux := &gemini.Mux{}
	mux.Handle("/", gemini.FileServer(os.DirFS("var/www")))

	server := &gemini.Server{
		Handler:        gemini.LoggingMiddleware(mux),
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   1 * time.Minute,
		GetCertificate: certificates.Get,
	}

	listener, err := sam.I2PListener("gemini","127.0.0.1:7656","gemini")
	if err != nil {
		log.Fatal(err)
	}

	// Listen for interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	errch := make(chan error)
	go func() {
		ctx := context.Background()
		errch <- server.Serve(ctx, listener)
	}()

	select {
	case err := <-errch:
		log.Fatal(err)
	case <-c:
		// Shutdown the server
		log.Println("Shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
}

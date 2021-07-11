// This example illustrates a Gemini server using I2P

package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"git.sr.ht/~adnano/go-gemini"
	"git.sr.ht/~adnano/go-gemini/certificate"

	"github.com/eyedeekay/sam3/helper"
	"github.com/eyedeekay/sam3/i2pkeys"
)

var (
	certdir  = flag.String("certs", "var/lib/gemini/certs", "Directory where server certificates(TLS) will be stored")
	filesdir = flag.String("files", "www", "Directory of files to serve up with Gemini")
	name     = flag.String("name", "i2pgemini", "Name of the service to pass to SAM")
	samaddr  = flag.String("sam", "127.0.0.1:7656", "SAM API to connect to and user")
)

func main() {
	flag.Parse()
	os.MkdirAll(*certdir, 0755)
	listener, err := sam.I2PListener(*name, *samaddr, filepath.Join(*certdir, "gemini"))
	if err != nil {
		log.Fatal(err)
	}
	certificates := &certificate.Store{}
	certificates.Register(listener.Addr().String())
	if err := certificates.Load(*certdir); err != nil {
		log.Fatal(err)
	}
	log.Println("gemini://" + listener.Addr().(i2pkeys.I2PAddr).Base32())

	mux := &gemini.Mux{}
	mux.Handle("/", gemini.FileServer(os.DirFS(*filesdir)))

	server := &gemini.Server{
		Handler:        gemini.LoggingMiddleware(mux),
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   1 * time.Minute,
		GetCertificate: certificates.Get,
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

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"postman/internal/config"
	"postman/internal/storage"
	"time"

	postman "postman"
)

func main() {
	var wait time.Duration

	config := config.LoadConfiguration("/Users/niranjan.prajapati/practise/postman/properties/local.json")
	storage := storage.InitStorage(config.DataBase)
	server := postman.NewServer(config, storage)

	log.Printf("starting server at %s", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other servicesd
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

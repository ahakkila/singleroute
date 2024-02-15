package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ahakkila/singleroute/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/catchall/person", handlers.GetPerson)
	r.GET("/catchall/person/:id", handlers.GetPersonByID)
	r.POST("/catchall/person", handlers.CreatePerson)
	r.GET("/catchall/account", handlers.GetAccount)
	r.GET("/catchall/account/:id", handlers.GetAccountByID)
	r.POST("/catchall/account", handlers.CreateAccount)

	srv := &http.Server{
		Addr:         resolvePort(),
		Handler:      r,
		IdleTimeout:  90 * time.Second,
		WriteTimeout: 90 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errstr := fmt.Sprintf("HTTP listener exited: %v", err)
			log.Println("error:", errstr)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	log.Println("Server exiting")

}

func resolvePort() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	
	"github.com/spf13/viper"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
}

func run() error {
	// Initialize viper
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	
	viper.SetDefault("APP_SERVER_PORT", 8000)
	
	srv := &http.Server{
		Addr: ":" + viper.GetString("APP_SERVER_PORT"),
	}

	serverErr := make(chan error, 1)
	
	go func() {
		log.Println("starting server...")
		serverErr <- srv.ListenAndServe()
	}()
	
	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %w", err)
		
	// TODO: Add support for graceful server shutdown if time permits
	}
}

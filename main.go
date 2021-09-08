package main

import (
	"fmt"
	"github.com/bwheel181/uniswap181/routers"
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
	viper.SetDefault("UNISWAP_QUERY_URL_V3", "https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-alt")

	log.Println("setting up routes")
	uniswapQueryURL := viper.GetString("UNISWAP_QUERY_URL_V3")
	router := routers.Router(uniswapQueryURL)

	srv := &http.Server{
		Addr:    ":" + viper.GetString("APP_SERVER_PORT"),
		Handler: router,
	}
	serverErr := make(chan error, 1)

	go func() {
		log.Println("starting server on port" + srv.Addr)
		serverErr <- srv.ListenAndServe()
	}()

	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %w", err)

		// TODO: Add support for graceful server shutdown if time permits
	}
}

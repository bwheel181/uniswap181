package routers

import (
	"fmt"
	controllerV1 "github.com/bwheel181/uniswap181/controllers/api/v1"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
)

func Router(queryURL string) *mux.Router {
	//router := mux.NewRouter().PathPrefix("/find-a-store").Subrouter()
	fmt.Println(queryURL)
	gac := generateV1Controllers()
	router := mux.NewRouter().PathPrefix("/uniswarm").Subrouter()
	router.NotFoundHandler = http.HandlerFunc(controllerV1.NotFoundHandler)
	
	router.
		Path("/assets").
		Methods(http.MethodGet).
		Handler(http.HandlerFunc(gac.GetAssetByID))
	
	return router
}

func generateV1Controllers() controllerV1.GraphAssetController {
	return controllerV1.NewGraphAssetController(viper.GetString("UNISWAP_QUERY_URL_V3"))
}

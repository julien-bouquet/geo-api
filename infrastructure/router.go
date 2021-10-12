package infrastructure

import (
	"log"
	"net/http"
	"os"

	"github.com/geo-api/interfaces"
	"github.com/gorilla/mux"
)

func setSubRouterApi(router *mux.Router, pointControler interfaces.PointController) {
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/point", pointControler.List).Methods(http.MethodGet)
	//api.HandleFunc("/point", pointControler.Add).Methods(http.MethodPost)
}

func HandleRequest(noSQLHander interfaces.NoSQLHandler) {
	pointControler := interfaces.NewPointController(noSQLHander)

	router := mux.NewRouter()
	setSubRouterApi(router, *pointControler)

	// Set PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening and Serve on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

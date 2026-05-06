package server

import (
	"log"
	"net/http"
	"os"

	"sprint-final/pkg/api"
)

func Run() {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	mux := http.NewServeMux()

	api.InitMux(mux)

	webDir := "web"

	fs := http.FileServer(http.Dir(webDir))

	mux.Handle("/web/", http.StripPrefix("/web/", fs))

	mux.Handle("/", fs)

	log.Println("Server starting on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

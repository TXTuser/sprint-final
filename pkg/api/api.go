package api

import "net/http"

func InitMux(mux *http.ServeMux) {
	mux.HandleFunc("/api/task", taskHandler)
	mux.HandleFunc("/api/task/", taskHandler)
	mux.HandleFunc("/api/tasks", tasksHandler)
	mux.HandleFunc("/api/nextdate", nextDateHandler)
	mux.HandleFunc("/api/task/done", taskDoneHandler)
}

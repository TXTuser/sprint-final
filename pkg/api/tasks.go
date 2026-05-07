package api

import (
	"net/http"
	"sprint-final/pkg/db"
	"strconv"
)

const tasksLimit = 50

type TaskResponse struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJSON(w, map[string]string{"error": "method not allowed"})
		return
	}

	tasks, err := db.Tasks(tasksLimit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	if tasks == nil {
		tasks = []db.Task{}
	}

	response := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		response[i] = TaskResponse{
			ID:      strconv.FormatInt(task.ID, 10),
			Date:    task.Date,
			Title:   task.Title,
			Comment: task.Comment,
			Repeat:  task.Repeat,
		}
	}

	writeJSON(w, map[string]interface{}{
		"tasks": response,
	})
}

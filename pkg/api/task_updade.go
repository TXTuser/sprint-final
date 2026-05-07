package api

import (
	"encoding/json"
	"net/http"
	"sprint-final/pkg/db"
	"strconv"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var bodyMap map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&bodyMap); err != nil {
		w.WriteHeader(http.StatusBadRequest) 
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	var task db.Task

	if idVal, ok := bodyMap["id"]; ok {
		switch v := idVal.(type) {
		case string:
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				writeJSON(w, map[string]string{"error": "Неверный идентификатор"})
				return
			}
			task.ID = id
		case float64:
			task.ID = int64(v)
		case int64:
			task.ID = v
		default:
			w.WriteHeader(http.StatusBadRequest)
			writeJSON(w, map[string]string{"error": "Неверный тип идентификатора"})
			return
		}
	}

	if dateVal, ok := bodyMap["date"]; ok {
		if dateStr, ok := dateVal.(string); ok {
			task.Date = dateStr
		}
	}
	if titleVal, ok := bodyMap["title"]; ok {
		if titleStr, ok := titleVal.(string); ok {
			task.Title = titleStr
		}
	}
	if commentVal, ok := bodyMap["comment"]; ok {
		if commentStr, ok := commentVal.(string); ok {
			task.Comment = commentStr
		}
	}
	if repeatVal, ok := bodyMap["repeat"]; ok {
		if repeatStr, ok := repeatVal.(string); ok {
			task.Repeat = repeatStr
		}
	}

	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, map[string]string{"error": "Не указан идентификатор"})
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, map[string]string{"error": "empty title"})
		return
	}

	if err := checkDate(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, map[string]interface{}{})
}

package api

import (
	"net/http"
	"time"
)

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	if date == "" || repeat == "" {
		http.Error(w, "missing params", http.StatusBadRequest)
		return
	}

	now := time.Now()
	if nowStr != "" {
		t, err := time.Parse(DateFormat, nowStr)
		if err != nil {
			http.Error(w, "invalid now param", http.StatusBadRequest)
			return
		}
		now = t
	}

	next, err := NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if _, err := w.Write([]byte(next)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

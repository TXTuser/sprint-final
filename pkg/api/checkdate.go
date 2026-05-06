package api

import (
	"time"

	"sprint-final/pkg/db"
)

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(DateFormat)
	}

	t, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		return err
	}

	if t.Before(now.Truncate(24 * time.Hour)) {
		if task.Repeat == "" {
			task.Date = now.Format(DateFormat)
			return nil
		}

		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}

		task.Date = next
	}

	return nil
}

package db

import (
	"database/sql"
	"fmt"
)

type Task struct {
	ID      int64  `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func Tasks(limit int) ([]Task, error) {
	rows, err := DB.Query(`
		SELECT id, date, title, comment, repeat
		FROM scheduler
		ORDER BY date ASC
		LIMIT ?
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTask(id int64) (*Task, error) {
	t := &Task{}

	err := DB.QueryRow(`
		SELECT id, date, title, comment, repeat
		FROM scheduler
		WHERE id = ?
	`, id).Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Задача не найдена")
		}
		return nil, err
	}

	return t, nil
}

func AddTask(t *Task) (int64, error) {
	res, err := DB.Exec(`
		INSERT INTO scheduler(date, title, comment, repeat)
		VALUES (?, ?, ?, ?)
	`, t.Date, t.Title, t.Comment, t.Repeat)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func UpdateTask(t *Task) error {
	res, err := DB.Exec(`
		UPDATE scheduler
		SET date=?, title=?, comment=?, repeat=?
		WHERE id=?
	`, t.Date, t.Title, t.Comment, t.Repeat, t.ID)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("Задача не найдена")
	}

	return nil
}

func UpdateTaskDate(id int64, date string) error {
	res, err := DB.Exec(`
		UPDATE scheduler
		SET date=?
		WHERE id=?
	`, date, id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("Задача не найдена")
	}

	return nil
}

func DeleteTask(id int64) error {
	res, err := DB.Exec(`
		DELETE FROM scheduler
		WHERE id=?
	`, id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("Задача не найдена")
	}

	return nil
}

package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title TEXT NOT NULL DEFAULT "",
    comment TEXT,
    repeat VARCHAR(128) NOT NULL DEFAULT ""
);

CREATE INDEX IF NOT EXISTS idx_date ON scheduler(date);
`

func Init(dbFile string) error {
	var err error

	_, err = os.Stat(dbFile)
	needCreate := err != nil

	DB, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	if needCreate {
		if _, err = DB.Exec(schema); err != nil {
			DB.Close()
			return err
		}
	}

	return DB.Ping()
}

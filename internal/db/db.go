package db

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

func ConnectDB() *sql.DB {
	logrus.Info("Connecting database")
	appPath, err := os.Executable()
	if err != nil {
		logrus.Error("failed to return the path.", "err", err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), "kode.db")

	_, err = os.Stat(dbFile)
	var install bool
	if err != nil {
		install = true
	}

	db, err := sql.Open("sqlite", "kode.db")
	if err != nil {
		logrus.Error("failed to connect db.", "err", err)
		os.Exit(1)
	}

	create := `
	    CREATE TABLE notes(
			id INTEGER PRIMARY KEY,
			author_id INTEGER,
			title VARCHAR,
			text VARCHAR
		);
		CREATE INDEX author_idx ON notes (author_id);
		`

	if install {
		logrus.Info("Creating db with script `create`")
		if _, err := db.Exec(create); err != nil {
			logrus.Error("failed to create db: ", err)
			os.Exit(1)
		}
	}
	return db
}

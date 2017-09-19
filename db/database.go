package db

import (
	"fmt"
	"database/sql"
	 _ "github.com/go-sql-driver/mysql"
	 "github.com/Simulalex/haveyouconsidered/models"
)

const dbConnectionString = "%v:%v@%v"

type Database interface {
	GetSummary(name string) (*models.Technology, error)
}

type database struct{
	mysql *sql.DB
}

type Config struct {
	Username string
	Password string
	Database string
}
func NewDatabase(cfg *Config) (Database, error) {
	// "user:password@/dbname"
	db, err := sql.Open("mysql", fmt.Sprintf(dbConnectionString, cfg.Username, cfg.Password, cfg.Database))

	if err != nil {
		return nil, err
	}

	outgoing := database{db}
	return outgoing, nil
}

func (db database) GetSummary(name string) (*models.Technology, error) {
	var technologyId, authorId int
	var summary, username, title string

	err := db.mysql.QueryRow("SELECT technologies.id, title, summary, authors.id, authors.username FROM technologies JOIN authors ON authors.id = technologies.author_id WHERE technologies.name = ?", name).Scan(&technologyId, &title, &summary, &authorId, &username)
	if err != nil{ //sql.ErrNoRows
        return nil, err
	}

	author := models.Author{authorId, username}

	technology := models.Technology{
		technologyId,
		name,
		title,
		summary,
		&author,
	}

	return &technology, nil
}


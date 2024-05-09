package db

import (
	"database/sql"
	"log"
)

type Database interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

type DBService struct {
	DB Database
}

func New(db Database) DBService {
	return DBService{DB: db}
}

func (service DBService) GetNames() ([]string, error) {
	query := "SELECT name FROM users"

	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return names, err
}

func (service DBService) SelectUniqueValues(columnName string, tableName string) ([]string, error) {
	query := "SELECT DISTINCT " + columnName + " FROM " + tableName
	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var values []string

	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}
		values = append(values, value)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return values, err
}

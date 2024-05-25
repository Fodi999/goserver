package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS custom_data (
		"key" TEXT NOT NULL PRIMARY KEY, 
		"value" TEXT,
		"color" TEXT
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveCustomData(key, value, color string) error {
	insertSQL := `INSERT INTO custom_data(key, value, color) VALUES (?, ?, ?)`
	_, err := db.Exec(insertSQL, key, value, color)
	return err
}

func GetCustomData(key string) (string, string, string, error) {
	querySQL := `SELECT key, value, color FROM custom_data WHERE key = ?`
	row := db.QueryRow(querySQL, key)

	var value, color string
	err := row.Scan(&key, &value, &color)
	if err != nil {
		return "", "", "", err
	}

	return key, value, color, nil
}

func UpdateCustomData(key, value, color string) error {
	updateSQL := `UPDATE custom_data SET value = ?, color = ? WHERE key = ?`
	_, err := db.Exec(updateSQL, value, color, key)
	return err
}

func DeleteCustomData(key string) error {
	deleteSQL := `DELETE FROM custom_data WHERE key = ?`
	_, err := db.Exec(deleteSQL, key)
	return err
}

func GetAllCustomData() ([]map[string]string, error) {
	querySQL := `SELECT key, value, color FROM custom_data`
	rows, err := db.Query(querySQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []map[string]string
	for rows.Next() {
		var key, value, color string
		err := rows.Scan(&key, &value, &color)
		if err != nil {
			return nil, err
		}
		entry := map[string]string{
			"key":   key,
			"value": value,
			"color": color,
		}
		data = append(data, entry)
	}
	if len(data) == 0 {
		return []map[string]string{}, nil
	}
	return data, nil
}

package db

import (
	"database/sql"
	"errors"
	"notes/constants"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

func Auth(username string, password string) (int, error) {

	var id int
	db, err := sql.Open("mysql", constants.DbPath)
	if err != nil {
		return 0, errors.New("error with sql access")
	}
	defer db.Close()
	var hash_pass string
	row := db.QueryRow("SELECT password FROM user WHERE username = ?;", username)
	row.Scan(&hash_pass)
	if hash_pass == "" {
		return 0, errors.New("user do not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash_pass), []byte(password))
	if err != nil {
		return 0, errors.New("invalid password")
	}

	row = db.QueryRow("SELECT id FROM user WHERE username = ?;", username)
	row.Scan(&id)

	return id, nil
}

func Reg(username string, password string) (int, error) {
	db, err := sql.Open("mysql", constants.DbPath)
	if err != nil {
		return 0, errors.New("error with sql access")
	}
	defer db.Close()
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return 0, errors.New("error with password generation")
	}
	var id int
	_, err = db.Exec("INSERT INTO user (username, password) VALUES (?, ?);", username, pass)
	row := db.QueryRow("SELECT id FROM user WHERE username = ?;", username)
	row.Scan(&id)

	if err != nil {
		return 0, errors.New("user already exist")
	}

	return id, nil
}

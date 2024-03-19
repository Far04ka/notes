package db

import (
	"database/sql"
	"errors"
	"notes/constants"

	_ "github.com/go-sql-driver/mysql"
)

func GetNotes(usrId int) (*[]constants.Note, error) {
	db, err := sql.Open("mysql", constants.DbPath)
	if err != nil {
		return nil, errors.New("error with sql access")
	}
	defer db.Close()

	usr_notes := []constants.Note{}
	row, err := db.Query("SELECT id, header, date, text FROM notes WHERE userId = ?;", usrId)
	if err != nil {
		return nil, errors.New("error with sql access")
	}
	defer row.Close()

	for row.Next() {
		var header, date, text string
		var id int
		row.Scan(&id, &header, &date, &text)
		note := constants.Note{Text: text, Date: date, Header: header, UserId: usrId, Id: id}
		usr_notes = append(usr_notes, note)
	}

	return &usr_notes, nil
}

func NewNoteDb(header string, date string, text string, usrId int) error {
	db, err := sql.Open("mysql", constants.DbPath)
	if err != nil {
		return errors.New("error with sql access")
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO notes (header, date, text, userId) VALUES (?, ?, ?, ?)", header, date, text, usrId)
	if err != nil {
		return errors.New("error with data")
	}

	return nil
}

func GetNote(id int) (*constants.Note, error) {
	db, err := sql.Open("mysql", constants.DbPath)
	if err != nil {
		return nil, errors.New("error with sql access")
	}
	defer db.Close()

	row := db.QueryRow("SELECT header, date, text, userId FROM notes WHERE id = ?;", id)

	var header, date, text string
	var userId int
	err = row.Scan(&header, &date, &text, &userId)
	if err != nil {
		return nil, errors.New("error with sql access")
	}
	note := constants.Note{Text: text, Date: date, Header: header, UserId: userId, Id: id}
	return &note, nil
}

func DelNote(id int) error {
	db, err := sql.Open("mysql", constants.DbPath)
	if err != nil {
		return errors.New("error with sql access")
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM notes WHERE id = ?;", id)
	if err != nil {
		return errors.New("error with sql access")
	}
	return nil
}

func RedNote(id int, header string, text string, date string) error {
	db, err := sql.Open("mysql", constants.DbPath)
	if err != nil {
		return errors.New("error with sql access")
	}
	defer db.Close()

	_, err = db.Exec("UPDATE notes SET header = ?, text = ?, date = ? WHERE id = ?;", header, text, date, id)
	if err != nil {
		return errors.New("error with sql access")
	}
	return nil
}

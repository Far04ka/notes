package pages

import (
	"errors"
	"html/template"
	"net/http"
	"notes/db"
	"strconv"
	"time"
)

func NewNote(w http.ResponseWriter, r *http.Request) {
	cookie_id, _ := Check(w, r)
	usr_id, _ := strconv.Atoi(cookie_id)
	if r.Method == http.MethodGet {
		tmp, err := template.ParseFiles("./html/head_foot.html", "./html/new_note.html")
		if ErrMessage(w, err) {
			return
		}
		err = tmp.ExecuteTemplate(w, "head_foot.html", nil)
		if ErrMessage(w, err) {
			return
		}
	} else if r.Method == http.MethodPost {
		header := r.FormValue("header")
		text := r.FormValue("text")
		date := r.FormValue("date")
		tm, err := time.Parse("2006-01-02", date)
		if ErrMessage(w, err) {
			return
		}
		if tm.Before(time.Now()) {
			err = errors.New("дата до сегодняшней")
		} else {
			err = nil
		}
		if ErrMessage(w, err) {
			return
		}
		err = db.NewNoteDb(header, date, text, usr_id)
		if ErrMessage(w, err) {
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)

	} else {
		ErrMessage(w, errors.New("method not allowed"))
	}
}

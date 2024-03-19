package pages

import (
	"errors"
	"html/template"
	"net/http"
	"notes/db"
	"strconv"
)

func Note(w http.ResponseWriter, r *http.Request) {
	Check(w, r)
	if r.Method == http.MethodPost {
		form_id := r.FormValue("note_id")
		id, _ := strconv.Atoi(form_id)
		note, err := db.GetNote(id)
		if ErrMessage(w, err) {
			return
		}
		tmp, err := template.ParseFiles("./html/head_foot.html", "./html/note.html")
		if ErrMessage(w, err) {
			return
		}
		err = tmp.ExecuteTemplate(w, "head_foot.html", note)
		if ErrMessage(w, err) {
			return
		}
	} else {
		ErrMessage(w, errors.New("method not allowed"))
	}
}

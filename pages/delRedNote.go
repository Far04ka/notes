package pages

import (
	"html/template"
	"net/http"
	"notes/db"
	"strconv"
)

func DelRedNote(w http.ResponseWriter, r *http.Request) {
	Check(w, r)
	if r.Method == http.MethodGet {
		w.WriteHeader(404)
		return
	}
	act := r.FormValue("act")
	note_id := r.FormValue("note_id")
	id, _ := strconv.Atoi(note_id)
	if act == "del" {
		err := db.DelNote(id)
		if ErrMessage(w, err) {
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else if act == "red" {
		note, err := db.GetNote(id)
		if ErrMessage(w, err) {
			return
		}
		tmp, err := template.ParseFiles("./html/head_foot.html", "./html/red_note.html")
		if ErrMessage(w, err) {
			return
		}
		err = tmp.ExecuteTemplate(w, "head_foot.html", note)
		if ErrMessage(w, err) {
			return
		}
	} else if act == "apply" {
		header := r.FormValue("header")
		text := r.FormValue("text")
		date := r.FormValue("date")
		err := db.RedNote(id, header, text, date)
		if ErrMessage(w, err) {
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

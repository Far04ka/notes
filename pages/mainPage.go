package pages

import (
	"html/template"
	"net/http"
	"notes/constants"
	"notes/db"
	"strconv"
)

func Route(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

	cookie_id, name := Check(w, r)

	id, _ := strconv.Atoi(cookie_id)
	usr := constants.UsrInf{UserName: name, Id: id}
	listNotes, err := db.GetNotes(id)
	if ErrMessage(w, err) {
		return
	}
	usr.ListNotes = listNotes

	tmp, err := template.ParseFiles("./html/head_foot.html", "./html/header.html", "./html/notes.html")
	if ErrMessage(w, err) {
		return
	}
	err = tmp.ExecuteTemplate(w, "head_foot.html", usr)
	if ErrMessage(w, err) {
		return
	}

}

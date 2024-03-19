package pages

import (
	"html/template"
	"log"
	"net/http"
	"notes/constants"
	"notes/db"
	"strconv"
)

func ErrMessage(w http.ResponseWriter, e error) bool {
	if e != nil {
		tmp, err := template.ParseFiles("./html/head_foot.html", "./html/err.html")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Template error"))
		}
		err = tmp.ExecuteTemplate(w, "head_foot.html", e.Error())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Template error"))
		}
		log.Print(e)
		return true
	}
	return false
}

func Auth_form(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmp, err := template.ParseFiles("./html/head_foot.html", "./html/reg.html")
		if ErrMessage(w, err) {
			return
		}
		err = tmp.ExecuteTemplate(w, "head_foot.html", nil)
		if ErrMessage(w, err) {
			return
		}
	} else if r.Method == http.MethodPost {
		var id int
		var err error
		u_name := r.FormValue("username")
		pass := r.FormValue("password")
		usr := constants.UsrInf{UserName: u_name, Password: pass}
		err = usr.RegValidate()
		if ErrMessage(w, err) {
			return
		}

		act := r.FormValue("act")
		if act == "register" {
			id, err = db.Reg(u_name, pass)
			if ErrMessage(w, err) {
				return
			}

		} else if act == "login" {
			id, err = db.Auth(u_name, pass)
			if ErrMessage(w, err) {
				return
			}
		}
		id_cookie := http.Cookie{
			Name:  "id",
			Value: strconv.Itoa(id),
		}
		name_cookie := http.Cookie{
			Name:  "username",
			Value: u_name,
		}
		http.SetCookie(w, &id_cookie)
		http.SetCookie(w, &name_cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

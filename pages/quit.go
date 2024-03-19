package pages

import (
	"net/http"
)

func Quit(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "id",
		Value:  "",
		MaxAge: -1,
	}
	c1 := http.Cookie{
		Name:   "username",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)
	http.SetCookie(w, &c1)
	http.Redirect(w, r, "/", http.StatusFound)
}

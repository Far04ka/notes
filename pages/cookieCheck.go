package pages

import (
	"net/http"
)

func Check(w http.ResponseWriter, r *http.Request) (string, string) {
	id, err1 := r.Cookie("id")
	name, err2 := r.Cookie("username")
	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/auth", http.StatusFound)
		return "", ""
	}
	return id.Value, name.Value
}

package auth

import "net/http"

func SetCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

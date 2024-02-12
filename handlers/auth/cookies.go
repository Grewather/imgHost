package auth

import "net/http"

func SetCookie(w http.ResponseWriter, token string, expiresIn int) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   expiresIn,
	}
	http.SetCookie(w, &cookie)
}

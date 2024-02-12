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

func GetCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

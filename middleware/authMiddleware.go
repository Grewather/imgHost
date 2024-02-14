package middleware

import (
	"imgHost/handlers/auth"
	"imgHost/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static") || strings.HasPrefix(r.URL.Path, "/api") {
			next.ServeHTTP(w, r)
			return
		}
		cookie, err := r.Cookie("token")
		if err != nil {
			if r.URL.Path != "/" {
				http.Redirect(w, r, "/", http.StatusMovedPermanently)
				return
			}
		} else {
			userInfo, err := utils.GetUserInfo(cookie.Value)
			if err != nil || len(userInfo.ID) == 0 {
				auth.SetCookie(w, "", -1)
				http.Redirect(w, r, "/", http.StatusMovedPermanently)
				return
			}
			if !strings.HasPrefix(r.URL.Path, "/upload") && !strings.HasPrefix(r.URL.Path, "/images") {
				http.Redirect(w, r, "/upload", http.StatusMovedPermanently)
				return
			} else {
				next.ServeHTTP(w, r)
				return
			}

		}

		next.ServeHTTP(w, r)
	})
}

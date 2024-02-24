package middleware

import (
	"github.com/joho/godotenv"
	"imgHost/db"
	"imgHost/handlers/auth"
	"imgHost/utils"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	godotenv.Load(".env")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/admin") || strings.HasPrefix(r.URL.Path, "/api/admin") {
			authToken, err := r.Cookie("token")
			if err != nil {
				http.Error(w, "Access token not found", http.StatusUnauthorized)
				return
			}
			userinfo, err := utils.GetUserInfo(authToken.Value)
			if err != nil || userinfo.ID != os.Getenv("ADMIN_ID") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/static") || strings.HasPrefix(r.URL.Path, "/api") || strings.HasPrefix(r.URL.Path, "/i") {
			next.ServeHTTP(w, r)
			return
		}
		cookie, err := r.Cookie("token")
		if err != nil {
			if r.URL.Path != "/" {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		} else {
			userInfo, err := utils.GetUserInfo(cookie.Value)
			if err != nil || len(userInfo.ID) == 0 {
				auth.SetCookie(w, "", -1)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			res := db.GetIdFromDb(userInfo.ID)
			if !res {

				auth.SetCookie(w, "", -1)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			if !strings.HasPrefix(r.URL.Path, "/upload") && !strings.HasPrefix(r.URL.Path, "/gallery") {
				http.Redirect(w, r, "/upload", http.StatusFound)
				return
			} else {
				next.ServeHTTP(w, r)
				return
			}

		}

		next.ServeHTTP(w, r)
	})
}

package admin

import (
	"github.com/joho/godotenv"
	"html/template"
	"imgHost/utils"
	"net/http"
	"os"
)

func AdminPage(w http.ResponseWriter, r *http.Request) {
	godotenv.Load(".env")
	authToken, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Access token not found", http.StatusUnauthorized)
		return
	}
	userinfo, err := utils.GetUserInfo(authToken.Value)
	if err != nil {
		http.Error(w, "Error getting user info", http.StatusInternalServerError)
		return
	}
	if userinfo.ID != os.Getenv("ADMIN_ID") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	tmpl := template.Must(template.ParseFiles("./templates/admin.html"))
	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

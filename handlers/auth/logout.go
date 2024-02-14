package auth

import (
	"github.com/joho/godotenv"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	godotenv.Load(".env")
	cookie, err := r.Cookie("token")
	data := url.Values{}
	data.Set("token", cookie.Value)
	data.Set("token_type_hint", "access_token")
	req, err := http.NewRequest("POST", "https://discord.com/api/v10/oauth2/token/revoke", strings.NewReader(data.Encode()))
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(os.Getenv("DISCORD_CLIENT_ID"), os.Getenv("DISCORD_CLIENT_SECRET"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error sending request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	SetCookie(w, "", -1)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

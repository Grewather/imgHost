package auth

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"imgHost/db"
	"imgHost/models"
	"imgHost/utils"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func LoginCallback(w http.ResponseWriter, r *http.Request) {
	godotenv.Load(".env")

	code := r.URL.Query().Get("code")

	if code != "" {
		params := url.Values{}
		params.Add("grant_type", "authorization_code")
		params.Add("code", code)
		params.Add("redirect_uri", os.Getenv("DISCORD_REDIRECT_URI"))
		params.Add("scope", "identify")

		req, err := http.NewRequest("POST", "https://discord.com/api/v10/oauth2/token", strings.NewReader(params.Encode()))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error creating request"))
			return
		}

		req.SetBasicAuth(os.Getenv("DISCORD_CLIENT_ID"), os.Getenv("DISCORD_CLIENT_SECRET"))

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error sending request"))
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading response body"))
			return
		}

		var tokenResponse TokenResponse
		err = json.Unmarshal(body, &tokenResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error decoding JSON response"))
			return
		}

		userInfo, err := utils.GetUserInfo(tokenResponse.AccessToken)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error getting user info"))
			return
		}
		res := db.GetIdFromDb(userInfo.DiscordId)
		if !res {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("You are not invited"))
			return
		}
		account := models.Account{
			DiscordId: userInfo.DiscordId,
			Username:  userInfo.Username,
		}
		db.AddToDb(account)

		SetCookie(w, tokenResponse.AccessToken, tokenResponse.ExpiresIn)
		http.Redirect(w, r, "/upload", http.StatusMovedPermanently)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No code provided"))
	}
}

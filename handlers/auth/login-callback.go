package auth

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"imgHost/db"
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
type UserInfo struct {
	ID       string `bson:"id"`
	Username string `bson:"username"`
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

		userInfo, err := getUserInfo(tokenResponse.AccessToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error getting user info"))
			return
		}
		response := struct {
			TokenResponse
			UserInfo UserInfo `json:"userInfo"`
		}{
			TokenResponse: tokenResponse,
			UserInfo:      userInfo,
		}
		db.AddToDb(userInfo.ID, userInfo.Username)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No code provided"))
	}
}

func getUserInfo(accessToken string) (UserInfo, error) {
	req, err := http.NewRequest("GET", "https://discord.com/api/v10/users/@me", nil)
	if err != nil {
		return UserInfo{}, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return UserInfo{}, err
	}
	defer resp.Body.Close()

	var userInfo UserInfo
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return UserInfo{}, err
	}

	return userInfo, nil
}

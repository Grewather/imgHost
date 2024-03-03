package utils

import (
	"encoding/json"
	"imgHost/models"
	"net/http"
)

func GetUserInfo(accessToken string) (models.Account, error) {
	req, err := http.NewRequest("GET", "https://discord.com/api/v10/users/@me", nil)
	if err != nil {
		return models.Account{}, err

	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.Account{}, err
	}
	defer resp.Body.Close()
	var userInfo models.Account
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return models.Account{}, err
	}

	return userInfo, err
}

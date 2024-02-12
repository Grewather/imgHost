package utils

import (
	"encoding/json"
	"net/http"
)

type UserInfo struct {
	ID       string `bson:"id"`
	Username string `bson:"username"`
}

func GetUserInfo(accessToken string) (UserInfo, error) {
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

	return userInfo, err
}

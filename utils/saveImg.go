package utils

import (
	"errors"
	"fmt"
	"imgHost/models"
	"io"
	"os"
)

func SaveFileToUploads(file io.Reader, userInfo models.Account, extension string, randString string) error {
	_, err := os.Stat("uploads/" + userInfo.DiscordId)
	if os.IsNotExist(err) {
		err = os.MkdirAll("uploads/"+userInfo.DiscordId, 0755)
		if err != nil {
			fmt.Println(err.Error())
			return errors.New("error creating directory")
		}
	}
	ext := CheckFileExt(extension)
	if !ext {
		return errors.New("invalid file extension")
	}
	for {
		if checkIfYouCanAdd(randString, userInfo.DiscordId, extension) {
			break
		}
		randString = GetRandomString()
	}

	outFile, err := os.Create("uploads/" + userInfo.DiscordId + "/" + randString + extension)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("error creating file")
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("error creating file")
	}
	return nil

}
func checkIfYouCanAdd(randString, discordid, extension string) bool {
	pathfile := "uploads/" + discordid + "/" + randString + extension

	if _, err := os.Stat(pathfile); err == nil {
		return false
	} else if os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

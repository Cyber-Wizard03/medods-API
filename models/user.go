package models

import (
	"fmt"
	"guzkod/database"

	"guzkod/utils"
)

func VerifyPassword(password, hashedPassword string) bool {
	if utils.Hash(password, "nswpwjcppfjwlzppiwocjwoejofjsiowhf") == hashedPassword {
		return true
	} else {
		return false
	}
}

func LoginCheck(username string, password string, fingerprint string) (map[string]string, string, error) {

	var err error

	data, _ := database.GetUser("login", username)

	if VerifyPassword(password, data.Password) {

		createone := database.Generation()
		createtwo := database.Generation()

		database.CreateSession(data.ID, createone, createtwo, fingerprint)

		keyone := database.Static() + createone
		keytwo := database.Static() + createtwo

		at, _ := utils.GenerateAccessToken(data.ID, keyone)
		rt, _ := utils.GenerateRefreshToken(data.ID, keytwo)
		tokens := map[string]string{
			"accessToken":  at,
			"refreshToken": rt,
		}
		str := "Проверка прошла успешно"
		fmt.Println(str)
		return tokens, str, nil

	} else {
		str := "Неверный логин или пароль"
		fmt.Println(str)
		return map[string]string{}, str, err

	}

}
func GetUserByID(userID uint) (database.User, error) {

	data, _ := database.GetUser("id", userID)
	// if data == " " {
	// 	return data, errors.New("User not found!")
	// }

	data.Password = ""

	return data, nil

}

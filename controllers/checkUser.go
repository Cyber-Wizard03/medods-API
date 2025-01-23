package controllers

import (
	"fmt"
	"guzkod/database"
	"os"
	"os/exec"

	"guzkod/utils"
)

func CheckUser() {
	verificotion := database.GetCheckUser()
	if verificotion {
		fmt.Println("[Пользователи существуют]")
	} else {
		var (
			name     string
			password string
			start    string
		)
		fmt.Print("Введите имя нового пользователя: ")
		fmt.Scan(&name)
		fmt.Print("Создайте пароль для нового пользователя: ")
		fmt.Scan(&password)
		hashpass := utils.Hash(password, "nswpwjcppfjwlzppiwocjwoejofjsiowhf")
		fmt.Println("Нажмите Y чтобы продолжить...")
		fmt.Scan(&start)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		NewUser := database.User{}
		NewUser.Login = name
		NewUser.Password = hashpass
		database.CreateTask(NewUser, database.CollectionUser)
	}
}

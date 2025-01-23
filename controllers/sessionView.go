package controllers

import (
	"fmt"
	"guzkod/database"
	"time"
)

func CheckSession() {
	times := time.Now()
	T := times.Format("2006-01-02 15:04")
	fmt.Println("[Запуск службы проверки сессии] [" + T + "]")
	for {
		t := time.NewTimer(5 * time.Minute)
		<-t.C
		times := time.Now()
		T := times.Format("2006-01-02 15:04")
		currentTime := time.Now()
		TimeNow := currentTime.Unix()
		a := database.GetSession()
		Completion := a.Date

		if Completion < TimeNow {
			database.DropSession()
			fmt.Println("[Время сессии истекло] [" + T + "]")
		} else {
			fmt.Println("[Проверка прошла успешно] [" + T + "]")
		}
	}
}

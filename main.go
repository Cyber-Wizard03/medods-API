package main

import (
	"guzkod/controllers"
	"guzkod/database"
	"guzkod/server"
)

func main() {

	database.Connect()
	controllers.CheckUser()
	go controllers.CheckSession()
	server.Start()

}

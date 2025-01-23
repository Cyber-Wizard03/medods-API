package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"guzkod/database"
	"guzkod/models"
	"guzkod/utils"
)

type LoginInput struct {
	Login       string `json:"login" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Fingerprint string `bson:"fingerprint"`
}
type MesFeedback struct {
	// ID           primitive.ObjectID `bson:"_id"`
	FIO          string `bson:"FIO" json:"FIO"`
	Phone        string `bson:"Phone" json:"Phone"`
	Organization string `bson:"Organization" json:"Organization"`
	Department   string `bson:"Department" json:"Department"`
	Doctor       string `bson:"Doctor" json:"Doctor"`
	Type_appeal  string `bson:"Type_appeal" json:"Type_appeal"`
	Messages     string `bson:"Messages" json:"Messages"`
}

// func Feedback(c *gin.Context) {

// 	var input MesFeedback

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	fmt.Println(input.FIO)
// 	fmt.Println(input)
// 	var text = "Данные пришли"
// 	c.JSON(http.StatusOK, text)

// }
func Exit(c *gin.Context) {
	c.SetCookie("refresh", "", -1, "/", "https://www.guzkod.ru/", false, false)
	c.SetCookie("access", "", -1, "/", "https://www.guzkod.ru/", false, false)
	database.DropSession()
}

type Token struct {
	Access string `json:"actoken"`
}

func TokenVerification(c *gin.Context) {
	c.JSON(http.StatusOK, true)
}

func CurrentUser(c *gin.Context) {

	user_id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("До")
	u, err := models.GetUserByID(user_id)
	fmt.Println("После")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

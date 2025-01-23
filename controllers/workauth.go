package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Password struct {
	Pas string `json:"password" binding:"required"`
}

func WorkAuth(c *gin.Context) {
	var input Password

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dt := time.Now()
	dt.Format("01-02-2006")

}

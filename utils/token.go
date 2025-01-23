package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func GenerateAccessToken(user_id uint, secret string) (string, error) {
	lifetimeA := time.Second * 60 * 2
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["aut"] = true
	claims["uid"] = user_id
	claims["exp"] = time.Now().Add(lifetimeA).Unix()
	token, _ := accessToken.SignedString([]byte(secret))

	return token, nil
}

func GenerateRefreshToken(user_id uint, secret string) (string, error) {
	lifetimeR := time.Second * 60 * 60 * 3
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["uid"] = user_id
	claims["exp"] = time.Now().Add(lifetimeR).Unix()
	token, _ := refreshToken.SignedString([]byte(secret))

	return token, nil
}

func TokenValid(t string, mysecret string) bool {

	tokenString := t

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(mysecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["uid"], claims["exp"])
		return true
	} else {
		fmt.Println(err)
		return false
		// "signature is invalid"
	}
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("refreshToken")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

package middlewares

import (
	"fmt"
	"guzkod/database"
	"guzkod/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		cookie__access, _ := c.Cookie("access")
		cookie__refresh, _ := c.Cookie("refresh")

		if len(cookie__refresh) != 0 {

			if len(cookie__access) != 0 {

				db := database.GetSession()
				key := database.Static() + db.KeyOne

				term := utils.TokenValid(cookie__access, key)

				if term {

					fmt.Println("[Токен действителен]")

				} else {

					db := database.GetSession()
					key_2 := database.Static() + db.KeyTwo

					examination := utils.TokenValid(cookie__refresh, key_2)

					if examination {
						db := database.GetSession()
						id := db.ID_USER

						OldKey := db.KeyOne
						createone := database.Generation()

						database.UpdateSession(OldKey, createone)
						keyone := database.Static() + createone

						token, _ := utils.GenerateAccessToken(id, keyone)
						c.SetCookie("access", token, 900, "/", "https://www.guzkod.ru/", false, false)
						fmt.Println("[Токен обновлен]")

					} else {
						c.String(http.StatusUnauthorized, "Unauthorized")
						c.Abort()

					}
				}
			} else {
				db := database.GetSession()
				id := db.ID_USER
				OldKey := db.KeyOne
				createone := database.Generation()
				fmt.Println("1")
				database.UpdateSession(OldKey, createone)
				fmt.Println("2")
				keyone := database.Static() + createone
				fmt.Println("3")
				token, _ := utils.GenerateAccessToken(id, keyone)
				c.SetCookie("access", token, 900, "/", "https://www.guzkod.ru/", false, false)
			}

		} else {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()

		}
	}
}

package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get cookie
		if cookie, err := c.Cookie("label"); err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}

		// Cookie verification failed
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
		c.Abort()
	}
}

// func CreateCookie(){
// 	func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
// 	// Имя первого параметра — это имя файла cookie;
// 	// Второе значение параметра — это значение cookie;
// 	// Третий параметр maxAge — это срок действия файла cookie.Если файл cookie существует дольше установленного времени, он становится недействительным, он больше не является нашим действительным файлом cookie, и его единицей времени является секунда;
// 	// Путь четвертого параметра — это каталог, в котором находится файл cookie;
// 	// Пятый домен — это домен, что означает область действия нашего файла cookie, который может быть localhost или вашим доменным именем, в зависимости от вашей собственной ситуации;
// 	// Шестое безопасное указывает, можно ли получить к нему доступ только через https, если это правда, это может быть только https;
// 	// Седьмой httpOnly указывает, можно ли манипулировать файлом cookie с помощью кода js.
// }

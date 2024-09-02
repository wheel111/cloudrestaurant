package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const CookieName = "cookie_user"
const CookieTimeLength = 10 * 60

func CookieAuth(c *gin.Context) (*http.Cookie, error) {
	cookie, err := c.Request.Cookie(CookieName)
	if err == nil {
		c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		return cookie, nil
	} else {
		return nil, err
	}
}

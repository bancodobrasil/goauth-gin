package goauthgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticateGin(handler func(h http.Handler) http.Handler) gin.HandlerFunc {

	return func(c *gin.Context) {

		handler(nil).ServeHTTP(c.Writer, c.Request)

		if c.Writer.Status() == http.StatusUnauthorized || c.Writer.Status() == http.StatusForbidden {

			c.Abort()

		}

	}

}

package goauthgin

import (
	"net/http"

	"github.com/bancodobrasil/goauth"
	"github.com/gin-gonic/gin"
)

// Authenticate is a gin middleware that runs the authentication handlers
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authenticate(c)
	}
}

// AuthenticateH is a gin middleware that runs the authentication handlers with an optional http.Handler argument
func AuthenticateH(handler func(h http.Handler) http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := authenticate(c)

		handler(nil).ServeHTTP(c.Writer, request)

		if c.Writer.Status() == http.StatusUnauthorized || c.Writer.Status() == http.StatusForbidden {
			c.Abort()
		}
	}
}

func authenticate(c *gin.Context) *http.Request {
	var err error
	var statusCode int
	request := c.Request

	for _, handler := range goauth.GetHandlers() {
		request, statusCode, err = handler.Handle(c.Request)
		if err == nil {
			return request
		}
	}

	if err != nil {
		respondWithError(c, &goauth.AuthMiddlewareError{
			Code:    statusCode,
			Message: err.Error(),
		})
		return request
	}
	return request
}

func respondWithError(c *gin.Context, e *goauth.AuthMiddlewareError) {
	c.AbortWithStatusJSON(e.Code, gin.H{"error": e.Message})
}

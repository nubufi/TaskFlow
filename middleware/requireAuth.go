package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
//
// Parameters:
//
// - c: The gin context
func AuthMiddleware(c *gin.Context) {
	// Get the token
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/signin")
		return
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/signin")
		return
	}

	// Check if the token is valid
	if !token.Valid {
		c.Redirect(http.StatusTemporaryRedirect, "/signin")
		return
	}

	// Get the user ID
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check if the token is expired
		if !claims.VerifyExpiresAt(jwt.TimeFunc().Unix(), true) {
			c.Redirect(http.StatusTemporaryRedirect, "/signin")
			return
		}
		c.Set("userID", claims["sub"])
	}
	c.Next()
}

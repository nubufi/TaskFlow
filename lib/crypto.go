package lib

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

	"taskflow/models"

	"github.com/golang-jwt/jwt"
)

// GenerateRandomUserID generates a random user ID as a hexadecimal string
func GenerateRandomUserID() string {
	// Define the length of the ID in bytes (e.g., 16 bytes for 128-bit ID)
	length := 16
	// Create a byte slice to hold the random data
	randomBytes := make([]byte, length)
	// Generate random bytes
	rand.Read(randomBytes)
	// Encode the random bytes as a hexadecimal string
	userID := hex.EncodeToString(randomBytes)

	return userID
}

// CreateToken creates a JWT token for a user
//
// Parameters:
//
// - user: The user for whom to create the token
//
// Returns:
//
// - string: The token string
//
// - error: An error if the token could not be created
func CreateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

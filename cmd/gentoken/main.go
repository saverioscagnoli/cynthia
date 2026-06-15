// cmd/gentoken/main.go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	claims := jwt.MapClaims{
		"discord_id": "123456789",
		"username":   "testuser",
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	}
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	fmt.Println(token)
}

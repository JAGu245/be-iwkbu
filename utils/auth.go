package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte

func init() {
	// Load .env proactively before fetching environment variables
	_ = godotenv.Load("../.env", ".env")
	
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("WARNING: JWT_SECRET not found in .env, generating a random one for this session. Users will be logged out on restart.")
		bytes := make([]byte, 32)
		rand.Read(bytes)
		jwtSecret = bytes
	} else {
		jwtSecret = []byte(secret)
	}
}

func GetJWTKey() []byte {
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		return []byte(secret)
	}
	return jwtSecret
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(GetJWTKey())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return GetJWTKey(), nil
	})

	return token, err
}

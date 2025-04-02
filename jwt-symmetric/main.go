package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type CustomerClaims struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func getSecreteKey() []byte {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("JWT_SECRET_KEY environment variable not set")
	}

	return []byte(secretKey)
}

func GenerateToken(userID int, username, role string, expiryMinutes int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expiryMinutes) * time.Minute)

	claims := CustomerClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "apps",
			Subject:   fmt.Sprintf("TOKEN-%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(getSecreteKey())
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
func validateToken(tokenString string) (*CustomerClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomerClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return getSecreteKey(), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*CustomerClaims)
	if !ok {
		return nil, fmt.Errorf("couldn't parse claims")
	}

	return claims, nil
}

func main() {

	tokenString, err := GenerateToken(123, "johndoe", "admin", 60)
	if err != nil {
		log.Fatalf("Error generating token: %v", err)

	}

	fmt.Println("Generated JWT Token:")
	fmt.Println(tokenString)
	fmt.Println()

	claims, err := validateToken(tokenString)
	if err != nil {
		log.Fatalf("Error validating token: %v", err)
	}

	fmt.Println("Token is valid!")
	fmt.Printf("User ID: %d\n", claims.UserID)
	fmt.Printf("Username : %s\n", claims.Username)
	fmt.Printf("Role : %s\n", claims.Role)
	fmt.Printf("Expires at: %v\n", claims.ExpiresAt.Time)
}

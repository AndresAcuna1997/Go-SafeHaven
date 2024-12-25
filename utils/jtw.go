package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateJWT(orgId int64) (string, error) {
	err := godotenv.Load()

	if err != nil {
		return "", err
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"orgId": orgId,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (int64, error) {
	err := godotenv.Load()

	if err != nil {
		return 0, err
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("token differs from signing method")
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return 0, errors.New("error parsing the token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("token invalida")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("token differs from signing method")
	}

	orgId, ok := claims["orgId"].(float64)

	if !ok {
		return 0, errors.New("orgId is not a valid number")
	}

	intOrgID := int64(orgId)

	return intOrgID, nil
}

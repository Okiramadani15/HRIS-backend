package utils

import (
    "time"
    "os"
    "github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userID uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

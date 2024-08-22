package auth

import (
	"fmt"
	"net/http"

	"github.com/HaikalRFadhilahh/auth/helper"
	"github.com/golang-jwt/jwt/v5"
)

func Validate(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	secretKey := []byte(helper.GetEnv("JWT_SECRET", ""))
	if tokenString == "" {
		helper.ErrorResponse(w, http.StatusUnauthorized, "error", "Token Cannot be Null!", nil)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Memastikan algoritma yang digunakan adalah HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		fmt.Print(err.Error())
		helper.ErrorResponse(w, http.StatusUnauthorized, "error", "Token Invalid!", nil)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		helper.ErrorResponse(w, http.StatusOK, "success", "Data Users from JWT", claims)
		return
	} else {
		helper.ErrorResponse(w, http.StatusUnauthorized, "error", "Invalid Token!", nil)
		return
	}

}

package auth

import (
	"net/http"

	"github.com/HaikalRFadhilahh/auth/helper"
	"github.com/golang-jwt/jwt/v5"
)

func Validate(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		helper.ErrorResponse(w, http.StatusForbidden, "Error", "Invalid Token!", nil)
		return
	}

	
}

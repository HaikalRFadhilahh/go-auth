package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/HaikalRFadhilahh/auth/db"
	"github.com/HaikalRFadhilahh/auth/helper"
	"github.com/HaikalRFadhilahh/auth/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Create Connection
	db, err := db.InitDB(db.EnvironmentDB())
	if err != nil {
		helper.ErrorResponse(w, http.StatusInternalServerError, "error", "Database Connection Refused!", nil)
		return
	}
	defer db.Close()

	var result models.UsersModels

	// Decode Data Request
	err = json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		helper.ErrorResponse(w, http.StatusInternalServerError, "error", "Internal Server Error", nil)
		return
	}

	pass := result.Password

	// Get Data Users
	err = db.QueryRow("select * from users where username=?", result.Username).Scan(&result.Id, &result.Nama, &result.Username, &result.Password, &result.Created_at, &result.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			helper.ErrorResponse(w, http.StatusUnauthorized, "error", "Unauthorized", nil)
			return
		}

		helper.ErrorResponse(w, http.StatusInternalServerError, "error", "Internal Server Error", nil)
		return
	}

	// Check Hash Password
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(pass)); err != nil {
		helper.ErrorResponse(w, http.StatusUnauthorized, "error", "Username or Password Invalid!", nil)
		return
	}

	// Generate JSON Web Token
	expired, err := strconv.Atoi(helper.GetEnv("JWT_EXPIRED_MINUTE", ""))
	if err != nil {
		helper.ErrorResponse(w, http.StatusInternalServerError, "error", "Internal Server Error", nil)
		return
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       result.Id,
		"nama":     result.Nama,
		"username": result.Username,
		"exp":      time.Now().Add(time.Minute * time.Duration(expired)).Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenJWT, err := claims.SignedString([]byte(helper.GetEnv("JWT_SECRET", "")))
	if err != nil {
		helper.ErrorResponse(w, http.StatusInternalServerError, "error", "Cannot Generate Token!!", nil)
		return
	}

	helper.ErrorResponse(w, http.StatusOK, "success", "Users Authentificated Token", tokenJWT)
}

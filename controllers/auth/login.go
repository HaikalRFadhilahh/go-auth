package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/HaikalRFadhilahh/auth/db"
	"github.com/HaikalRFadhilahh/auth/helper"
	"github.com/HaikalRFadhilahh/auth/models"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db, err := db.InitDB(db.EnvironmentDB())
	if err != nil {
		helper.ErrorResponse(w, http.StatusInternalServerError, "error", "Database Connection Refused!", nil)
		return
	}
	defer db.Close()

	var result models.UsersModels

	err = json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		helper.ErrorResponse(w, 500, "error", "Internal Server Error", nil)
		return
	}

	pass := result.Password

	err = db.QueryRow("select * from users where username=?", result.Username).Scan(&result.Id, &result.Nama, &result.Username, &result.Password, &result.Created_at, &result.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			helper.ErrorResponse(w, http.StatusUnauthorized, "error", "Unauthorized", nil)
			return
		}

		helper.ErrorResponse(w, 500, "error", "Internal Server Error", nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(pass)); err != nil {
		helper.ErrorResponse(w, http.StatusUnauthorized, "error", "Unauthorized", nil)
		return
	}

	helper.ErrorResponse(w, http.StatusOK, "success", "Users Data Authenticated", result)
}

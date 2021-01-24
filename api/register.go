package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ck46/ebus/user"
	"github.com/ck46/ebus/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

func Signup(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	if req.Method == "POST" {
		var usr User
		_ = json.NewDecoder(req.Body).Decode(&usr)

		config, _ := utils.LoadConfig("config.json")
		db, err := utils.DBCon(config)
		utils.PanicOnError(err)
		defer db.Close()

		_, err = user.Signup(db, &user.Request{
			FirstName: usr.FirstName,
			LastName:  usr.LastName,
			Phone:     usr.Phone,
			Email:     usr.Email,
			Password:  usr.Password,
		})

		if err != nil {
			json.NewEncoder(w).Encode(Exception{Message: "User signup failed."})
			return
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":    usr.Email,
			"password": usr.Password,
		})
		tokenString, error := token.SignedString([]byte(secret))
		if error != nil {
			fmt.Println(error)
		}

		res := &Response{
			Code:   200,
			Result: JwtToken{Token: tokenString},
		}
		json.NewEncoder(w).Encode(res)
	}
}

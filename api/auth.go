package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

func Login(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	if req.Method == "POST" {
		var usr User
		_ = json.NewDecoder(req.Body).Decode(&usr)

		// Test login credentials using user.Login
		_, err := VerifyUser(usr)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(Exception{Message: "User login failed."})
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
			json.NewEncoder(w).Encode(Exception{Message: "User login failed."})
			return
		}

		res := Response{
			Code:   200,
			Result: JwtToken{Token: tokenString},
		}

		json.NewEncoder(w).Encode(res)
	}
}

func Logout(w http.ResponseWriter, req *http.Request) {}

func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user User
		mapstructure.Decode(claims, &user)
		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
	}
}

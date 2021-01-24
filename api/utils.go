package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/ck46/ebus/store"
	"github.com/ck46/ebus/user"
	"github.com/ck46/ebus/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

func ServeFileHandler(res http.ResponseWriter, req *http.Request) {
	fname := path.Base(req.URL.Path)
	filePath := "files/" + fname
	http.ServeFile(res, req, filePath)
}

func VerifyUser(usr User) (*user.User, error) {
	// Test login credentials using user.Login
	config, _ := utils.LoadConfig("config.json")
	db, err := utils.DBCon(config)
	utils.PanicOnError(err)
	defer db.Close()

	res, err := user.Login(db, &user.Request{
		Email:    usr.Email,
		Password: usr.Password,
	})
	if err != nil {
		return nil, err
	}
	return res.User, nil
}

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(secret), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(req, decode, token.Claims)
					next.ServeHTTP(w, req)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			} else {
				json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}

func TestEndpoint(w http.ResponseWriter, req *http.Request) {
	decoded := context.Get(req, decode)
	var usr User
	mapstructure.Decode(decoded.(jwt.MapClaims), &usr)

	res, err := VerifyUser(usr)
	if err != nil {
		json.NewEncoder(w).Encode(Exception{Message: "Wrong user login credentials."})
		return
	}
	json.NewEncoder(w).Encode(res)
}

func StoreItemToAPIItem(item *store.Item) *Item {
	var images []string
	for i := 0; i < len(item.Images); i++ {
		images = append(images, item.Images[i].S3link)
	}
	return &Item{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		StoreName:   item.Store.Name,
		PromoImage:  item.PromoImage,
		Price:       item.Price,
		Stock:       item.Stock,
		Images:      images,
	}
}

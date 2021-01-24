package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ck46/ebus/store"
	"github.com/ck46/ebus/utils"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		params, ok := r.URL.Query()["id"] // get department name

		if !ok || len(params[0]) < 1 {
			res := ErrorResponse{
				Code:  301,
				Error: "Get parameter 'id' is missing.",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		id, err := strconv.ParseUint(params[0], 10, 64)

		if err != nil {
			res := ErrorResponse{
				Code:  301,
				Error: err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		config, _ := utils.LoadConfig("config.json")
		db, err := utils.DBCon(config)
		utils.PanicOnError(err)
		defer db.Close()

		dbitem, err := store.FindItemByID(db, uint(id))

		if err != nil {
			res := ErrorResponse{
				Code:  302,
				Error: err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}
		item := StoreItemToAPIItem(dbitem)
		json.NewEncoder(w).Encode(&Response{Code: 200, Result: item})
	}
}

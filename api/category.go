package api

import (
	"encoding/json"
	"net/http"

	"github.com/ck46/ebus/store"
	"github.com/ck46/ebus/utils"
)

func GetDepartments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
	if r.Method == "GET" {
		config, _ := utils.LoadConfig("config.json")
		db, err := utils.DBCon(config)
		utils.PanicOnError(err)
		defer db.Close()

		departments := store.GetAllDepartments(db)

		res := Response{
			Code:   200,
			Result: departments,
		}
		json.NewEncoder(w).Encode(res)
	}
}
func GetDepartmentCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
	if r.Method == "GET" {
		params, ok := r.URL.Query()["name"] // get department name

		if !ok || len(params[0]) < 1 {
			res := ErrorResponse{
				Code:  301,
				Error: "Get parameter 'name' is missing.",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		name := params[0]

		config, _ := utils.LoadConfig("config.json")
		db, err := utils.DBCon(config)
		utils.PanicOnError(err)
		defer db.Close()

		department, err := store.FindDepartmentByName(db, name)
		if err != nil {
			res := ErrorResponse{
				Code:  302,
				Error: err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		res := Response{
			Code:   200,
			Result: department.Children,
		}
		json.NewEncoder(w).Encode(res)
	}
}

func GetCategoryItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
	if r.Method == "GET" {
		params, ok := r.URL.Query()["name"]

		if !ok || len(params[0]) < 1 {
			res := ErrorResponse{
				Code:  301,
				Error: "Get parameter 'name' is missing.",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		name := params[0]

		config, _ := utils.LoadConfig("config.json")
		db, err := utils.DBCon(config)
		utils.PanicOnError(err)
		defer db.Close()

		category, err := store.FindCategoryByName(db, name)
		if err != nil {
			res := ErrorResponse{
				Code:  302,
				Error: err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}
		dbitems := store.FindItemsByCategory(db, category)

		var items []*Item
		for i := 0; i < len(dbitems); i++ {
			item := StoreItemToAPIItem(dbitems[i])
			items = append(items, item)
		}

		res := Response{
			Code:   200,
			Result: items,
		}
		json.NewEncoder(w).Encode(res)
	}
}

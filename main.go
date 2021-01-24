package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ck46/ebus/utils"
)

func main() {
	r := router()
	config, err := utils.LoadConfig("config.json")
	utils.PanicOnError(err)

	fmt.Println("Making connection to database")
	db, err := utils.DBCon(config)
	utils.PanicOnError(err)
	fmt.Println(fmt.Sprintf("Connected to [%s] database", config.Deployment))
	err = dbMigrate(db)
	utils.PanicOnError(err)

	defer db.Close()

	fmt.Println(fmt.Sprintf("Serving on port [%s]", config.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

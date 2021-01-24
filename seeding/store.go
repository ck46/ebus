package seeding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ck46/ebus/store"
	"github.com/jinzhu/gorm"
)

func SeedStores(db *gorm.DB) error {
	fmt.Println("Seeding stores...")
	jsonFile, err := os.Open("data/stores.json")
	defer jsonFile.Close()

	if err != nil {
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var stores Stores
	json.Unmarshal(byteValue, &stores)

	for i := 0; i < len(stores.Stores); i++ {
		_, err = store.FindStoreByName(db, stores.Stores[i].Name)
		if err != nil {
			_, _ = store.MakeStore(db, &store.Request{
				Name:        stores.Stores[i].Name,
				Description: stores.Stores[i].Description,
				ImageUrl:    stores.Stores[i].ImageUrl,
			})
		}
	}
	return nil
}

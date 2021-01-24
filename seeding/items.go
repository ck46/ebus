package seeding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ck46/ebus/store"
	"github.com/jinzhu/gorm"
)

func SeedItems(db *gorm.DB) error {
	fmt.Println("Seeding product items...")
	jsonFile, err := os.Open("data/subs_plans.json")
	defer jsonFile.Close()

	if err != nil {
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var items Items
	json.Unmarshal(byteValue, &items)

	for i := 0; i < len(items.Items); i++ {
		_, err = store.FindItemByName(db, items.Items[i].Name)
		if err != nil {
			_, err1 := store.MakeItem(db, &store.Request{
				Name:        items.Items[i].Name,
				StoreName:   items.Items[i].Store,
				Promo:       items.Items[i].PromoImage,
				Description: items.Items[i].Description,
				Stock:       items.Items[i].Stock,
				Price:       items.Items[i].Price,
				Categories:  items.Items[i].Categories,
			})
			if err1 == nil {
				_, _ = store.MakeItemImages(db, &store.Request{
					Name:   items.Items[i].Name,
					Images: items.Items[i].Images,
				})
			}
		}
	}

	return nil
}

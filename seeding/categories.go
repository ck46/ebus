package seeding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ck46/ebus/store"

	"github.com/jinzhu/gorm"
)

func SeedCategories(db *gorm.DB) error {
	fmt.Println("Seeding product categories...")
	jsonFile, err := os.Open("data/categories.json")
	defer jsonFile.Close()

	if err != nil {
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var categories Categories
	json.Unmarshal(byteValue, &categories)

	for i := 0; i < len(categories.Categories); i++ {
		_, err = store.FindCategoryByName(db, categories.Categories[i].Name)
		if err != nil {
			_, _ = store.MakeCategory(db, &store.Request{
				Name:        categories.Categories[i].Name,
				Description: categories.Categories[i].Description,
			})
		}
	}
	return nil
}

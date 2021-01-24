package seeding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ck46/ebus/store"
	"github.com/jinzhu/gorm"
)

func SeedDepartments(db *gorm.DB) error {
	fmt.Println("Seeding product departments...")
	jsonFile, err := os.Open("data/subs_plans.json")
	defer jsonFile.Close()

	if err != nil {
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var departments Departments
	json.Unmarshal(byteValue, &departments)

	for i := 0; i < len(departments.Departments); i++ {
		_, err = store.FindDepartmentByName(db, departments.Departments[i].Name)
		if err != nil {
			_, _ = store.MakeDepartment(db, &store.Request{
				Name:       departments.Departments[i].Name,
				Categories: departments.Departments[i].Categories,
			})
		}
	}

	return nil
}

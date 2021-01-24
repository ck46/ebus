package seeding

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func SeedAll(db *gorm.DB) error {
	err := SeedCategories(db)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = SeedDepartments(db)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = SeedStores(db)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = SeedItems(db)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

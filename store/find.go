package store

import (
	"github.com/jinzhu/gorm"
)

func FindStoreByID(db *gorm.DB, id uint) (*Store, error) {
	var store Store
	res := db.Find(&store, id)
	if res.RecordNotFound() {
		return nil, &StoreNotExistError{}
	}
	return &store, nil
}

func FindItemByID(db *gorm.DB, id uint) (*Item, error) {
	var item Item
	res := db.Preload("Images").Preload("Categories").Preload("Store").Find(&item, id)
	if res.RecordNotFound() {
		return nil, &ItemNotExistError{}
	}
	return &item, nil
}

func GetStoreItems(db *gorm.DB, store *Store) []*Item {
	var items []*Item
	res := db.Model(&store).Preload("Categories").Preload("Images").Preload("Store").Related(&items)
	if res.RecordNotFound() {
		return nil
	}
	return items
}

func GetItems(db *gorm.DB) []*Item {
	var items []*Item
	res := db.Preload("Categories").Preload("Images").Preload("Store").Find(&items)
	if res.RecordNotFound() {
		return nil
	}
	return items
}

func GetStores(db *gorm.DB) []*Store {
	var stores []*Store
	res := db.Find(&stores)
	if res.RecordNotFound() {
		return nil
	}
	return stores
}

func FindItemsByCategory(db *gorm.DB, category *Category) []*Item {
	var items []*Item
	res := db.Model(&category).Preload("Categories").Preload("Images").Preload("Store").Related(&items)
	if res.RecordNotFound() {
		return nil
	}
	return items
}

func FindCategoryByName(db *gorm.DB, name string) (*Category, error) {
	var category Category
	res := db.Where("name = ?", name).Find(&category)
	if res.RecordNotFound() {
		return nil, &CategoryNotExistError{}
	}
	return &category, nil
}

func FindDepartmentByCategory(db *gorm.DB, category *Category) (*Department, error) {
	var department Department
	res := db.Where("category_id = ?", category.ID).Preload("Category").Preload("Children").Find(&department)
	if res.RecordNotFound() {
		return nil, &DepartmentNotExistError{}
	}
	return &department, nil
}

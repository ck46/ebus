package store

import "github.com/jinzhu/gorm"

func CreateItem(db *gorm.DB, item *Item) (uint, error) {
	err := db.Create(item).Error
	if err != nil {
		return 0, err
	}
	return item.ID, err
}

func CreateStore(db *gorm.DB, store *Store) (uint, error) {
	err := db.Create(store).Error
	if err != nil {
		return 0, err
	}
	return store.ID, err
}

func CreateCategory(db *gorm.DB, category *Category) (uint, error) {
	err := db.Create(category).Error
	if err != nil {
		return 0, err
	}
	return category.ID, err
}

func CreateDepartment(db *gorm.DB, department *Department) (uint, error) {
	err := db.Create(department).Error
	if err != nil {
		return 0, err
	}
	return department.ID, err
}

func CreateImage(db *gorm.DB, image *Image) (uint, error) {
	err := db.Create(image).Error
	if err != nil {
		return 0, err
	}
	return image.ID, err
}

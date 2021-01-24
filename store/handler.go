package store

import "github.com/jinzhu/gorm"

func MakeItem(db *gorm.DB, req *Request) (*Response, error) {
	catalog := &Item{
		Description:   req.Description,
		Store:         *req.Store,
		StoreQuantity: req.Quantity,
		Price:         req.Price,
	}
	id, err := CreateItem(db, catalog)
	if err != nil {
		return nil, err
	}
	return &Response{Id: id}, err
}

func MakeStore(db *gorm.DB, req *Request) (*Response, error) {
	store := &Store{
		Name:        req.Name,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
	}
	id, err := CreateStore(db, store)
	if err != nil {
		return nil, err
	}
	return &Response{Id: id}, err
}

func MakeImage(db *gorm.DB, req *Request) (*Response, error) {
	image := &Image{
		Item:   *req.Item,
		S3link: req.S3link,
	}
	id, err := CreateImage(db, image)
	if err != nil {
		return nil, err
	}
	return &Response{Id: id}, err
}

func MakeCategory(db *gorm.DB, req *Request) (*Response, error) {
	category := &Category{
		Name:        req.Name,
		Description: req.Description,
	}
	id, err := CreateCategory(db, category)
	if err != nil {
		return nil, err
	}
	return &Response{Id: id}, err
}

func MakeDepartment(db *gorm.DB, req *Request) (*Response, error) {
	department := &Department{
		Category: *req.Category,
	}

	id, err := CreateDepartment(db, department)
	if err != nil {
		return nil, err
	}
	return &Response{Id: id}, err
}

func PushCategoryToDepartment(db *gorm.DB, req *Request) (*Response, error) {
	department := req.Department
	department.Children = append(department.Children, req.Category)
	err := db.Save(&department).Error
	if err != nil {
		return nil, err
	}
	return &Response{Department: department}, nil
}

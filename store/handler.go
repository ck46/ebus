package store

import "github.com/jinzhu/gorm"

func MakeItem(db *gorm.DB, req *Request) (*Response, error) {
	var categories []*Category
	for i := 0; i < len(req.Categories); i++ {
		category, err := FindCategoryByName(db, req.Categories[i])
		if err == nil {
			categories = append(categories, category)
		}
	}
	store, err := FindStoreByName(db, req.StoreName)
	if err != nil {
		return nil, err
	}
	item := &Item{
		Name:        req.Name,
		Description: req.Description,
		Store:       *store,
		Stock:       req.Quantity,
		Price:       req.Price,
		Categories:  categories,
	}
	id, err := CreateItem(db, item)
	if err != nil {
		return nil, err
	}
	return &Response{Id: id}, err
}

func MakeItemImages(db *gorm.DB, req *Request) (*Response, error) {
	item, err := FindItemByName(db, req.Name)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(req.Images); i++ {
		_, _ = MakeImage(db, &Request{Item: item, S3link: req.Images[i]})
	}

	return &Response{Item: item}, nil
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
	var categories []*Category
	for i := 0; i < len(req.Categories); i++ {
		category, err := FindCategoryByName(db, req.Categories[i])
		if err == nil {
			categories = append(categories, category)
		}
	}

	department := &Department{
		Category: *req.Category,
		Children: categories,
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

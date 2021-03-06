package store

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	Name        string `gorm:"not null; unique_index"`
	StoreID     uint
	Store       Store
	PromoImage  string
	Images      []*Image
	Categories  []*Category `gorm:"many2many:item_categories"`
	Description string
	Stock       uint
	Price       float64
}

type Store struct {
	gorm.Model
	Approved    bool   `gorm:"default:false"`
	Name        string `gorm:"not null; unique_index"`
	ImageUrl    string
	Description string
}

type Category struct {
	gorm.Model
	Name        string `gorm:"not null; unique"`
	Description string
}

type Department struct {
	gorm.Model
	CategoryID uint
	Category   Category
	Children   []*Category
}

type Image struct {
	gorm.Model
	ItemID uint
	Item   Item
	S3link string `gorm:"not null; unique"`
}

type Request struct {
	Id          uint
	Name        string
	StoreName   string
	Store       *Store
	Promo       string
	Description string
	S3link      string
	ImageUrl    string
	Item        *Item
	Price       float64
	Quantity    uint
	Category    *Category
	Department  *Department
	Categories  []string
	Stock       uint
	Images      []string
}

type Response struct {
	Id         uint
	Item       *Item
	Store      *Store
	Image      *Image
	Department *Department
	Category   *Category
}

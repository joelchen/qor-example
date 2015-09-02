package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/media_library"
)

type Product struct {
	gorm.Model
	Name           string
	Code           string
	Price          float32
	Description    string `sql:"size:2000"`
	Images         []ProductImage
	ColorVariation []ColorVariation
	Category       Category
}

type ProductImage struct {
	gorm.Model
	Image media_library.FileSystem
}

type ColorVariation struct {
	gorm.Model
	ProductID      uint
	Color          Color
	SizeVariations []SizeVariation
}

type SizeVariation struct {
	gorm.Model
	ColorVariationID  uint
	Size              Size
	AvailableQuantity uint
}

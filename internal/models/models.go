package models

import "gorm.io/gorm"

type (
	Model struct {
		gorm.Model

		Brand       Brand
		Number      string
		OriginalBox bool
		Price       float32
		Scale       string
		State       string
		Type        Type
	}

	Brand struct {
		gorm.Model

		Name string
	}

	Type struct {
		gorm.Model

		Name string
	}
)

func AllModels() []interface{} {
	return []interface{}{&Model{}, &Brand{}, &Type{}}
}

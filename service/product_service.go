package product

import (
	"../db"
	"../entity"
)

type Service struct {
}

type Product entity.Product

type Parameter struct {
	title string
}

func (s Service) Search(title string) ([]Product, error) {
	db := db.GetDB()

	// create product from product model
	var product []Product

	// set parameter
	p := Parameter{title}

	// confirm db connecting
	if err := db.Take(&product).Error; err != nil {
		return nil, err
	}

	// product search query
	tx := db

	// Get all record
	tx = tx.Find(&product)

	if p.title != "" {
		tx = tx.Where("title = ?", p.title).Find(&product)
	}

	return product, nil
}

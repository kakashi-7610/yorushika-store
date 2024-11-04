package repository

import (
	"log"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Text  string
	Name  string
	Price int
	Img   string
}

// func (db *Database) CreateProducts(products []*Product) (int, error) {
// 	result := db.db.Create(products)
// 	if result.Error != nil {
// 		log.Fatalf("failed to create products. error: %v", result.Error)
// 		return 0, result.Error
// 	}

// 	log.Printf("created products. count: %d", result.RowsAffected)
// 	return int(result.RowsAffected), nil
// }

func (db *Database) ListProducts() ([]Product, error) {
	products := []Product{}
	result := db.db.Find(&products)
	if result.Error != nil {
		log.Fatalf("failed to list products. error: %v", result.Error)
		return nil, result.Error
	}

	log.Printf("list products. count: %d", result.RowsAffected)
	return products, nil
}

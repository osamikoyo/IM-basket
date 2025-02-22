package models

type Product struct {
	ID uint64
	Name string
	Price uint64
}

type Basket struct {
	UserID uint64
	FullPrice uint64
	Products []Product
}


package models

import "github.com/osamikoyo/IM-basket/pkg/proto/pb"

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

func ToPbProduct(p Product) *pb.Product {
	return &pb.Product{
		ID: p.ID,
		Price: p.Price,
		Name: p.Name,
	}
}

func ToPbProducts(p []Product) []*pb.Product {
	var res []*pb.Product
	for _, pr := range p{
		res = append(res, ToPbProduct(pr))
	}
	return res
}

func ToPBbasket(b *Basket) *pb.Basket {
	return &pb.Basket{
		Products: ToPbProducts(b.Products),
		UserID: b.UserID,
		FullPrice: b.UserID,
	}
}

func ToModelsProduct(b *pb.Product) Product{
	return Product{
		ID: b.ID,
		Price: b.Price,
		Name: b.Name,
	}
}
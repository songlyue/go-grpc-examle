package service

import "context"

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (p *ProductService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	stock := p.GetStockById(request.GetProdId())
	return &ProductResponse{
		ProdStock: stock,
	}, nil
}

func (p *ProductService) mustEmbedUnimplementedProdServiceServer() {
}

func (p *ProductService) GetStockById(id int32) int32 {
	return 100
}

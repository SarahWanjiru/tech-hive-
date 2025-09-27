package service

import (
	"context"
	"github.com/tech-hive/ecommerce/model"
)

type ProductService interface {
 	Create(ctx context.Context, model model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel
 	Update(ctx context.Context, productModel model.ProductCreateOrUpdateModel, id string) model.ProductCreateOrUpdateModel
 	Delete(ctx context.Context, id string)
 	FindById(ctx context.Context, id string) model.ProductModel
 	FindAll(ctx context.Context) []model.ProductModel
 	Search(ctx context.Context, searchModel model.ProductSearchModel) ([]model.ProductModel, int64)
 }

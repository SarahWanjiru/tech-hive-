package repository

import (
 	"context"
 	"github.com/tech-hive/ecommerce/entity"
 	"github.com/tech-hive/ecommerce/model"
 )

type ProductRepository interface {
 	Insert(ctx context.Context, product entity.Product) entity.Product
 	Update(ctx context.Context, product entity.Product) entity.Product
 	Delete(ctx context.Context, product entity.Product)
 	FindById(ctx context.Context, id string) (entity.Product, error)
 	FindByProductId(ctx context.Context, productId string) (entity.Product, error)
 	FindAl(ctx context.Context) []entity.Product
 	Search(ctx context.Context, searchModel model.ProductSearchModel) ([]entity.Product, int64)
 }

package impl

import (
	"context"
	"github.com/tech-hive/ecommerce/common"
	"github.com/tech-hive/ecommerce/configuration"
	"github.com/tech-hive/ecommerce/entity"
	"github.com/tech-hive/ecommerce/exception"
	"github.com/tech-hive/ecommerce/model"
	"github.com/tech-hive/ecommerce/repository"
	"github.com/tech-hive/ecommerce/service"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
)

func NewProductServiceImpl(productRepository *repository.ProductRepository, cache *redis.Client) service.ProductService {
	return &productServiceImpl{ProductRepository: *productRepository, Cache: cache}
}

type productServiceImpl struct {
	repository.ProductRepository
	Cache *redis.Client
}

func (service *productServiceImpl) Create(ctx context.Context, productModel model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel {
	common.Validate(productModel)
	product := entity.Product{
		Name:        productModel.Name,
		Description: productModel.Description,
		Price:       productModel.Price,
		Stock:       productModel.Stock,
		ImageUrl:    productModel.ImageUrl,
	}
	service.ProductRepository.Insert(ctx, product)
	return productModel
}

func (service *productServiceImpl) Update(ctx context.Context, productModel model.ProductCreateOrUpdateModel, id string) model.ProductCreateOrUpdateModel {
	common.Validate(productModel)
	product := entity.Product{
		ProductId:   uuid.MustParse(id),
		Name:        productModel.Name,
		Description: productModel.Description,
		Price:       productModel.Price,
		Stock:       productModel.Stock,
		ImageUrl:    productModel.ImageUrl,
	}
	service.ProductRepository.Update(ctx, product)
	return productModel
}

func (service *productServiceImpl) Delete(ctx context.Context, id string) {
	product, err := service.ProductRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	service.ProductRepository.Delete(ctx, product)
}

func (service *productServiceImpl) FindById(ctx context.Context, id string) model.ProductModel {
	productCache := configuration.SetCache[entity.Product](service.Cache, ctx, "product", id, service.ProductRepository.FindById)
	return model.ProductModel{
		Id:          productCache.ProductId.String(),
		Name:        productCache.Name,
		Description: productCache.Description,
		Price:       productCache.Price,
		Stock:       productCache.Stock,
		ImageUrl:    productCache.ImageUrl,
	}
}

func (service *productServiceImpl) FindAll(ctx context.Context) (responses []model.ProductModel) {
  	products := service.ProductRepository.FindAl(ctx)
  	for _, product := range products {
  		responses = append(responses, model.ProductModel{
  			Id:          product.ProductId.String(),
  			Name:        product.Name,
  			Description: product.Description,
  			Price:       product.Price,
  			Stock:       product.Stock,
  			ImageUrl:    product.ImageUrl,
  		})
  	}
  	if len(products) == 0 {
  		return []model.ProductModel{}
  	}
  	return responses
  }

func (service *productServiceImpl) Search(ctx context.Context, searchModel model.ProductSearchModel) ([]model.ProductModel, int64) {
  	products, totalCount := service.ProductRepository.Search(ctx, searchModel)

  	var responses []model.ProductModel
  	for _, product := range products {
  		responses = append(responses, model.ProductModel{
  			Id:          product.ProductId.String(),
  			Name:        product.Name,
  			Description: product.Description,
  			Price:       product.Price,
  			Stock:       product.Stock,
  			ImageUrl:    product.ImageUrl,
  		})
  	}

  	return responses, totalCount
  }

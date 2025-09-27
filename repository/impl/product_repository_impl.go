package impl

import (
 	"context"
 	"errors"
 	"github.com/tech-hive/ecommerce/entity"
 	"github.com/tech-hive/ecommerce/exception"
 	"github.com/tech-hive/ecommerce/model"
 	"github.com/tech-hive/ecommerce/repository"
 	"github.com/google/uuid"
 	"gorm.io/gorm"
 )

func NewProductRepositoryImpl(DB *gorm.DB) repository.ProductRepository {
	return &productRepositoryImpl{DB: DB}
}

type productRepositoryImpl struct {
	*gorm.DB
}

func (repository *productRepositoryImpl) Insert(ctx context.Context, product entity.Product) entity.Product {
 	product.ProductId = uuid.New()
 	err := repository.DB.WithContext(ctx).Create(&product).Error
 	exception.PanicLogging(err)
 	return product
 }

func (repository *productRepositoryImpl) Update(ctx context.Context, product entity.Product) entity.Product {
	err := repository.DB.WithContext(ctx).Where("product_id = ?", product.Id).Updates(&product).Error
	exception.PanicLogging(err)
	return product
}

func (repository *productRepositoryImpl) Delete(ctx context.Context, product entity.Product) {
	err := repository.DB.WithContext(ctx).Delete(&product).Error
	exception.PanicLogging(err)
}

func (repository *productRepositoryImpl) FindById(ctx context.Context, id string) (entity.Product, error) {
 	var product entity.Product
 	result := repository.DB.WithContext(ctx).Unscoped().Where("product_id = ?", id).First(&product)
 	if result.RowsAffected == 0 {
 		return entity.Product{}, errors.New("product Not Found")
 	}
 	return product, nil
 }

func (repository *productRepositoryImpl) FindByProductId(ctx context.Context, productId string) (entity.Product, error) {
 	var product entity.Product
 	result := repository.DB.WithContext(ctx).Where("product_id = ?", productId).First(&product)
 	if result.Error != nil {
 		if result.Error == gorm.ErrRecordNotFound {
 			return entity.Product{}, gorm.ErrRecordNotFound
 		}
 		return entity.Product{}, result.Error
 	}
 	return product, nil
 }

func (repository *productRepositoryImpl) FindAl(ctx context.Context) []entity.Product {
 	var products []entity.Product
 	repository.DB.WithContext(ctx).Find(&products)
 	return products
 }

func (repository *productRepositoryImpl) Search(ctx context.Context, searchModel model.ProductSearchModel) ([]entity.Product, int64) {
 	query := repository.DB.WithContext(ctx).Model(&entity.Product{})

 	// Add search filters
 	if searchModel.Name != "" {
 		query = query.Where("name LIKE ?", "%"+searchModel.Name+"%")
 	}

 	if searchModel.MinPrice > 0 {
 		query = query.Where("price >= ?", searchModel.MinPrice)
 	}

 	if searchModel.MaxPrice > 0 {
 		query = query.Where("price <= ?", searchModel.MaxPrice)
 	}

 	if searchModel.InStock != nil && *searchModel.InStock {
 		query = query.Where("stock > ?", 0)
 	}

 	// Get total count
 	var totalCount int64
 	query.Count(&totalCount)

 	// Set default pagination values
 	page := searchModel.Page
 	if page <= 0 {
 		page = 1
 	}

 	limit := searchModel.Limit
 	if limit <= 0 {
 		limit = 10
 	}

 	offset := (page - 1) * limit
 	query = query.Offset(offset).Limit(limit)

 	// Add sorting
 	sortBy := searchModel.SortBy
 	if sortBy == "" {
 		sortBy = "created_at"
 	}

 	sortOrder := searchModel.SortOrder
 	if sortOrder == "" {
 		sortOrder = "desc"
 	}

 	query = query.Order(sortBy + " " + sortOrder)

 	var products []entity.Product
 	query.Find(&products)

 	return products, totalCount
 }

package impl

import (
	"context"
	"github.com/tech-hive/ecommerce/exception"
	"github.com/tech-hive/ecommerce/model"
	"github.com/tech-hive/ecommerce/repository"
	"github.com/tech-hive/ecommerce/service"
)

func NewTransactionDetailServiceImpl(transactionDetailRepository *repository.TransactionDetailRepository) service.TransactionDetailService {
	return &transactionDetailServiceImpl{TransactionDetailRepository: *transactionDetailRepository}
}

type transactionDetailServiceImpl struct {
	repository.TransactionDetailRepository
}

func (transactionDetailService *transactionDetailServiceImpl) FindById(ctx context.Context, id string) model.TransactionDetailModel {
	transactionDetail, err := transactionDetailService.TransactionDetailRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return model.TransactionDetailModel{
		Id:            transactionDetail.Id.String(),
		SubTotalPrice: transactionDetail.SubTotalPrice,
		Price:         transactionDetail.Price,
		Quantity:      transactionDetail.Quantity,
		Product: model.ProductModel{
			Id:          transactionDetail.Product.Id.String(),
			Name:        transactionDetail.Product.Name,
			Description: transactionDetail.Product.Description,
			Price:       transactionDetail.Product.Price,
			Stock:       transactionDetail.Product.Stock,
			ImageUrl:    transactionDetail.Product.ImageUrl,
		},
	}
}

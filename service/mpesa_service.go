package service

import (
	"context"
	"github.com/tech-hive/ecommerce/model"
)

type MpesaService interface {
	InitiateSTKPush(ctx context.Context, request model.MpesaPaymentRequest) (model.MpesaPaymentResponse, error)
	ProcessCallback(ctx context.Context, callback model.MpesaCallbackRequest) error
	GeneratePassword() string
	GenerateTimestamp() string
}
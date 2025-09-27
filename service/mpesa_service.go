package service

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
)

type MpesaService interface {
	InitiateSTKPush(ctx context.Context, request model.MpesaPaymentRequest) (model.MpesaPaymentResponse, error)
	ProcessCallback(ctx context.Context, callback model.MpesaCallbackRequest) error
	GeneratePassword() string
	GenerateTimestamp() string
}
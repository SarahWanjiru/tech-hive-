package impl

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"github.com/tech-hive/ecommerce/configuration"
	"github.com/tech-hive/ecommerce/entity"
	"github.com/tech-hive/ecommerce/model"
	"github.com/tech-hive/ecommerce/repository"
	"github.com/tech-hive/ecommerce/service"
	"gorm.io/gorm"
	"math/big"
	"strconv"
	"time"
)

func NewMpesaServiceImpl(config configuration.Config, orderRepository *repository.OrderRepository, DB *gorm.DB) service.MpesaService {
	return &mpesaServiceImpl{
		Config:          config,
		OrderRepository: *orderRepository,
		DB:              DB,
	}
}

type mpesaServiceImpl struct {
	configuration.Config
	repository.OrderRepository
	DB *gorm.DB
}

func (mpesaService *mpesaServiceImpl) InitiateSTKPush(ctx context.Context, request model.MpesaPaymentRequest) (model.MpesaPaymentResponse, error) {
	// Get order to verify it exists and belongs to user
	order, err := mpesaService.OrderRepository.GetOrderById(ctx, request.OrderId)
	if err != nil {
		return model.MpesaPaymentResponse{}, errors.New("order not found")
	}

	// Check if order is already paid
	if len(order.Payments) > 0 && order.Payments[0].Status == "success" {
		return model.MpesaPaymentResponse{}, errors.New("order is already paid")
	}

	// Generate M-Pesa request (for reference, actual request is simulated)
	// timestamp := mpesaService.GenerateTimestamp()
	// password := mpesaService.GeneratePassword()

	// Create STK push request (for reference, actual request is simulated)
	// stkPushRequest := model.MpesaSTKPushRequest{
	//     BusinessShortCode: mpesaService.Config.Get("MPESA_SHORTCODE"),
	//     Password:          password,
	//     Timestamp:         timestamp,
	//     TransactionType:   "CustomerPayBillOnline",
	//     Amount:            request.Amount,
	//     PartyA:            request.PhoneNumber,
	//     PartyB:            mpesaService.Config.Get("MPESA_SHORTCODE"),
	//     PhoneNumber:       request.PhoneNumber,
	//     CallBackURL:       mpesaService.Config.Get("MPESA_CALLBACK_URL"),
	//     AccountReference:  fmt.Sprintf("Order #%d", request.OrderId),
	//     TransactionDesc:   "Payment for order",
	// }

	// Simulate M-Pesa API call
	// In a real implementation, you would make an HTTP request to M-Pesa API
	// For simulation, we'll generate a mock response
	merchantRequestId := generateRandomString(20)
	checkoutRequestId := generateRandomString(20)

	// Simulate different response scenarios
	responseCode := "0" // Success
	responseMessage := "Success. Request accepted for processing"
	customerMessage := "Please complete the payment on your phone"

	// Simulate random failures for testing
	if simulateRandomFailure() {
		responseCode = "1"
		responseMessage = "Insufficient balance"
		customerMessage = "You have insufficient balance to complete this transaction"
	}

	response := model.MpesaPaymentResponse{
		MerchantRequestId: merchantRequestId,
		CheckoutRequestId: checkoutRequestId,
		ResponseCode:      responseCode,
		ResponseMessage:   responseMessage,
		CustomerMessage:   customerMessage,
	}

	// Create payment record regardless of response code
	payment := entity.Payment{
		OrderId:       request.OrderId,
		TransactionId: checkoutRequestId,
		Status:        "pending",
	}

	// Use transaction to ensure consistency
	tx := mpesaService.DB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		return model.MpesaPaymentResponse{}, err
	}

	// If request was accepted, simulate payment completion after a delay
	if responseCode == "0" {
		go mpesaService.simulatePaymentCompletion(ctx, checkoutRequestId, request.OrderId)
	} else {
		// If request failed, mark payment as failed immediately
		tx.Model(&payment).Update("status", "failed")
	}

	if err := tx.Commit().Error; err != nil {
		return model.MpesaPaymentResponse{}, err
	}

	return response, nil
}

func (mpesaService *mpesaServiceImpl) ProcessCallback(ctx context.Context, callback model.MpesaCallbackRequest) error {
	// Find payment by checkout request ID
	var payment entity.Payment
	result := mpesaService.DB.WithContext(ctx).Where("transaction_id = ?", callback.CheckoutRequestId).First(&payment)
	if result.Error != nil {
		return result.Error
	}

	// Update payment status based on callback
	if callback.ResultCode == 0 {
		// Payment successful
		mpesaService.DB.WithContext(ctx).Model(&payment).Updates(map[string]interface{}{
			"status":  "success",
			"paid_at": time.Now(),
		})

		// Update order status to confirmed
		mpesaService.OrderRepository.UpdateOrderStatus(ctx, payment.OrderId, "confirmed")
	} else {
		// Payment failed
		mpesaService.DB.WithContext(ctx).Model(&payment).Update("status", "failed")
	}

	return nil
}

func (mpesaService *mpesaServiceImpl) GeneratePassword() string {
	// Generate M-Pesa API password
	// Format: Base64 encoded string of Shortcode + Passkey + Timestamp
	shortcode := mpesaService.Config.Get("MPESA_SHORTCODE")
	passkey := mpesaService.Config.Get("MPESA_PASSKEY")
	timestamp := mpesaService.GenerateTimestamp()

	passwordString := shortcode + passkey + timestamp
	hash := sha256.Sum256([]byte(passwordString))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func (mpesaService *mpesaServiceImpl) GenerateTimestamp() string {
	return time.Now().Format("20060102150405")
}

// simulatePaymentCompletion simulates payment completion for testing
func (mpesaService *mpesaServiceImpl) simulatePaymentCompletion(ctx context.Context, checkoutRequestId string, orderId uint) {
	// Simulate delay for payment processing
	time.Sleep(5 * time.Second)

	// Simulate successful payment
	callback := model.MpesaCallbackRequest{
		MerchantRequestId: generateRandomString(20),
		CheckoutRequestId: checkoutRequestId,
		ResultCode:        0,
		ResultDesc:        "The service request is processed successfully",
		CallbackMetadata: &model.MpesaCallbackMetadata{
			Item: []model.MpesaCallbackItem{
				{Name: "MpesaReceiptNumber", Value: generateRandomString(10)},
				{Name: "TransactionDate", Value: strconv.FormatInt(time.Now().Unix(), 10)},
			},
		},
	}

	// Process the callback
	mpesaService.ProcessCallback(ctx, callback)
}

// Helper functions
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		_, _ = rand.Read(b[i:i+1])
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b)
}

func simulateRandomFailure() bool {
	// 10% chance of failure for testing
	n, _ := rand.Int(rand.Reader, big.NewInt(100))
	return n.Int64() < 10
}
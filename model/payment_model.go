package model

type MpesaPaymentRequest struct {
	OrderId     uint    `json:"order_id" validate:"required"`
	PhoneNumber string  `json:"phone_number" validate:"required,len=12" example:"254712345678"`
	Amount      float64 `json:"amount" validate:"required,min=1" example:"1000.00"`
}

type MpesaPaymentResponse struct {
	MerchantRequestId string `json:"merchant_request_id"`
	CheckoutRequestId string `json:"checkout_request_id"`
	ResponseCode      string `json:"response_code"`
	ResponseMessage   string `json:"response_message"`
	CustomerMessage   string `json:"customer_message"`
}

type MpesaCallbackRequest struct {
	MerchantRequestId   string `json:"merchant_request_id"`
	CheckoutRequestId   string `json:"checkout_request_id"`
	ResultCode          int    `json:"result_code"`
	ResultDesc          string `json:"result_desc"`
	CallbackMetadata    *MpesaCallbackMetadata `json:"callback_metadata,omitempty"`
}

type MpesaCallbackMetadata struct {
	Item []MpesaCallbackItem `json:"item"`
}

type MpesaCallbackItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MpesaSTKPushRequest struct {
	BusinessShortCode string  `json:"business_short_code"`
	Password          string  `json:"password"`
	Timestamp         string  `json:"timestamp"`
	TransactionType   string  `json:"transaction_type"`
	Amount            float64 `json:"amount"`
	PartyA            string  `json:"party_a"`
	PartyB            string  `json:"party_b"`
	PhoneNumber       string  `json:"phone_number"`
	CallBackURL       string  `json:"call_back_url"`
	AccountReference  string  `json:"account_reference"`
	TransactionDesc   string  `json:"transaction_desc"`
}
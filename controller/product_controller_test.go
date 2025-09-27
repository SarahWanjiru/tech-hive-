package controller

import (
	"bytes"
	"encoding/json"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestProductController_Create_Success(t *testing.T) {
	// Setup
	config := configuration.New()

	// Create a mock product service
	mockProductService := new(MockProductService)

	// Create controller with mock service
	controller := NewProductController(&mockProductService, config)

	// Create Fiber app
	app := fiber.New()

	// Setup route
	controller.Route(app)

	// Test data
	productRequest := model.ProductCreateOrUpdateModel{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       10,
		ImageUrl:    "https://example.com/test.jpg",
	}

	expectedResponse := productRequest

	// Mock the service call
	mockProductService.On("Create", productRequest).Return(expectedResponse)

	// Create request
	requestBody, _ := json.Marshal(productRequest)
	req := httptest.NewRequest("POST", "/v1/api/product", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	resp, err := app.Test(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	// Clean up
	mockProductService.AssertExpectations(t)
}

func TestProductController_Create_ValidationError(t *testing.T) {
	// Setup
	config := configuration.New()
	mockProductService := new(MockProductService)
	controller := NewProductController(&mockProductService, config)
	app := fiber.New()
	controller.Route(app)

	// Test with invalid data (empty name)
	productRequest := model.ProductCreateOrUpdateModel{
		Name:        "", // Invalid: empty name
		Description: "Test Description",
		Price:       99.99,
		Stock:       10,
	}

	requestBody, _ := json.Marshal(productRequest)
	req := httptest.NewRequest("POST", "/v1/api/product", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	resp, err := app.Test(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode) // Should return validation error

	// Clean up
	mockProductService.AssertExpectations(t)
}

// MockProductService is a mock implementation of the ProductService interface
type MockProductService struct {
	// Add any necessary fields for mocking
}

func (m *MockProductService) Create(productModel model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel {
	// Mock implementation
	return productModel
}

func (m *MockProductService) Update(productModel model.ProductCreateOrUpdateModel, id string) model.ProductCreateOrUpdateModel {
	// Mock implementation
	return productModel
}

func (m *MockProductService) Delete(id string) {
	// Mock implementation
}

func (m *MockProductService) FindById(id string) model.ProductModel {
	// Mock implementation
	return model.ProductModel{}
}

func (m *MockProductService) FindAll() []model.ProductModel {
	// Mock implementation
	return []model.ProductModel{}
}

func (m *MockProductService) Search(searchModel model.ProductSearchModel) ([]model.ProductModel, int64) {
	// Mock implementation
	return []model.ProductModel{}, 0
}

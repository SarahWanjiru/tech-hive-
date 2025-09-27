package impl

import (
	"context"
	"github.com/tech-hive/ecommerce/entity"
	"github.com/tech-hive/ecommerce/repository"
	"github.com/tech-hive/ecommerce/service"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewSeedServiceImpl(userRepository *repository.UserRepository, productRepository *repository.ProductRepository, DB *gorm.DB) service.SeedService {
	return &seedServiceImpl{
		UserRepository:    *userRepository,
		ProductRepository: *productRepository,
		DB:                DB,
	}
}

type seedServiceImpl struct {
	repository.UserRepository
	repository.ProductRepository
	DB *gorm.DB
}

func (seedService *seedServiceImpl) SeedUsers(ctx context.Context) error {
	// Check if users already exist
	var count int64
	seedService.DB.WithContext(ctx).Model(&entity.User{}).Count(&count)
	if count > 0 {
		return nil // Users already seeded
	}

	// Hash passwords
	adminPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	customerPassword, _ := bcrypt.GenerateFromPassword([]byte("customer123"), bcrypt.DefaultCost)

	users := []entity.User{
		{
			Name:     "Administrator",
			Email:    "admin@example.com",
			Password: string(adminPassword),
			Role:     "admin",
			IsActive: true,
		},
		{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: string(customerPassword),
			Role:     "customer",
			IsActive: true,
		},
		{
			Name:     "Jane Smith",
			Email:    "jane@example.com",
			Password: string(customerPassword),
			Role:     "customer",
			IsActive: true,
		},
	}

	for _, user := range users {
		_, err := seedService.UserRepository.Register(ctx, user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (seedService *seedServiceImpl) SeedProducts(ctx context.Context) error {
	// Check if products already exist
	var count int64
	seedService.DB.WithContext(ctx).Model(&entity.Product{}).Count(&count)
	if count > 0 {
		return nil // Products already seeded
	}

	products := []entity.Product{
		{
			ProductId:   uuid.New(),
			Name:        "iPhone 15 Pro",
			Description: "Latest iPhone with advanced camera system and titanium design",
			Price:       999.99,
			Stock:       50,
			ImageUrl:    "https://example.com/iphone15pro.jpg",
		},
		{
			ProductId:   uuid.New(),
			Name:        "Samsung Galaxy S24",
			Description: "Premium Android smartphone with AI features",
			Price:       899.99,
			Stock:       30,
			ImageUrl:    "https://example.com/galaxy-s24.jpg",
		},
		{
			ProductId:   uuid.New(),
			Name:        "MacBook Pro 16-inch",
			Description: "Professional laptop with M3 chip and stunning display",
			Price:       2499.99,
			Stock:       20,
			ImageUrl:    "https://example.com/macbook-pro.jpg",
		},
		{
			ProductId:   uuid.New(),
			Name:        "Dell XPS 13",
			Description: "Ultra-portable laptop with InfinityEdge display",
			Price:       1299.99,
			Stock:       25,
			ImageUrl:    "https://example.com/dell-xps13.jpg",
		},
		{
			ProductId:   uuid.New(),
			Name:        "Sony WH-1000XM5",
			Description: "Industry-leading noise canceling wireless headphones",
			Price:       399.99,
			Stock:       100,
			ImageUrl:    "https://example.com/sony-wh1000xm5.jpg",
		},
		{
			ProductId:   uuid.New(),
			Name:        "iPad Air",
			Description: "Versatile tablet with M1 chip and all-screen design",
			Price:       599.99,
			Stock:       40,
			ImageUrl:    "https://example.com/ipad-air.jpg",
		},
		{
			ProductId:   uuid.New(),
			Name:        "Nintendo Switch OLED",
			Description: "Gaming console with vibrant OLED screen",
			Price:       349.99,
			Stock:       60,
			ImageUrl:    "https://example.com/switch-oled.jpg",
		},
		{
			ProductId:   uuid.New(),
			Name:        "Apple Watch Series 9",
			Description: "Advanced smartwatch with health monitoring features",
			Price:       399.99,
			Stock:       80,
			ImageUrl:    "https://example.com/apple-watch9.jpg",
		},
	}

	for _, product := range products {
		seedService.ProductRepository.Insert(ctx, product)
	}

	return nil
}

func (seedService *seedServiceImpl) SeedAll(ctx context.Context) error {
	// Seed users first
	if err := seedService.SeedUsers(ctx); err != nil {
		return err
	}

	// Then seed products
	if err := seedService.SeedProducts(ctx); err != nil {
		return err
	}

	return nil
}
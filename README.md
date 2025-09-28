 Tech Hive E-commerce Application

A comprehensive mini-ecommerce application built with Go Fiber backend and Vue.js frontend using clean architecture principles, featuring user authentication, product management, shopping cart, order processing, and M-Pesa payment integration with Kenyan Shillings (KSh) currency support.

### Core Features
- **User Authentication & Authorization**: JWT-based authentication with role-based access control
- **Product Management**: Full CRUD operations with search and filtering capabilities
- **Shopping Cart**: Add, update, remove items with quantity management
- **Order Management**: Complete order lifecycle from cart to delivery
- **Payment Integration**: M-Pesa STK Push simulation for payments
- **Database Migrations**: Comprehensive migration system for schema management
- **Database Seeding**: Sample data generation for testing and development

### Technical Features
- **Clean Architecture**: Separation of concerns with entities, repositories, services, and controllers
- **Database Support**: MySQL with GORM ORM
- **Caching**: Redis integration for performance optimization
- **API Documentation**: Swagger/OpenAPI documentation
- **Docker Support**: Complete containerization with Docker Compose
- **Testing**: Unit and integration test structure
- **Error Handling**: Comprehensive error handling and logging

## 🛠 Technology Stack

- **Backend**: Go 1.21, Go Fiber
- **Frontend**: Vue.js 3, TypeScript, Pinia
- **Database**: MySQL 8.0
- **Cache**: Redis
- **ORM**: GORM
- **Authentication**: JWT
- **Documentation**: Swagger
- **Containerization**: Docker & Docker Compose
- **Architecture**: Clean Architecture

## 📋 Prerequisites

- Docker and Docker Compose
- Go 1.21+ (for local development)
- Node.js 18+ and npm (for frontend development)
- Git

## 🚀 Quick Start

### Using Docker Compose (Recommended)

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd tech-hive
   ```

2. **Start all services**
   ```bash
   docker compose up -d
   ```

3. **Run database migrations**
   ```bash
   docker compose exec app migrate -database "mysql://gofiber:gofiber123@tcp(db:3306)/tech_hive" -path db/migrations up
   ```

4. **Seed sample data (optional)**
   ```bash
   # Using curl
   curl -X POST http://localhost:9999/v1/api/seed/all

   # Or access via browser
   # http://localhost:9999/swagger/
   ```

5. **Access the application**
   - Frontend: http://localhost:3000
   - API: http://localhost:9999
   - API Documentation: http://localhost:9999/swagger/
   - Database Admin: http://localhost:8080 (Adminer)
   - Redis Admin: http://localhost:8081 (Redis Commander)

### Local Development Setup

1. **Start databases**
   ```bash
   docker compose up -d db redis
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Run migrations**
   ```bash
   migrate -database "mysql://root:root@tcp(localhost:3306)/tech_hive" -path db/migrations up
   ```

5. **Run the application**
    ```bash
    go run main.go
    ```

 ### Frontend Development Setup

 1. **Install frontend dependencies**
    ```bash
    cd client
    npm install
    ```

 2. **Start development server**
    ```bash
    cd client
    npm run dev
    ```

 3. **Build for production**
    ```bash
    cd client
    npm run build
    ```

 4. **Preview production build**
    ```bash
    cd client
    npm run preview
    ```

## 📚 API Documentation

### Authentication Endpoints

#### Register User
```http
POST /v1/api/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "role": "customer"
}
```

#### Login
```http
POST /v1/api/authentication
Content-Type: application/json

{
  "username": "john@example.com",
  "password": "password123"
}
```

### Product Endpoints

#### Get All Products
```http
GET /v1/api/product
Authorization: Bearer <token>
```

#### Search Products
```http
POST /v1/api/product/search
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "iPhone",
  "min_price": 500,
  "max_price": 1500,
  "in_stock": true,
  "page": 1,
  "limit": 10,
  "sort_by": "price",
  "sort_order": "asc"
}
```

#### Create Product (Admin Only)
```http
POST /v1/api/product
Authorization: Bearer <admin-token>
Content-Type: application/json

{
  "name": "New Product",
  "description": "Product description",
  "price": 99.99,
  "stock": 100,
  "image_url": "https://example.com/image.jpg"
}
```

### Cart Endpoints

#### Get Cart
```http
GET /v1/api/cart
Authorization: Bearer <token>
```

#### Add to Cart
```http
POST /v1/api/cart/items
Authorization: Bearer <token>
Content-Type: application/json

{
  "product_id": "product-uuid",
  "quantity": 2
}
```

#### Update Cart Item
```http
PUT /v1/api/cart/items/1
Authorization: Bearer <token>
Content-Type: application/json

{
  "quantity": 3
}
```

### Order Endpoints

#### Create Order
```http
POST /v1/api/orders
Authorization: Bearer <token>
Content-Type: application/json

{
  "cart_id": 1,
  "shipping_address": "123 Main St, City, Country"
}
```

#### Get User Orders
```http
GET /v1/api/orders
Authorization: Bearer <token>
```

#### Cancel Order
```http
DELETE /v1/api/orders/1
Authorization: Bearer <token>
```

### Payment Endpoints (M-Pesa)

#### Initiate Payment
```http
POST /v1/api/mpesa/stkpush
Authorization: Bearer <token>
Content-Type: application/json

{
  "order_id": 1,
  "phone_number": "254712345678",
  "amount": 1000.00
}
```

## 🗄️ Database Schema

### Tables Created
- `tb_user`: User accounts and authentication
- `tb_product`: Product catalog
- `tb_order`: Order management
- `tb_order_item`: Order line items
- `tb_payment`: Payment transactions
- `tb_cart`: Shopping cart
- `tb_cart_item`: Cart line items

## 🧪 Testing

### Run Tests
```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test ./controller/...
```

### Seed Test Data

#### Option 1: Using the Seed Script (Recommended)
```bash
# Make the script executable and run it
chmod +x seed-data.sh
./seed-data.sh
```

#### Option 2: Using cURL
```bash
# Seed all sample data (users and products)
curl -X POST http://localhost:9999/v1/api/seed/all

# Or seed separately
curl -X POST http://localhost:9999/v1/api/seed/users
curl -X POST http://localhost:9999/v1/api/seed/products
```

#### Option 3: Using Admin Dashboard
1. Start the frontend: `cd client && npm run dev`
2. Go to http://localhost:3000
3. Login as admin: `admin@example.com` / `admin123`
4. Navigate to Admin Dashboard
5. Click "Seed Data" button

### Sample Data Created

#### Users
- **Admin User**:
  - Email: `admin@example.com`
  - Password: `admin123`
  - Role: Admin (full access)

- **Customer Users**:
  - Email: `john@example.com` / Password: `customer123`
  - Email: `jane@example.com` / Password: `customer123`
  - Role: Customer (shopping access)

#### Products (8 Sample Items)
 - **iPhone 15 Pro** - KSh 999.99 (50 in stock)
 - **Samsung Galaxy S24** - KSh 899.99 (30 in stock)
 - **MacBook Pro 16-inch** - KSh 2499.99 (20 in stock)
 - **Dell XPS 13** - KSh 1299.99 (25 in stock)
 - **Sony WH-1000XM5** - KSh 399.99 (100 in stock)
 - **iPad Air** - KSh 599.99 (40 in stock)
 - **Nintendo Switch OLED** - KSh 349.99 (60 in stock)
 - **Apple Watch Series 9** - KSh 399.99 (80 in stock)

## 🔧 Configuration

### Environment Variables
```env
# Database
DB_HOST=localhost
DB_PORT=3306
DB_USER=gofiber
DB_PASSWORD=gofiber123
DB_NAME=tech_hive

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379

# JWT
JWT_SECRET_KEY=your-super-secret-key
JWT_EXPIRE_MINUTES_COUNT=60

# Server
SERVER_PORT=9999

# M-Pesa (for simulation)
MPESA_SHORTCODE=174379
MPESA_PASSKEY=your-mpesa-passkey
MPESA_CALLBACK_URL=http://localhost:9999/v1/api/mpesa/callback
```

## 🐳 Docker Services

- **app**: Go Fiber application
- **db**: MySQL database
- **dbtest**: Test database
- **redis**: Redis cache
- **adminer**: Database management
- **redis-commander**: Redis management

## 📁 Project Structure

```
├── client/                 # Frontend application (Vue.js)
│   ├── public/            # Static assets
│   ├── src/
│   │   ├── api/          # API client configuration
│   │   ├── assets/       # Images, styles, icons
│   │   ├── components/   # Reusable Vue components
│   │   ├── composables/  # Vue composables
│   │   ├── middleware/   # Route guards
│   │   ├── router/       # Vue router configuration
│   │   ├── services/     # API service layer
│   │   ├── stores/       # Pinia state management
│   │   ├── types/        # TypeScript type definitions
│   │   ├── utils/        # Utility functions
│   │   └── views/        # Vue pages/views
│   ├── Dockerfile        # Frontend container config
│   └── package.json      # Frontend dependencies
├── common/                 # Shared utilities
├── configuration/          # Configuration management
├── controller/             # HTTP controllers
├── db/migrations/          # Database migrations
├── entity/                 # Domain entities
├── exception/              # Error handling
├── middleware/             # HTTP middleware
├── model/                  # Request/Response models
├── repository/             # Data access layer
│   └── impl/
├── service/                # Business logic layer
│   └── impl/
└── docs/                   # Documentation
```

## 🔒 Security Features

- JWT-based authentication
- Role-based authorization
- Password hashing with bcrypt
- Input validation and sanitization
- CORS configuration
- SQL injection prevention (GORM)
- XSS protection headers

## 🚀 Deployment

### Production Deployment
1. Update environment variables for production
2. Use strong JWT secret key
3. Configure proper M-Pesa credentials
4. Set up reverse proxy (nginx)
5. Configure SSL/TLS certificates
6. Set up monitoring and logging

### Docker Production
```bash
# Build and deploy
docker compose -f docker-compose.prod.yml up -d

# Run migrations
docker compose -f docker-compose.prod.yml exec app migrate -database "$DATABASE_URL" -path db/migrations up
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Run the test suite
6. Submit a pull request

## 📝 API Testing

### Using cURL

#### Register User
```bash
curl -X POST http://localhost:9999/v1/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123",
    "role": "customer"
  }'
```

#### Login
```bash
curl -X POST http://localhost:9999/v1/api/authentication \
  -H "Content-Type: application/json" \
  -d '{
    "username": "test@example.com",
    "password": "password123"
  }'
```

#### Create Product (Admin)
```bash
curl -X POST http://localhost:9999/v1/api/product \
  -H "Authorization: Bearer <admin-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Product",
    "description": "A test product",
    "price": 29.99,
    "stock": 100,
    "image_url": "https://example.com/test.jpg"
  }'
```

## 📊 Monitoring

- Health check endpoint: `GET /health`
- Database monitoring via Adminer
- Redis monitoring via Redis Commander
- Application logs in Docker containers

## 🔧 Troubleshooting

### Common Issues

1. **Database connection failed**
   - Ensure Docker containers are running
   - Check database credentials in environment variables

2. **Migration errors**
   - Verify migration file syntax
   - Check database user permissions

3. **JWT authentication failed**
   - Verify JWT secret key configuration
   - Check token expiration

4. **Redis connection failed**
   - Ensure Redis container is running
   - Check Redis host and port configuration

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- Go Fiber team for the excellent web framework
- GORM team for the powerful ORM
- Docker team for containerization
- All contributors and open source community

---

For more information, visit the [API Documentation](http://localhost:9999/swagger/) once the application is running.
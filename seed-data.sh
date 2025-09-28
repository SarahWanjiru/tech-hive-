#!/bin/bash

# Seed Data Script for Tech Hive E-commerce
# This script seeds the database with sample users and products for testing

echo "🌱 Seeding sample data for Tech Hive E-commerce..."

# Check if the server is running
SERVER_URL="http://localhost:9999"

echo "📡 Checking if server is running on $SERVER_URL..."

if curl -s "$SERVER_URL/health" > /dev/null; then
    echo "✅ Server is running!"
else
    echo "❌ Server is not running on $SERVER_URL"
    echo "Please start the server first:"
    echo "  go run main.go"
    echo "  OR"
    echo "  docker-compose up -d"
    exit 1
fi

echo ""
echo "🔐 Seeding users..."
curl -s -X POST "$SERVER_URL/v1/api/seed/users" -H "Content-Type: application/json" | jq -r '.message // "Users seeded"'

echo ""
echo "📦 Seeding products..."
curl -s -X POST "$SERVER_URL/v1/api/seed/products" -H "Content-Type: application/json" | jq -r '.message // "Products seeded"'

echo ""
echo "🎉 All sample data has been seeded successfully!"
echo ""
echo "📋 Sample Users Created:"
echo "  🔑 Admin User:"
echo "     Email: admin@example.com"
echo "     Password: admin123"
echo ""
echo "  👥 Customer Users:"
echo "     Email: john@example.com"
echo "     Password: customer123"
echo ""
echo "     Email: jane@example.com"
echo "     Password: customer123"
echo ""
echo "🛍️  Sample Products Created:"
echo "  • iPhone 15 Pro - $999.99 (50 in stock)"
echo "  • Samsung Galaxy S24 - $899.99 (30 in stock)"
echo "  • MacBook Pro 16-inch - $2499.99 (20 in stock)"
echo "  • Dell XPS 13 - $1299.99 (25 in stock)"
echo "  • Sony WH-1000XM5 - $399.99 (100 in stock)"
echo "  • iPad Air - $599.99 (40 in stock)"
echo "  • Nintendo Switch OLED - $349.99 (60 in stock)"
echo "  • Apple Watch Series 9 - $399.99 (80 in stock)"
echo ""
echo "🚀 You can now:"
echo "  • Login as admin: admin@example.com / admin123"
echo "  • Login as customer: john@example.com / customer123"
echo "  • Browse products and test the full e-commerce functionality!"
echo ""
echo "💡 Alternative: Use the Admin Dashboard"
echo "  1. Start the frontend: cd client && npm run dev"
echo "  2. Go to http://localhost:3000"
echo "  3. Login as admin"
echo "  4. Click 'Seed Data' button in the Admin Dashboard"
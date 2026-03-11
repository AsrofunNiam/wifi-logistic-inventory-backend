# Login
curl -X POST http://localhost:8089/users/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "password123"}'



# Get all users
curl -X GET http://localhost:8089/users \
  -H "Authorization: Bearer <token>"

# Get user by ID
curl -X GET http://localhost:8089/users/1 \
  -H "Authorization: Bearer <token>"

# Create user
curl -X POST http://localhost:8089/users \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"full_name": "John Doe", "legal_name": "John Doe", "email": "john@example.com", "password": "password123", "number_phone": "08123456789", "role": "admin", "place_of_birth": "Jakarta", "date_of_birth": "1990-01-01", "salary": 5000000}'

# Update user
curl -X PUT http://localhost:8089/users/1 \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"full_name": "John Updated", "email": "john.updated@example.com"}'

# Delete user
curl -X DELETE http://localhost:8089/users/1 \
  -H "Authorization: Bearer <token>"

# Change password
curl -X PUT http://localhost:8089/users/1/password \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"old_password": "oldpass", "new_password": "newpass"}'


# Get all products
curl -X GET http://localhost:8089/products \
  -H "Authorization: Bearer <token>"

# Get product by ID
curl -X GET http://localhost:8089/products/1 \
  -H "Authorization: Bearer <token>"

# Create product
curl -X POST http://localhost:8089/products \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"code": "PRD001", "name": "Router WiFi", "category_id": 1, "supplier_id": 1, "company_code": 1, "description": "High-speed router", "stock": 100, "min_stock": 10, "unit": "Unit", "price": 500000}'

# Update product
curl -X PUT http://localhost:8089/products/1 \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name": "Router WiFi Updated", "stock": 150}'

# Delete product
curl -X DELETE http://localhost:8089/products/1 \
  -H "Authorization: Bearer <token>"

# Get all suppliers
curl -X GET http://localhost:8089/suppliers \
  -H "Authorization: Bearer <token>"

# Get supplier by ID
curl -X GET http://localhost:8089/suppliers/1 \
  -H "Authorization: Bearer <token>"

# Create supplier
curl -X POST http://localhost:8089/suppliers \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"code": "SUP001", "name": "PT Supplier Indonesia", "contact": "John Doe", "phone": "08123456789", "email": "supplier@example.com", "address": "Jakarta, Indonesia", "status": "active"}'

# Update supplier
curl -X PUT http://localhost:8089/suppliers/1 \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name": "PT Supplier Updated", "phone": "08987654321"}'

# Delete supplier
curl -X DELETE http://localhost:8089/suppliers/1 \
  -H "Authorization: Bearer <token>"

# Get all categories
curl -X GET http://localhost:8089/categories \
  -H "Authorization: Bearer <token>"

# Get category by ID
curl -X GET http://localhost:8089/categories/1 \
  -H "Authorization: Bearer <token>"

# Create category
curl -X POST http://localhost:8089/categories \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name": "Network Equipment", "description": "Routers, switches, etc."}'

# Update category
curl -X PUT http://localhost:8089/categories/1 \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name": "Network Devices", "description": "Updated description"}'

# Delete category
curl -X DELETE http://localhost:8089/categories/1 \
  -H "Authorization: Bearer <token>"

# Get all stock in records
curl -X GET http://localhost:8089/stock-in \
  -H "Authorization: Bearer <token>"

# Get stock in by ID
curl -X GET http://localhost:8089/stock-in/1 \
  -H "Authorization: Bearer <token>"

# Create stock in
curl -X POST http://localhost:8089/stock-in \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"code": "SI-20260312-001", "date": "2026-03-12", "product_id": 1, "supplier_id": 1, "quantity": 50, "notes": "Monthly restock"}'

# Update stock in
curl -X PUT http://localhost:8089/stock-in/1 \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"quantity": 60, "notes": "Updated quantity"}'

# Delete stock in
curl -X DELETE http://localhost:8089/stock-in/1 \
  -H "Authorization: Bearer <token>"

# Get all stock out records
curl -X GET http://localhost:8089/stock-out \
  -H "Authorization: Bearer <token>"

# Get stock out by ID
curl -X GET http://localhost:8089/stock-out/1 \
  -H "Authorization: Bearer <token>"

# Create stock out
curl -X POST http://localhost:8089/stock-out \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"code": "SO-20260312-001", "date": "2026-03-12", "product_id": 1, "destination": "Customer A - Jakarta", "quantity": 10, "notes": "Delivery to customer"}'

# Update stock out
curl -X PUT http://localhost:8089/stock-out/1 \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"quantity": 15, "notes": "Updated quantity"}'

# Delete stock out
curl -X DELETE http://localhost:8089/stock-out/1 \
  -H "Authorization: Bearer <token>"

# Get stock summary report
curl -X GET http://localhost:8089/reports/stock-summary \
  -H "Authorization: Bearer <token>"

# Get stock in report (with optional date filters)
curl -X GET "http://localhost:8089/reports/stock-in?start_date=2026-01-01&end_date=2026-03-12" \
  -H "Authorization: Bearer <token>"

# Get stock out report (with optional date filters)
curl -X GET "http://localhost:8089/reports/stock-out?start_date=2026-01-01&end_date=2026-03-12" \
  -H "Authorization: Bearer <token>"

# Get low stock report
curl -X GET http://localhost:8089/reports/low-stock \
  -H "Authorization: Bearer <token>"
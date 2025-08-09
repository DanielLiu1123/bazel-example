# Gateway API Documentation

The Gateway service provides a unified HTTP API for all backend services. It acts as a reverse proxy and handles authentication, routing, and request/response transformation.

## Base URL

```
http://localhost:8080
```

## Authentication

Most endpoints require authentication via JWT token in the Authorization header:

```
Authorization: Bearer <jwt-token>
```

## Endpoints

### Health Check

#### GET /health

Returns the health status of the gateway service.

**Response:**
```json
{
  "status": "healthy",
  "service": "gateway",
  "timestamp": 1640995200
}
```

### Authentication

#### POST /api/auth/login

Authenticate a user and receive a JWT token.

**Request:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "user-123",
    "email": "user@example.com",
    "username": "testuser"
  },
  "service": "user-service"
}
```

### User Management

#### GET /api/users/{id}

Get user details by ID.

**Parameters:**
- `id` (path): User ID

**Response:**
```json
{
  "id": "user-123",
  "email": "user@example.com",
  "username": "testuser",
  "first_name": "John",
  "last_name": "Doe",
  "status": "active",
  "created_at": "2023-01-01T00:00:00Z"
}
```

#### POST /api/users

Create a new user.

**Request:**
```json
{
  "email": "newuser@example.com",
  "username": "newuser",
  "password": "password123",
  "first_name": "Jane",
  "last_name": "Smith",
  "phone": "+1234567890"
}
```

**Response:**
```json
{
  "id": "user-124",
  "email": "newuser@example.com",
  "username": "newuser",
  "status": "created",
  "service": "user-service"
}
```

#### PUT /api/users/{id}

Update user information.

**Parameters:**
- `id` (path): User ID

**Request:**
```json
{
  "first_name": "Jane",
  "last_name": "Doe",
  "phone": "+1234567891"
}
```

**Response:**
```json
{
  "id": "user-123",
  "status": "updated",
  "service": "user-service"
}
```

#### DELETE /api/users/{id}

Delete a user.

**Parameters:**
- `id` (path): User ID

**Response:**
```json
{
  "id": "user-123",
  "status": "deleted",
  "service": "user-service"
}
```

### Order Management

#### GET /api/orders/{id}

Get order details by ID.

**Parameters:**
- `id` (path): Order ID

**Response:**
```json
{
  "id": "order-123",
  "user_id": "user-123",
  "status": "confirmed",
  "total": 99.99,
  "currency": "USD",
  "items": [
    {
      "product_id": "prod-1",
      "product_name": "Test Product",
      "quantity": 2,
      "unit_price": 49.99,
      "total_price": 99.98
    }
  ],
  "created_at": "2023-01-01T00:00:00Z"
}
```

#### POST /api/orders

Create a new order.

**Request:**
```json
{
  "user_id": "user-123",
  "items": [
    {
      "product_id": "prod-1",
      "product_name": "Test Product",
      "quantity": 2,
      "unit_price": 49.99
    }
  ],
  "shipping_address": {
    "street": "123 Main St",
    "city": "Anytown",
    "state": "CA",
    "postal_code": "12345",
    "country": "US"
  }
}
```

**Response:**
```json
{
  "id": "order-124",
  "user_id": "user-123",
  "status": "pending",
  "total": 99.99,
  "service": "order-service"
}
```

#### PUT /api/orders/{id}/status

Update order status.

**Parameters:**
- `id` (path): Order ID

**Request:**
```json
{
  "status": "shipped"
}
```

**Response:**
```json
{
  "id": "order-123",
  "status": "shipped",
  "service": "order-service"
}
```

#### GET /api/orders/user/{userId}

List orders for a specific user.

**Parameters:**
- `userId` (path): User ID
- `page` (query, optional): Page number (default: 1)
- `page_size` (query, optional): Items per page (default: 10)

**Response:**
```json
{
  "orders": [
    {
      "id": "order-123",
      "user_id": "user-123",
      "status": "confirmed",
      "total": 99.99
    }
  ],
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total_items": 1,
    "total_pages": 1
  },
  "service": "order-service"
}
```

## Error Responses

All endpoints may return error responses in the following format:

```json
{
  "error": {
    "code": "INVALID_ARGUMENT",
    "message": "Email is required",
    "details": "Validation failed for field: email"
  }
}
```

### Error Codes

- `INVALID_ARGUMENT`: Invalid request parameters
- `NOT_FOUND`: Resource not found
- `ALREADY_EXISTS`: Resource already exists
- `PERMISSION_DENIED`: Access denied
- `UNAUTHENTICATED`: Authentication required
- `INTERNAL`: Internal server error
- `UNAVAILABLE`: Service unavailable

## Rate Limiting

The gateway implements rate limiting:
- **Default**: 100 requests per minute per IP
- **Authenticated users**: 1000 requests per minute per user

Rate limit headers are included in responses:
```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1640995260
```

## CORS

The gateway supports CORS for web applications:
- **Allowed Origins**: Configurable (default: localhost:3000)
- **Allowed Methods**: GET, POST, PUT, DELETE, OPTIONS
- **Allowed Headers**: Content-Type, Authorization

## WebSocket Support

The gateway supports WebSocket connections for real-time features:

```
ws://localhost:8080/ws
```

Authentication is required via query parameter:
```
ws://localhost:8080/ws?token=<jwt-token>
```

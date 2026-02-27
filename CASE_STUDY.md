# Case Study: Restaurant Menu Discovery API

## Overview

**FoodHub** is a mobile and web application that helps users discover and view restaurant menus in their city. Users can browse restaurants by location, see detailed menu items with prices, and check availability. Restaurant managers can add, update, and manage their menu items through the API.

This case study demonstrates a RESTful API for managing restaurant data and menu information. The system features two interconnected entities: **Restaurants** and **Menu Items**, allowing users to explore dining options without needing authentication for basic browsing. It showcases all CRUD operations in a real-world scenario.

## Data Models

### Restaurant Table
```
Restaurant {
  id: UUID (auto-generated)
  name: string (required)
  location: string (required)
  created_at: timestamp (auto-generated)
}
```

### MenuItem Table
```
MenuItem {
  id: UUID (auto-generated)
  restaurant_id: UUID (foreign key)
  name: string (required)
  price: decimal (required)
  available: boolean (default: true)
  created_at: timestamp (auto-generated)
  updated_at: timestamp (auto-generated)
}
```

## API Endpoints

### Restaurants

#### Create Restaurant (POST)
**Endpoint:** `POST /api/v1/restaurants`

**Request:**
```json
{
  "name": "Pasta House",
  "location": "Jakarta"
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Pasta House",
  "location": "Jakarta",
  "created_at": "2024-01-15T10:30:00Z"
}
```

---

#### Get All Restaurants (GET)
**Endpoint:** `GET /api/v1/restaurants?page=1&limit=10`

**Query Parameters:**
- `page` - Page number (default: 1)
- `limit` - Items per page (default: 10, max: 50)

**Response (200 OK):**
```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "Pasta House",
      "location": "Jakarta",
      "created_at": "2024-01-15T10:30:00Z"
    },
    {
      "id": "550e8400-e29b-41d4-a716-446655440001",
      "name": "Sushi Place",
      "location": "Jakarta",
      "created_at": "2024-01-15T11:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 25,
    "total_pages": 3
  }
}
```

---

#### Get Restaurant by ID (GET)
**Endpoint:** `GET /api/v1/restaurants/{id}`

**Response (200 OK):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Pasta House",
  "location": "Jakarta",
  "created_at": "2024-01-15T10:30:00Z"
}
```

---

#### Update Restaurant (PUT/PATCH)
**Endpoint:** `PUT /api/v1/restaurants/{id}` or `PATCH /api/v1/restaurants/{id}`

**Request:**
```json
{
  "name": "Pasta House Premium",
  "location": "Jakarta"
}
```

**Response (200 OK):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Pasta House Premium",
  "location": "Jakarta",
  "created_at": "2024-01-15T10:30:00Z"
}
```

---

#### Delete Restaurant (DELETE)
**Endpoint:** `DELETE /api/v1/restaurants/{id}`

**Response (204 No Content)**

---

### Menu Items

#### Create Menu Item (POST)
**Endpoint:** `POST /api/v1/restaurants/{restaurant_id}/items`

**Request:**
```json
{
  "name": "Carbonara",
  "price": 85000
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "restaurant_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Carbonara",
  "price": 85000,
  "available": true,
  "created_at": "2024-01-15T10:45:00Z",
  "updated_at": "2024-01-15T10:45:00Z"
}
```

---

#### Get All Items for Restaurant (GET)
**Endpoint:** `GET /api/v1/restaurants/{restaurant_id}/items?page=1&limit=20`

**Query Parameters:**
- `page` - Page number (default: 1)
- `limit` - Items per page (default: 20, max: 50)

**Response (200 OK):**
```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440002",
      "restaurant_id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "Carbonara",
      "price": 85000,
      "available": true,
      "created_at": "2024-01-15T10:45:00Z",
      "updated_at": "2024-01-15T10:45:00Z"
    },
    {
      "id": "550e8400-e29b-41d4-a716-446655440003",
      "restaurant_id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "Bolognese",
      "price": 75000,
      "available": true,
      "created_at": "2024-01-15T10:50:00Z",
      "updated_at": "2024-01-15T10:50:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 45,
    "total_pages": 3
  }
}
```

---

#### Get Menu Item by ID (GET)
**Endpoint:** `GET /api/v1/restaurants/{restaurant_id}/items/{item_id}`

**Response (200 OK):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "restaurant_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Carbonara",
  "price": 85000,
  "available": true,
  "created_at": "2024-01-15T10:45:00Z",
  "updated_at": "2024-01-15T10:45:00Z"
}
```

---

#### Update Menu Item (PUT/PATCH)
**Endpoint:** `PUT /api/v1/restaurants/{restaurant_id}/items/{item_id}` or `PATCH /api/v1/restaurants/{restaurant_id}/items/{item_id}`

**Request (PATCH - Mark unavailable):**
```json
{
  "available": false
}
```

**Response (200 OK):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "restaurant_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Carbonara",
  "price": 85000,
  "available": false,
  "created_at": "2024-01-15T10:45:00Z",
  "updated_at": "2024-01-15T11:30:00Z"
}
```

---

#### Delete Menu Item (DELETE)
**Endpoint:** `DELETE /api/v1/restaurants/{restaurant_id}/items/{item_id}`

**Response (204 No Content)**

---

## Database Schema

```sql
CREATE TABLE restaurants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    location VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE menu_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    restaurant_id UUID NOT NULL REFERENCES restaurants(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    available BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

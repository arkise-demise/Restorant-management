# Restaurant Management System

This is a RESTful API for managing a restaurant's operations, built with the Go programming language and the Gin framework. It interacts with a MongoDB database to manage data related to users, food items, menus, tables, orders, order items, and invoices.

## Features

- User authentication and management
- CRUD operations for food items, menus, tables, orders, order items, and invoices
- Middleware for authentication and logging
- MongoDB integration

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16 or later
- MongoDB

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/arkise-demise/restaurant-management.git
   cd restaurant-management
   ```

2. Install Go dependencies:
   ```sh
   go mod tidy
   ```

### Running the Application

1. Start your MongoDB server.

2. Run the application:

   ```sh
   go run main.go
   ```

   The server will start on `http://localhost:8000`.

## API Endpoints

Here are some of the main endpoints available in the API:

### User Routes

- `POST /users/signup` - Register a new user
- `POST /users/login` - Login a user

### Food Routes

- `GET /foods` - Get all food items
- `POST /foods` - Create a new food item
- `GET /foods/:id` - Get a food item by ID
- `PUT /foods/:id` - Update a food item by ID
- `DELETE /foods/:id` - Delete a food item by ID

### Menu Routes

- `GET /menus` - Get all menus
- `POST /menus` - Create a new menu
- `GET /menus/:id` - Get a menu by ID
- `PUT /menus/:id` - Update a menu by ID
- `DELETE /menus/:id` - Delete a menu by ID

### Table Routes

- `GET /tables` - Get all tables
- `POST /tables` - Create a new table
- `GET /tables/:id` - Get a table by ID
- `PUT /tables/:id` - Update a table by ID
- `DELETE /tables/:id` - Delete a table by ID

### Order Routes

- `GET /orders` - Get all orders
- `POST /orders` - Create a new order
- `GET /orders/:id` - Get an order by ID
- `PUT /orders/:id` - Update an order by ID
- `DELETE /orders/:id` - Delete an order by ID

### Order Item Routes

- `GET /orderitems` - Get all order items
- `POST /orderitems` - Create a new order item
- `GET /orderitems/:id` - Get an order item by ID
- `PUT /orderitems/:id` - Update an order item by ID
- `DELETE /orderitems/:id` - Delete an order item by ID

### Invoice Routes

- `GET /invoices` - Get all invoices
- `POST /invoices` - Create a new invoice
- `GET /invoices/:id` - Get an invoice by ID
- `PUT /invoices/:id` - Update an invoice by ID
- `DELETE /invoices/:id` - Delete an invoice by ID

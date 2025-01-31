# GoQuick - E-commerce Backend in Go

GoQuick is a high-performance, scalable e-commerce backend built with **Golang**, **GORM**, and **JWT authentication**. It follows clean architecture principles, implements structured logging with **Zap**, and ensures modularity with dependency inversion.

---

## 🚀 Features

✅ **User Authentication** (JWT-based login, registration, and role-based access control) ✅ **Product Management** (CRUD operations for products) ✅ **Cart & Orders** (Manage shopping cart and process orders) ✅ **Structured Logging** (Using Uber Zap for production-ready logging) ✅ **Database Integration** (PostgreSQL with GORM ORM) ✅ **Middleware Support** (JWT authentication & role-based access control) ✅ **Clean Architecture** (Separation of concerns: handlers, services, repositories) ✅ **Scalability & Maintainability** (Dependency inversion, modular structure)

---

## 🏗️ Architecture

The project follows a **layered architecture**:

```
GoQuick/
│── internal/
│   ├── handlers/        # HTTP handlers
│   ├── services/        # Business logic
│   ├── models/          # Database models
│   ├── middleware/      # Authentication & authorization
│   ├── logger/          # Structured logging (Zap)
│   ├── config/          # Configuration management
│── db/
│   ├── migrations/      # Database migration scripts
│── main.go              # Entry point
```

---

## 🛠️ Tech Stack

- **Language:** Golang
- **Frameworks/Libraries:** Gorilla Mux, GORM, Zap, JWT-Go
- **Database:** PostgreSQL
- **Authentication:** JWT (JSON Web Tokens)
- **Logging:** Uber Zap
- **Dependency Management:** Go Modules

---

## 🔧 Installation & Setup

### Prerequisites

- Install [Go](https://go.dev/dl/)
- Install [PostgreSQL](https://www.postgresql.org/download/)

### Clone the Repository

```sh
$ git clone https://github.com/Namith667/GoQuick.git
$ cd GoQuick
```

### Configure Environment Variables

Create a `.env` file in the root directory and update the values:

```sh
DB_HOST=localhost
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=goquick
JWT_SECRET_KEY=your_secret_key
```

### Run Database Migrations

```sh
$ go run db/migrations/migrate.go
```

### Start the Server

```sh
$ go run main.go
```

---

## 📌 API Endpoints

### **Authentication**

#### **User Registration**

`POST /api/auth/register`

```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "securepassword",
  "role": "customer"
}
```

#### **User Login**

`POST /api/auth/login`

```json
{
  "email": "john@example.com",
  "password": "securepassword"
}
```

*Response:*

```json
{
  "token": "your_jwt_token_here"
}
```

### **Products**

#### **Get All Products**

`GET /api/products`

#### **Create Product (Admin Only)**

`POST /api/products`

```json
{
  "name": "Laptop",
  "price": 1200.00,
  "stock": 10
}
```

### **Cart & Orders**

#### **Add to Cart**

`POST /api/cart`

```json
{
  "product_id": 1,
  "quantity": 2
}
```

#### **Place Order**

`POST /api/orders`

```json
{
  "cart_id": 123
}
```

---

## 📂 Folder Structure

```
GoQuick/
│── internal/
│   ├── handlers/        # HTTP handlers
│   ├── services/        # Business logic
│   ├── models/          # Database models
│   ├── middleware/      # Authentication & authorization
│   ├── logger/          # Structured logging (Zap)
│   ├── config/          # Configuration management
│── db/
│   ├── migrations/      # Database migration scripts
│── main.go              # Entry point
```

---

## 🛠️ Future Improvements

🚀 **Dockerization** - Containerize the app for easy deployment 🚀 **Unit & Integration Testing** - Improve test coverage 🚀 **GraphQL API** - Add support for GraphQL 🚀 **CI/CD Pipeline** - Automate deployment with GitHub Actions

---

## 💡 Contributing

Contributions are welcome! Feel free to **fork the repository**, create a new branch, and submit a pull request.

---

## 📜 License

MIT License. See `LICENSE` for more details.

---

### ⭐ Star the Repository if You Like It!

If you found this useful, consider giving it a ⭐ on [GitHub](https://github.com/Namith667/GoQuick)!


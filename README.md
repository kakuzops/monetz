# Monetz

The project was an study to understand more golang and use some patterns some ideias and create an Sass for a friends.
The project idea was create an api with LLM to underst the rent, payments and catalog this.


A Go-based REST API for user management using Gin framework and GORM.

## 🔧 Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/monetz.git
cd monetz
```

2. Install dependencies:
```bash
go mod download
```

3. Create your environment file:
```bash
cp .env.example .local.env
```

4. Update the `.local.env` file with your database credentials:
```env
DB_HOST=HOST
DB_PORT=PORT
DB_USER=USER
DB_PASSWORD=PASSWORD
DB_NAME=DB_NAME
```

## 🚀 Running the Application

```bash
go run main.go
```

The server will start at `http://localhost:8080`

## 🛣️ API Endpoints

### Users
- `GET /user` - Get all users
- `POST /user` - Create a new user
- `GET /user/:id` - Get user by ID

### Health Check
- `GET /health` - Check API and database health

## 📦 Project Structure

```
monetz/
├── main.go
├── src/
│   ├── app/
│   │   └── user/
│   │       ├── handler/
│   │       ├── service/
│   │       └── repositories/
│   └── config/
│       ├── database/
│       └── models/
└── .env.example
```


This project is licensed under the MIT License - see the LICENSE file for details.
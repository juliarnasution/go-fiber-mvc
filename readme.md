```bash
go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/go-playground/validator/v10
go get golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt/v4
go get github.com/gofiber/template
```

```bash
go run main.go
```
```bash
# API Routes
POST /register - Register a new user
POST /login - Login a user
GET /users - List all users
GET /users/:id - Get a specific user by ID
POST /users - Create a new user
PUT /users/:id - Update an existing user by ID
DELETE /users/:id - Delete a user by ID
```
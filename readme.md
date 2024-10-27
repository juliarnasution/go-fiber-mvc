go mod init go-fiber-mvc
go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/go-playground/validator/v10
go get golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt/v4



go run main.go


GET /users - List all users
GET /users/:id - Get a specific user by ID
POST /users - Create a new user
PUT /users/:id - Update an existing user by ID
DELETE /users/:id - Delete a user by ID
package middlewares

import (
    "fmt"
    "os"
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware checks for a valid JWT token in the Authorization header
func AuthMiddleware(c *fiber.Ctx) error {
    // Get token from the Authorization header
    tokenString := c.Get("Authorization")

    if tokenString == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Missing or malformed JWT",
        })
    }

    // Parse and validate the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Verify the signing method is what you expect
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        // Return the secret key
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil || !token.Valid {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Invalid or expired JWT",
        })
    }

    return c.Next()
}

// Function to generate JWT token (e.g., after a successful login)
func GenerateJWT(userID uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expiry
    })

    // Generate encoded token and send it as response.
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

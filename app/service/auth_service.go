package service

import (
    "os"
    "time"
    "prestasi_backend/app/repository"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

// Repositories dipanggil dari sini agar tidak nil
var authUserRepo = repository.NewUserRepository()
var authRolePermRepo = repository.NewRolePermissionRepository()

// Generate JWT Token
func generateToken(userID string, roleID string, permissions []string) (string, error) {
    claims := jwt.MapClaims{
        "user_id":     userID,
        "role_id":     roleID,
        "permissions": permissions,
        "exp":         time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// LOGIN API
func AuthLogin(c *fiber.Ctx) error {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }

    user, err := authUserRepo.FindByUsername(req.Username)
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"error": "Invalid username or password"})
    }

    // Password check
    if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
        return c.Status(401).JSON(fiber.Map{"error": "Invalid username or password"})
    }

    permissions, _ := authRolePermRepo.GetPermissions(user.RoleID)

    token, err := generateToken(user.ID, user.RoleID, permissions)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
    }

    return c.JSON(fiber.Map{
        "token": token,
        "user": fiber.Map{
            "id":       user.ID,
            "username": user.Username,
            "email":    user.Email,
            "role_id":  user.RoleID,
        },
    })
}

// PROFILE
func AuthProfile(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "user_id":     c.Locals("user_id"),
        "role_id":     c.Locals("role_id"),
        "permissions": c.Locals("permissions"),
    })
}

// REFRESH TOKEN endpoint
func AuthRefresh(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "refresh not implemented"})
}

// LOGOUT endpoint
func AuthLogout(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "logout success"})
}

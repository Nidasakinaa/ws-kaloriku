package handler

import (
	"fmt"
	"net/http"
	"strings"

	iniconfig "github.com/Nidasakinaa/be_KaloriKu/config"
	"github.com/Nidasakinaa/be_KaloriKu/model"
	inimodel "github.com/Nidasakinaa/be_KaloriKu/model"
	cek "github.com/Nidasakinaa/be_KaloriKu/module"
	"github.com/Nidasakinaa/ws-kaloriku/config"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var loginDetails inimodel.User
	if err := c.BodyParser(&loginDetails); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid request",
		})
	}

	storedAdmin, err := cek.GetRoleByAdmin(config.Ulbimongoconn, "User", "Admin")
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status":  http.StatusUnauthorized,
			"message": "Invalid credentials",
		})
	}

	if !iniconfig.CheckPasswordHash(loginDetails.Password, storedAdmin.Password) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status":  http.StatusUnauthorized,
			"message": "Invalid credentials",
		})
	}

	token, err := iniconfig.GenerateJWT(*storedAdmin, "Admin") 
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Could not generate token",
		})
	}

	err = cek.SaveTokenToDatabase(config.Ulbimongoconn, "tokens", storedAdmin.ID.Hex(), token)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Could not save token",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Login successful",
		"token":   token,
	})
}


func Logout(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing token",
		})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token format",
		})
	}

	token := parts[1]

	err := cek.DeleteTokenFromMongoDB(config.Ulbimongoconn, "tokens", token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not delete token",
		})
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successful",
	})
}

func Register(c *fiber.Ctx) error {
    var newAdmin model.User
    if err := c.BodyParser(&newAdmin); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "status":  http.StatusBadRequest,
            "message": "Invalid request body",
        })
    }


    // Check if the username already exists for the specified role (Admin or Customer)
    existingUser, err := cek.GetByUsername(config.Ulbimongoconn, "User", newAdmin.Username)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "status":  http.StatusInternalServerError,
            "message": "Could not check existing username",
        })
    }

    if existingUser != nil {
        return c.Status(http.StatusConflict).JSON(fiber.Map{
            "status":  http.StatusConflict,
            "message": "Username already taken",
        })
    }

    hashedPassword, err := iniconfig.HashPassword(newAdmin.Password)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "status":  http.StatusInternalServerError,
            "message": "Could not hash password",
        })
    }

    // Save new user to the database
    insertedID, err := cek.InsertUser(config.Ulbimongoconn, "User", newAdmin.FullName, newAdmin.Phone, newAdmin.Username, hashedPassword, newAdmin.Role)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "status":  http.StatusInternalServerError,
            "message": "Could not register user",
        })
    }

    return c.Status(http.StatusCreated).JSON(fiber.Map{
        "status":  http.StatusCreated,
        "message": "Account registered successfully",
        "data":    insertedID,
    })
}


func DashboardPage(c *fiber.Ctx) error {
    adminID := c.Locals("admin_id")
    if adminID == nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "status":  http.StatusInternalServerError,
            "message": "Admin ID not found in context",
        })
    }

    adminIDStr := fmt.Sprintf("%v", adminID)

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "status":  http.StatusOK,
        "message": "Dashboard access successful",
        "admin_id": adminIDStr,
    })
}	

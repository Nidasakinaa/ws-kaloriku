package controller

import (
	"errors"
	"fmt"
	"net/http"

	inimodel "github.com/Nidasakinaa/be_KaloriKu/model"
	cek "github.com/Nidasakinaa/be_KaloriKu/module"
	"github.com/Nidasakinaa/ws-kaloriku/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

// GetMenuID godoc
// @Summary Get By ID Data Menu.
// @Description Ambil per ID data menu.
// @Tags MenuItem
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} MenuItem
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /menu/{id} [get]
func GetMenuID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := cek.GetMenuItemByID(objID, config.Ulbimongoconn, "Proyek3")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

// GetMenu godoc
// @Summary Get All Data Menu.
// @Description Mengambil semua data menu.
// @Tags MenuItem
// @Accept json
// @Produce json
// @Success 200 {object} MenuItem
// @Router /menu [get]
//GetAllMenuItem retrieves all menu items from the database
func GetMenu(c *fiber.Ctx) error {
	ps := cek.GetAllMenuItem(config.Ulbimongoconn, "Menu")
	return c.JSON(ps)
}

// InsertDataMenu godoc
// @Summary Insert data menu.
// @Description Input data menu.
// @Tags MenuItem
// @Accept json
// @Produce json
// @Param request body ReqMenuItem true "Payload Body [RAW]"
// @Success 200 {object} MenuItem
// @Failure 400
// @Failure 500
// @Router /insert [post]
func InsertDataMenu(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var menuItem inimodel.MenuItem
	if err := c.BodyParser(&menuItem); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := cek.InsertMenuItem(db, "Menu",
		menuItem.Name,
		menuItem.Description,
		menuItem.Category,
		menuItem.Image) // Assuming menuItem has a Calories field of type float64
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
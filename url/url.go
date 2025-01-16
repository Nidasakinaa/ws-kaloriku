package url

import (
	"github.com/Nidasakinaa/ws-kaloriku/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage)
	page.Get("/menu", controller.GetMenu)
	page.Get("/menu/:id", controller.GetMenuID)
	page.Post("/insert", controller.InsertDataMenu)
	// page.Put("/update/:id", controller.UpdateDataMenu)
	// page.Delete("/delete/:id", controller.DeleteMenuByID)

	// page.Get("/docs/*", swagger.HandlerDefault)
}

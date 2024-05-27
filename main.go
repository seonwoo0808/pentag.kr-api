package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"pentag.kr/api-server/configs"
	"pentag.kr/api-server/database"
	"pentag.kr/api-server/dto"
	"pentag.kr/api-server/utils"
)	



func main() {
	configs.LoadAllEnv()
	database.NewDBClient()	

    app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://pentag.kr",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

    app.Post("/contact", func (c *fiber.Ctx) error {
		var contactReq dto.ContactReq

		if err := c.BodyParser(&contactReq); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		if err := contactReq.Validate(); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		
		if !utils.CheckCaptcha(contactReq.Token) {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid captcha",
			})
		}

		contactReq.ReplaceSymbol()

		dbClient := database.GetDBClient()
		contactObj, err := dbClient.Contact.Create().
			SetFirstName(contactReq.FirstName).
			SetLastName(contactReq.LastName).
			SetEmail(contactReq.Email).
			SetPhone(contactReq.Phone).
			SetCategory(contactReq.Category).
			SetMessage(contactReq.Message).
			SetVerifyCode(utils.GenerateRandomString(10)).
			Save(context.Background())
		

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Cannot save contact",
			})
		}

		err = contactReq.SendVerifyEmail(contactObj.VerifyCode)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Cannot send email",
			})
		}

        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":3000"))
}
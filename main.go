package main


import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)
// fmt.Println("Hello, World!") => with JavaScript; console.log("Hello, World!")

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName: "Fiber Starter App",
	})


	/*
	In: 	EXPRESS.JS      |     GO
	-------------------------------------------
			app.delete()    |  app.Delete()
			app.get()       |  app.Get()
			app.put()       |  app.Put()
			app.patch()     |  app.Patch()
			app.post()      |  app.Post()
			next()          |  c.Next() (c *fiber.Ctx)
	-------------------------------------------
		
	*/

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Emotional Damage!")
		// http://localhost:3000/
		/*
			"Emotional Damage!"
		*/
	})

	app.Get("/@:id", func(c *fiber.Ctx) error {
		result := utils.ImmutableString(c.Params("id")) 
		return c.SendString(result)
		// http://localhost:3000/@clqu
		/*
			"clqu"
		*/
	})

	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
		// http://localhost:3000/api/users/premium
		/*
			API path: users/premium
		*/
	})

	app.Get("/page::page/filter::filter", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Page": c.Params("page"),"Filter": c.Params("filter")})
		// http://localhost:3000/page:1/filter:safety
		/*
			{
				"Filter": "safety",
				"Page": "1"
			}
		*/
	})


	app.Listen(":3000")
}
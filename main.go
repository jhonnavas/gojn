package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	//Instanciando el framework
	app := fiber.New()

	app.Use(cors.New())

	//Desde Static le vamos a decir que cuando visite la ruta inicial, despache la carpeta public
	app.Static("/", "./public")

	//Crear una ruta de GO, quiero que cuando le hagan una peticion GET quiero que crees un context
	/*
		Esta funcion recibe un objeto que tiene informacion de la peticion y la respuesta "func(c *fiber.Ctx)"
		es como el request y response solo que esta unido en uno solo
	*/
	//Esta funcion puede retornar un error o devolver algo al cliente
	app.Get("/users", func(c *fiber.Ctx) error {
		//Vamos a utilizar del contexto, la funcion JSON y con ello le enviamos un objeto
		//Fiber.map es un objeto que da el framework, el cual es un struct al que le puedo decir el dato que quiero enviarle
		return c.JSON(&fiber.Map{
			"data": "Estos son los usuarios desde el backend",
		})
	})

	app.Listen(":3000")
	fmt.Println("Server on port 3000")

}

package main

import (
	"zanuarmirza/goFiberSample/handler"
	"zanuarmirza/goFiberSample/model"
	"zanuarmirza/goFiberSample/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//  --- CARGO CONFIG ---
	cargo := model.Cargo{
		MaxWeigth: 20,
		Volume: model.Volume{
			Height: 10,
			Width:  10,
			Length: 10,
		},
		CurrentWeigth: 0,
		CurrentVolume: 0,
		Temperature: model.CargoTemperature{
			MinTemp: 20,
			MaxTemp: 30,
		},
		Door: model.Door{
			Height: 7,
			Width:  5,
		},
	}
	cargoService := service.NewCargoService(cargo)
	handler := handler.NewHandler(cargoService)
	app.Post("/validate", handler.InputItem)

	app.Listen(":3000")
}

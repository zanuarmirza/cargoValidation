package handler

import (
	"fmt"
	"zanuarmirza/goFiberSample/model"
	"zanuarmirza/goFiberSample/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	CargoService *service.CargoService
}

func NewHandler(cargoService *service.CargoService) *Handler {
	return &Handler{cargoService}
}

// CopyHandler ... , copy params to buffer then print it
func (h *Handler) InputItem(c *fiber.Ctx) error {
	// reset cargo variable
	h.CargoService.Cargo.RegisteredItems = nil
	h.CargoService.Cargo.CurrentVolume = 0
	h.CargoService.Cargo.CurrentWeigth = 0

	itemInput := new(model.ItemInputRequest)
	if err := c.BodyParser(itemInput); err != nil {
		fmt.Printf("%v\n", "InputItem parsing error")
		return err
	}
	fmt.Println("----------------------")
	for _, item := range itemInput.ListItem {
		if h.CargoService.InputItem(item) {
			fmt.Printf("%v:%v\n", item.ID, "PASS")
		} else {
			fmt.Printf("%v:%v\n", item.ID, "REJECT")
		}
	}
	fmt.Println("Item in Cargo : ")
	for _, rItem := range h.CargoService.Cargo.RegisteredItems {
		fmt.Println(rItem.ID)
	}
	fmt.Println("----------------------")
	// print item in cargo

	c.JSON(h.CargoService.Cargo.RegisteredItems)
	return nil
}

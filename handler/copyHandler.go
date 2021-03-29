package handler

import (
	"fmt"
	"zanuarmirza/goFiberSample/model"

	"github.com/gofiber/fiber/v2"
)

// CopyHandler ... , copy params to buffer then print it
func CopyHandler(c *fiber.Ctx) error {
	// Variable is only valid within this handler
	result := c.Params("foo")

	// Make a copy
	buffer := make([]byte, len(result))
	copy(buffer, result)
	resultCopy := string(buffer)
	// Variable is now valid forever
	fmt.Printf("%v\n", resultCopy)
	return nil
}

func WriteHandler(c *fiber.Ctx) error {
	p := new(model.Person)
	if err := c.BodyParser(p); err != nil {
		fmt.Printf("%v\n", "write Handler parsing error")
		return err
	}
	return c.Status(fiber.StatusOK).JSON(p)
}

package web

import (
	"ylanzinhoy-operator-management/internal/application/useCase/operatorUseCase"

	"github.com/gofiber/fiber/v2"
)

type WebHandlerAdapter struct {
	uc operatorUseCase.Port
}

func NewWebAdapter(uc operatorUseCase.Port) *WebHandlerAdapter {
	return &WebHandlerAdapter{
		uc: uc,
	}
}

func (w *WebHandlerAdapter) PostOperator(c *fiber.Ctx) error {
	requestCtx := c.Context()

	inputDto := operatorUseCase.InsertOperatorDTO{}

	if err := c.BodyParser(&inputDto); err != nil {
		return err
	}

	resultId, err := w.uc.InsertOperator(requestCtx, inputDto)

	if err != nil {
		return err
	}

	if resultId == nil {
		c.Status(500)
	}

	err = c.Status(201).JSON(&fiber.Map{
		"id": resultId,
	})

	if err != nil {
		return err
	}

	return nil
}

func (w *WebHandlerAdapter) GetOperators(c *fiber.Ctx) error {
	requestCtx := c.Context()

	operatorList, err := w.uc.ListOperators(requestCtx)

	if err != nil {
		return err
	}
	return c.Status(200).JSON(operatorList)
}

func (w *WebHandlerAdapter) GetOperatorById(c *fiber.Ctx) error {
	requestCtx := c.Context()
	id := c.Params("id")

	operator, err := w.uc.GetOperatorById(requestCtx, id)
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"error": "operator not found",
		})
	}

	return c.Status(200).JSON(operator)
}

func (w *WebHandlerAdapter) UpdateOperator(c *fiber.Ctx) error {
	requestCtx := c.Context()
	id := c.Params("id")

	inputDto := operatorUseCase.UpdateOperatorDTO{}

	if err := c.BodyParser(&inputDto); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": "invalid request body",
		})
	}

	err := w.uc.UpdateOperator(requestCtx, id, inputDto)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "operator updated successfully",
	})
}

func (w *WebHandlerAdapter) DeleteOperator(c *fiber.Ctx) error {
	requestCtx := c.Context()
	id := c.Params("id")

	err := w.uc.DeleteOperator(requestCtx, id)
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"error": "operator not found",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "operator deleted successfully",
	})
}

func (w *WebHandlerAdapter) ChangePassword(c *fiber.Ctx) error {
	requestCtx := c.Context()

	inputDto := operatorUseCase.UpdatePasswordDTO{}

	if err := c.BodyParser(&inputDto); err != nil {
		c.Status(400)
		return err
	}

	err := w.uc.UpdatePassword(requestCtx, inputDto)
	if err != nil {
		return err
	}

	return nil

}

package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/tracker"
)

type UpdateItemRequest struct {
	Name string `json:"name"`
}

func (s *Server) UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	var req UpdateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}

	item, err := s.Repository.Update(c.Context(), tracker.Item{ID: id, Name: req.Name})
	if err != nil {
		if errors.Is(err, tracker.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "item not found")
		}
		log.Errorw("s.Repository.Update", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

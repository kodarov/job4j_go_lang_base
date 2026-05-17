package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/tracker"
)

func (s *Server) DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	err := s.Repository.Delete(c.Context(), id)
	if err != nil {
		if errors.Is(err, tracker.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "item not found")
		}
		log.Errorw("s.Repository.Delete", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/tracker"
)

func (s *Server) GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	item, err := s.Repository.Get(c.Context(), id)
	if err != nil {
		if errors.Is(err, tracker.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "item not found")
		}
		log.Errorw("s.Repository.Get", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

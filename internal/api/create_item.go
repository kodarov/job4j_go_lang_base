package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"job4j.ru/go-lang-base/internal/tracker"
)

type CreateItemRequest struct {
	Name string `json:"name"`
}

func (s *Server) CreateItem(c *fiber.Ctx) error {
	var req CreateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}

	item := tracker.Item{
		ID:   uuid.New().String(),
		Name: req.Name,
	}

	err := s.Repository.Create(c.Context(), item)
	if err != nil {
		log.Errorw("s.Repository.Create", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusCreated).JSON(item)
}

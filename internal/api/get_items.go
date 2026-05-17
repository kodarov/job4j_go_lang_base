package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/tracker"
)

type ItemRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FindItemsResponse struct {
	Items []ItemRequest `json:"items"`
}

func (s *Server) FindItems(c *fiber.Ctx) error {
	name := c.Query("name")
	var items []tracker.Item
	var err error

	if name != "" {
		items, err = s.Repository.FindByName(c.Context(), name)
	} else {
		items, err = s.Repository.List(c.Context())
	}
	if err != nil {
		log.Errorw("s.Repository.List", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	res := make([]ItemRequest, 0, len(items))
	for _, item := range items {
		res = append(res, ItemRequest{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return c.Status(fiber.StatusOK).JSON(FindItemsResponse{Items: res})
}

package api

import "github.com/gofiber/fiber/v2"

func (s *Server) Route(route fiber.Router) {
	route.Post("/items/", s.CreateItem)
	route.Get("/items/", s.FindItems)
	route.Get("/items/:id", s.GetItem)
	route.Patch("/items/:id", s.UpdateItem)
	route.Delete("/items/:id", s.DeleteItem)
}

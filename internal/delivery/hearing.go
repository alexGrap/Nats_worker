package delivery

import (
	"L0/config"
	models "L0/internal"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Server struct {
	Server  *fiber.App
	UseCase models.UseCase
	Nats    models.Nats
}

func Fabric(conf *config.Config, useCase models.UseCase) models.Rest {
	server := Server{}
	server.Server = fiber.New()
	server.Server.Get("getById", server.GetOne)
	server.Server.Get("getAll", server.GetAll)
	server.UseCase = useCase
	return &server
}

func (server *Server) Hearing() error {
	err := server.Server.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (server *Server) GetOne(ctx *fiber.Ctx) error {
	key := ctx.Query("id")
	body, err := server.UseCase.Get(key)
	if err != nil {
		ctx.Status(404)
		return ctx.JSON(err)
	}
	return ctx.Send(body)
}

func (server *Server) GetAll(ctx *fiber.Ctx) error {
	cash, err := server.UseCase.GetAll()
	if err != nil {
		ctx.Status(404)
		return ctx.JSON(err)
	}
	var result []models.Model
	var model models.Model
	for i, _ := range cash {
		err = json.Unmarshal(cash[i].Body, &model)
		if err != nil {
			ctx.Status(500)
			return ctx.JSON(err)
		}
		result = append(result, model)
	}

	return ctx.JSON(result)
}

package http

import (
	"context"

	"github.com/andreoav/mock-helper-api/pkg/domain"
	"github.com/gofiber/fiber"
)

type FiberProjectHandler struct {
	projectService domain.ProjectService
}

func NewFiberProjectHandler(app *fiber.App, ps domain.ProjectService) {
	handler := &FiberProjectHandler{
		projectService: ps,
	}

	app.Get("/project", handler.FetchProject)
}

func (h *FiberProjectHandler) FetchProject(ctx *fiber.Ctx) {
	basePath := ctx.Query("basePath")

	project, err := h.projectService.FetchByBasePath(context.Background(), basePath)

	if err != nil {
		ctx.SendStatus(400)
		ctx.JSON(fiber.Map{"message": err.Error()})
		return
	}

	ctx.JSON(project)
}

package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/andreoav/mock-helper-api/pkg/domain"
	"github.com/gofiber/fiber"
)

type FiberMockHandler struct {
	mockService domain.MockService
}

func NewFiberMockHandler(app *fiber.App, ms domain.MockService) {
	handler := &FiberMockHandler{
		mockService: ms,
	}

	app.Use(handler.MockRequest)
}

// MockRequest handles a request and returns
// a mocked response for the active project/scenario
func (h *FiberMockHandler) MockRequest(ctx *fiber.Ctx) {
	request := domain.MockRequest{
		Project: ctx.Get("ProjectID"),
		Method:  ctx.Method(),
		Path:    ctx.Path(),
	}

	fmt.Println(request)

	scenario, err := h.mockService.FetchScenarioByRequest(context.Background(), request)

	if err != nil {
		ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	} else {
		var response fiber.Map
		_ = json.Unmarshal([]byte(scenario.Response), &response)
		ctx.JSON(response)
	}
}

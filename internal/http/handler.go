package http

import (
	"ApiService/internal/dto"
	"ApiService/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
)

type TaskHandler struct {
	service  service.TaskService
	log      *zap.SugaredLogger
	validate *validator.Validate
}

func NewTaskHandler(service service.TaskService, log *zap.SugaredLogger) *TaskHandler {
	return &TaskHandler{service: service, log: log, validate: validator.New()}
}

func (h *TaskHandler) CreateTask(ctx *fiber.Ctx) error {
	var req service.TaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		h.log.Error("Invalid request body", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}

	if err := h.validate.Struct(req); err != nil {
		h.log.Error("Validation error", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Validation failed: "+err.Error())
	}

	id, err := h.service.CreateTask(ctx.Context(), req)
	if err != nil {
		h.log.Error("Failed to create task", zap.Error(err))
		return dto.InternalServerError(ctx)
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Status: "success",
		Data:   map[string]int{"task_id": id},
	})
}

func (h *TaskHandler) GetTaskById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		h.log.Error("Get Task Error", zap.Error(err))
		return dto.NotFoundError(ctx, dto.FieldBadFormat, "Bad id")
	}
	task, err := h.service.GetTaskById(ctx.Context(), id)
	if err != nil {
		h.log.Error("Task not found", zap.Error(err))
		return dto.NotFoundError(ctx, dto.FieldBadFormat, "Task not found")
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Status: "success",
		Data:   task,
	})
}

func (h *TaskHandler) UpdateTask(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		h.log.Error("Invalid task id", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid task id")
	}
	var req service.TaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		h.log.Error("Invalid request body", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}

	updatedId, err := h.service.UpdateTask(ctx.Context(), id, req)
	if err != nil {
		h.log.Error("Failed to update task", zap.Error(err))
		return dto.NotFoundError(ctx, dto.NotFound, "Task not found")
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Status: "success",
		Data:   map[string]int{"task_id": updatedId},
	})
}

func (h *TaskHandler) DeleteTask(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		h.log.Error("Delete Task Error", zap.Error(err))
		return dto.NotFoundError(ctx, dto.FieldBadFormat, "Bad id")
	}
	taskId, err := h.service.DeleteTask(ctx.Context(), id)
	if err != nil {
		h.log.Error("Delete Task Error", zap.Error(err))
		return dto.NotFoundError(ctx, dto.NotFound, "Task not found")
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Status: "success",
		Data:   map[string]int{"task_id": taskId},
	})
}

package service

import (
	"ApiService/internal/repo"
	"context"
	"errors"
	"go.uber.org/zap"
)

type TaskService interface {
	CreateTask(ctx context.Context, req TaskRequest) (int, error)
	GetTaskById(ctx context.Context, id int) (repo.Task, error)
	UpdateTask(ctx context.Context, id int, req TaskRequest) (int, error)
	DeleteTask(ctx context.Context, id int) (int, error)
}

type taskService struct {
	repo repo.Repository
	log  *zap.SugaredLogger
}

func NewService(repo repo.Repository, logger *zap.SugaredLogger) TaskService {
	return &taskService{
		repo: repo,
		log:  logger,
	}
}

func (s *taskService) CreateTask(ctx context.Context, req TaskRequest) (int, error) {
	task := repo.Task{
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
	}
	taskID, err := s.repo.CreateTask(ctx, task)
	if err != nil {
		s.log.Error("Create Task Error", zap.Error(err))
		return 0, errors.New("create Task Error")
	}
	s.log.Infof("Task created task ID: %d", taskID)
	return taskID, nil
}

func (s *taskService) GetTaskById(ctx context.Context, id int) (repo.Task, error) {
	task, err := s.repo.GetTaskById(ctx, id)
	if err != nil {
		s.log.Error("Get Task Error", zap.Error(err))
		return repo.Task{}, errors.New("task not found")
	}

	return task, nil
}

func (s *taskService) UpdateTask(ctx context.Context, id int, req TaskRequest) (int, error) {
	task := repo.Task{
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	id, err := s.repo.UpdateTask(ctx, id, task)
	if err != nil {
		s.log.Error("Update Task Error", zap.Error(err))
		return 0, errors.New("task not found")
	}
	return id, nil
}

func (s *taskService) DeleteTask(ctx context.Context, id int) (int, error) {
	taskID, err := s.repo.DeleteTask(ctx, id)
	if err != nil {
		s.log.Error("Delete Task Error", zap.Error(err))
		return 0, errors.New("delete Task Error")
	}
	return taskID, nil
}

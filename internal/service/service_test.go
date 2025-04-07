package service

import (
	"ApiService/internal/repo"
	"ApiService/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
	"testing"
	"time"
)

func TestCreateTask(t *testing.T) {
	logger := zaptest.NewLogger(t).Sugar()

	t.Run("success", func(t *testing.T) {
		mockRepo := new(mocks.Repository)
		service := NewService(mockRepo, logger)

		req := TaskRequest{
			UserId:      1,
			Title:       "Test Task",
			Description: "Test Description",
		}

		expectedID := 123
		mockRepo.On("CreateTask", mock.Anything, mock.AnythingOfType("repo.Task")).
			Return(expectedID, nil)

		taskID, err := service.CreateTask(context.Background(), req)

		assert.NoError(t, err)
		assert.Equal(t, expectedID, taskID)
		mockRepo.AssertExpectations(t)
	})
	t.Run("repo error", func(t *testing.T) {
		mockRepo := new(mocks.Repository)
		service := NewService(mockRepo, logger)

		req := TaskRequest{
			UserId:      1,
			Title:       "Test Task",
			Description: "Test Description",
		}

		mockRepo.On("CreateTask", mock.Anything, mock.AnythingOfType("repo.Task")).
			Return(0, errors.New("unable to create task"))

		taskID, err := service.CreateTask(context.Background(), req)

		assert.Error(t, err)
		assert.Equal(t, 0, taskID)
		assert.EqualError(t, err, "create Task Error")
		mockRepo.AssertExpectations(t)
	})
}

func TestGetTaskById(t *testing.T) {
	logger := zaptest.NewLogger(t).Sugar()

	t.Run("success", func(t *testing.T) {
		mockRepo := new(mocks.Repository)
		service := NewService(mockRepo, logger)

		expectedTask := repo.Task{
			Id:          1,
			UserId:      2,
			Title:       "Test Task",
			Description: "Description",
			Status:      "new",
			CreateAt:    time.Now(),
		}

		mockRepo.On("GetTaskById", mock.Anything, expectedTask.Id).
			Return(expectedTask, nil)

		task, err := service.GetTaskById(context.Background(), expectedTask.Id)

		assert.NoError(t, err)
		assert.Equal(t, expectedTask, task)
		mockRepo.AssertExpectations(t)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo := new(mocks.Repository)
		service := NewService(mockRepo, logger)

		mockRepo.On("GetTaskById", mock.Anything, 999).
			Return(repo.Task{}, errors.New("task not found"))

		task, err := service.GetTaskById(context.Background(), 999)

		assert.Error(t, err)
		assert.EqualError(t, err, "task not found")
		assert.Equal(t, repo.Task{}, task)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateTask(t *testing.T) {
	logger := zaptest.NewLogger(t).Sugar()

	t.Run("success", func(t *testing.T) {
		mockRepo := new(mocks.Repository)
		service := NewService(mockRepo, logger)

		req := TaskRequest{
			UserId:      1,
			Title:       "Updated Task",
			Description: "Updated Description",
			Status:      "in_progress",
		}

		taskID := 42

		expectedTask := repo.Task{
			UserId:      req.UserId,
			Title:       req.Title,
			Description: req.Description,
			Status:      req.Status,
		}

		mockRepo.On("UpdateTask", mock.Anything, taskID, expectedTask).
			Return(taskID, nil)

		resultID, err := service.UpdateTask(context.Background(), taskID, req)

		assert.NoError(t, err)
		assert.Equal(t, taskID, resultID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo := new(mocks.Repository)
		service := NewService(mockRepo, logger)

		req := TaskRequest{
			UserId:      2,
			Title:       "Fail Task",
			Description: "Does not exist",
			Status:      "pending",
		}

		taskID := 999
		expectedTask := repo.Task{
			UserId:      req.UserId,
			Title:       req.Title,
			Description: req.Description,
			Status:      req.Status,
		}

		mockRepo.On("UpdateTask", mock.Anything, taskID, expectedTask).
			Return(0, errors.New("task not found"))

		resultID, err := service.UpdateTask(context.Background(), taskID, req)

		assert.Error(t, err)
		assert.EqualError(t, err, "task not found")
		assert.Equal(t, 0, resultID)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteTask(t *testing.T) {
	logger := zaptest.NewLogger(t).Sugar()
	t.Run("success", func(t *testing.T) {
		mockRepo := new(mocks.Repository)
		service := NewService(mockRepo, logger)
		taskID := 42

		mockRepo.On("DeleteTask", mock.Anything, taskID).Return(taskID, nil)

		id, err := service.DeleteTask(context.Background(), taskID)
		assert.NoError(t, err)
		assert.Equal(t, taskID, id)
		mockRepo.AssertExpectations(t)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo := new(mocks.Repository)
		taskID := 99
		service := NewService(mockRepo, logger)

		mockRepo.On("DeleteTask", mock.Anything, taskID).Return(0, errors.New("unable to delete task"))
		deletedID, err := service.DeleteTask(context.Background(), taskID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "delete Task Error")
		assert.Equal(t, 0, deletedID)
		mockRepo.AssertExpectations(t)
	})
}

package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	dtoreq "github.com/handiism/tasks-api-clone/dto/request"
	"github.com/handiism/tasks-api-clone/model"
	"github.com/handiism/tasks-api-clone/repo"
)

type SubtaskService interface {
	Add(userId uuid.UUID, listId uint, taskId uint, inp dtoreq.SaveSubtask) (model.Subtask, error)
	Update(userId uuid.UUID, listId uint, taskId uint, subtaskId uint, inp dtoreq.SaveSubtask) (model.Subtask, error)
	Fetch(userId uuid.UUID, listId uint, taskId uint, subtaskId uint) (model.Subtask, error)
	FetchAll(userId uuid.UUID, listId uint, taskId uint) ([]model.Subtask, error)
	Delete(id uuid.UUID, listId uint, taskId uint, subtaskId uint) error
}

type subtaskService struct {
	userRepo    repo.UserRepo
	listRepo    repo.ListRepo
	taskRepo    repo.TaskRepo
	subtaskRepo repo.SubtaskRepo
}

func (s *subtaskService) Add(userId uuid.UUID, listId uint, taskId uint, inp dtoreq.SaveSubtask) (model.Subtask, error) {
	if _, err := s.userRepo.FindByUUID(userId); err != nil {
		return model.Subtask{}, err
	}

	if list, err := s.listRepo.Find(listId); err != nil && list.UserID != userId {
		return model.Subtask{}, errors.New("list not found in targeted user")
	}

	if task, err := s.taskRepo.Find(taskId); err != nil && task.ListID != listId {
		return model.Subtask{}, errors.New("task not found in targeted list")
	} else {
		subtask, err := s.subtaskRepo.Store(model.Subtask{
			TaskID: taskId,
			Name:   inp.Name,
			IsDone: inp.IsDone,
		})
		if err != nil {
			return model.Subtask{}, err
		} else {
			return subtask, nil
		}
	}
}

func (s *subtaskService) Delete(id uuid.UUID, listId uint, taskId uint, subtaskId uint) error {
	if _, err := s.userRepo.FindByUUID(id); err != nil {
		return err
	}

	if list, err := s.listRepo.Find(listId); err != nil && list.UserID != id {
		return errors.New("list not found in targeted user")
	}

	if task, err := s.taskRepo.Find(taskId); err != nil && task.ListID != listId {
		return errors.New("task not found in targeted list")
	}

	if subtask, err := s.subtaskRepo.Find(subtaskId); err != nil {
		return err
	} else {
		if err := s.subtaskRepo.Delete(subtask); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func (s *subtaskService) Fetch(userId uuid.UUID, listId uint, taskId uint, subtaskId uint) (model.Subtask, error) {
	if _, err := s.userRepo.FindByUUID(userId); err != nil {
		return model.Subtask{}, err
	}

	if list, err := s.listRepo.Find(listId); err != nil && list.UserID != userId {
		return model.Subtask{}, errors.New("list not found in targeted user")
	}

	if task, err := s.taskRepo.Find(taskId); err != nil && task.ListID != listId {
		return model.Subtask{}, errors.New("task not found in targeted list")
	}

	if subtask, err := s.subtaskRepo.Find(subtaskId); err != nil && subtask.TaskID != taskId {
		return model.Subtask{}, errors.New("subtask not found in targeted task")
	} else {
		return subtask, nil
	}
}

func (s *subtaskService) FetchAll(userId uuid.UUID, listId uint, taskId uint) ([]model.Subtask, error) {
	if _, err := s.userRepo.FindByUUID(userId); err != nil {
		return nil, err
	}

	if list, err := s.listRepo.Find(listId); err != nil && list.UserID != userId {
		return nil, errors.New("list not found in targeted user")
	}

	if task, err := s.taskRepo.Find(taskId); err != nil && task.ListID != listId {
		return nil, errors.New("task not found in targeted list")
	} else {
		if task, err = s.taskRepo.Preload(task); err != nil {
			return nil, err
		} else {
			return task.Subtasks, nil
		}
	}
}

func (s *subtaskService) Update(userId uuid.UUID, listId uint, taskId uint, subtaskId uint, inp dtoreq.SaveSubtask) (model.Subtask, error) {
	if _, err := s.userRepo.FindByUUID(userId); err != nil {
		return model.Subtask{}, err
	}

	if list, err := s.listRepo.Find(listId); err != nil && list.UserID != userId {
		return model.Subtask{}, errors.New("list not found in targeted user")
	}

	if task, err := s.taskRepo.Find(taskId); err != nil && task.ListID != listId {
		return model.Subtask{}, errors.New("task not found in targeted list")
	}

	if subtask, err := s.subtaskRepo.Find(subtaskId); err != nil && subtask.TaskID != taskId {
		return model.Subtask{}, errors.New("subtask not found in targeted task")
	} else {
		subtask, err = s.subtaskRepo.Update(model.Subtask{
			ID:        subtask.ID,
			TaskID:    taskId,
			Name:      inp.Name,
			IsDone:    inp.IsDone,
			CreatedAt: subtask.CreatedAt,
			UpdatedAt: time.Now(),
		})
		if err != nil {
			return model.Subtask{}, err
		} else {
			return subtask, nil
		}
	}
}

func NewSubtaskService(userRepo repo.UserRepo, listRepo repo.ListRepo, taskRepo repo.TaskRepo, subtaskRepo repo.SubtaskRepo) SubtaskService {
	return &subtaskService{
		userRepo:    userRepo,
		listRepo:    listRepo,
		taskRepo:    taskRepo,
		subtaskRepo: subtaskRepo,
	}
}

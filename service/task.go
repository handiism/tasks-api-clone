package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	dtoreq "github.com/handiism/tasks-api-clone/dto/request"
	"github.com/handiism/tasks-api-clone/model"
	"github.com/handiism/tasks-api-clone/repo"
)

type TaskService interface {
	Add(userId uuid.UUID, listId uint, inp dtoreq.SaveTask) (model.Task, error)
	Update(userId uuid.UUID, listId uint, taskId uint, inp dtoreq.SaveTask) (model.Task, error)
	Fetch(userId uuid.UUID, listId uint, taskId uint) (model.Task, error)
	FetchAll(userId uuid.UUID, listId uint) ([]model.Task, error)
	Delete(id uuid.UUID, listId uint, taskId uint) error
}

type taskService struct {
	userRepo repo.UserRepo
	listRepo repo.ListRepo
	taskRepo repo.TaskRepo
}

func (t *taskService) Add(userId uuid.UUID, listId uint, inp dtoreq.SaveTask) (model.Task, error) {
	user, err := t.userRepo.FindByUUID(userId)
	if err != nil {
		return model.Task{}, err
	}

	user, _ = t.userRepo.Preload(user)
	found := false
	for _, v := range user.Lists {
		if v.ID == listId {
			found = true
		}
	}

	if !found {
		return model.Task{}, err
	}

	task, err := t.taskRepo.Store(model.Task{
		ListID:  listId,
		Name:    inp.Name,
		Detail:  inp.Detail,
		DueDate: inp.DueDate,
		IsDone:  false,
	})

	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (t *taskService) Delete(id uuid.UUID, listId uint, taskId uint) error {
	if _, err := t.userRepo.FindByUUID(id); err != nil {
		return err
	}

	if list, err := t.listRepo.Find(listId); err != nil && list.UserID != id {
		return errors.New("list not found in targeted user")
	}

	if task, err := t.taskRepo.Find(taskId); err != nil && task.ListID != listId {
		return errors.New("task not found in targeted list")
	} else {
		if err = t.taskRepo.Delete(task); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func (t *taskService) Fetch(userId uuid.UUID, listId uint, taskId uint) (model.Task, error) {
	if _, err := t.userRepo.FindByUUID(userId); err != nil {
		return model.Task{}, err
	}

	if list, err := t.listRepo.Find(listId); err != nil && list.UserID != userId {
		return model.Task{}, err
	}

	if task, err := t.taskRepo.Find(taskId); err != nil && task.ListID != listId {
		return model.Task{}, errors.New("task not found in targeted list")
	} else {
		return task, err
	}
}

func (t *taskService) FetchAll(userId uuid.UUID, listId uint) ([]model.Task, error) {
	if _, err := t.userRepo.FindByUUID(userId); err != nil {
		return nil, err
	}

	if list, err := t.listRepo.Find(listId); err != nil && list.UserID != userId {
		return nil, errors.New("list not found in targeted user")
	} else {
		return list.Tasks, nil
	}
}

func (t *taskService) Update(userId uuid.UUID, listId uint, taskId uint, inp dtoreq.SaveTask) (model.Task, error) {
	if _, err := t.userRepo.FindByUUID(userId); err != nil {
		return model.Task{}, err
	}

	if list, err := t.listRepo.Find(listId); err != nil && list.UserID != userId {
		return model.Task{}, errors.New("list not found in targeted user")
	}

	if task, err := t.taskRepo.Find(taskId); err != nil && task.ListID != listId {
		return model.Task{}, errors.New("task not found in targeted list")
	} else {
		task, err := t.taskRepo.Update(model.Task{
			ID:        task.ID,
			ListID:    listId,
			Name:      inp.Name,
			Detail:    inp.Detail,
			DueDate:   inp.DueDate,
			IsDone:    inp.IsDone,
			CreatedAt: task.CreatedAt,
			UpdatedAt: time.Now(),
		})
		if err != nil {
			return model.Task{}, err
		} else {
			return task, nil
		}
	}
}

func NewTaskService(userRepo repo.UserRepo, listRepo repo.ListRepo, taskRepo repo.TaskRepo) TaskService {
	return &taskService{
		userRepo: userRepo,
		listRepo: listRepo,
		taskRepo: taskRepo,
	}
}

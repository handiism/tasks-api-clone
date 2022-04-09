package service

import (
	"errors"

	"github.com/google/uuid"
	dtoreq "github.com/handiism/tasks-api-clone/dto/request"
	"github.com/handiism/tasks-api-clone/model"
	"github.com/handiism/tasks-api-clone/repo"
)

type ListService interface {
	Add(id uuid.UUID, inp dtoreq.SaveList) (model.List, error)
	Update(userId uuid.UUID, listId uint, inp dtoreq.SaveList) (model.List, error)
	Fetch(userId uuid.UUID, listId uint) (model.List, error)
	FetchAll(id uuid.UUID) ([]model.List, error)
	Delete(id uuid.UUID, listId uint) error
}

type listService struct {
	userRepo repo.UserRepo
	listRepo repo.ListRepo
}

func NewListService(userRepo repo.UserRepo, listRepo repo.ListRepo) ListService {
	return &listService{
		userRepo: userRepo,
		listRepo: listRepo,
	}
}

func (l *listService) Add(id uuid.UUID, inp dtoreq.SaveList) (model.List, error) {
	list := model.List{
		UserID: id,
		Title:  inp.Title,
	}
	if res, err := l.listRepo.Store(list); err != nil {
		return model.List{}, err
	} else {
		return res, nil
	}
}

func (l *listService) Update(userId uuid.UUID, listId uint, inp dtoreq.SaveList) (model.List, error) {
	user, err := l.userRepo.FindByUUID(userId)
	if err != nil {
		return model.List{}, err
	}

	user, err = l.userRepo.Preload(user)
	if err != nil {
		return model.List{}, err
	}

	list, err := l.listRepo.Find(listId)
	if err != nil {
		return model.List{}, err
	}

	found := false
	for _, v := range user.Lists {
		if v.ID == list.ID {
			found = true
			break
		}
	}

	if !found {
		return model.List{}, err
	}

	list.Title = inp.Title
	res, err := l.listRepo.Update(list)
	if err != nil {
		return model.List{}, err
	}

	return res, nil
}

func (l *listService) Fetch(userId uuid.UUID, listId uint) (model.List, error) {
	list, err := l.listRepo.Find(listId)
	if err != nil {
		return model.List{}, err
	}

	user, err := l.userRepo.FindByUUID(userId)
	if err != nil {
		return model.List{}, err
	}

	user, err = l.userRepo.PreloadAll(user)
	if err != nil {
		return model.List{}, err
	}

	valid := false
	for _, li := range user.Lists {
		if li.ID == list.ID {
			valid = true
			break
		}
	}
	if valid {
		return list, nil
	} else {
		return model.List{}, errors.New("list not found")
	}
}

func (l *listService) FetchAll(id uuid.UUID) ([]model.List, error) {
	user, err := l.userRepo.FindByUUID(id)
	if err != nil {
		return nil, err
	}
	userPreload, err := l.userRepo.Preload(user)
	return userPreload.Lists, nil
}

func (l *listService) Delete(id uuid.UUID, listId uint) error {
	user, err := l.userRepo.FindByUUID(id)
	if err != nil {
		return err
	}

	user, err = l.userRepo.Preload(user)
	if err != nil {
		return err
	}

	list, err := l.listRepo.Find(listId)
	if err != nil {
		return err
	}

	found := false
	for _, v := range user.Lists {
		if v.ID == list.ID {
			found = true
		}
	}

	if found {
		if err := l.listRepo.Delete(list); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return errors.New("list not registered with the user")
	}
}

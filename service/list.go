package service

import (
	"github.com/google/uuid"
	dtoreq "github.com/handiism/tasks-api-clone/dto/request"
	"github.com/handiism/tasks-api-clone/model"
	"github.com/handiism/tasks-api-clone/repo"
)

type ListService interface {
	Add(id uuid.UUID, inp dtoreq.SaveList) (model.List, error)
	Update(id uuid.UUID, inp dtoreq.SaveUser) (model.List, error)
	FetchAll(id uuid.UUID) ([]model.List, error)
	Delete(id uuid.UUID) error
}

type listService struct {
	listRepo repo.ListRepo
}

func NewListService(listRepo repo.ListRepo) ListService {
	return &listService{
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

func (l *listService) Update(id uuid.UUID, inp dtoreq.SaveUser) (model.List, error) {
	panic("not implemented") // TODO: Implement
}

func (l *listService) FetchAll(id uuid.UUID) ([]model.List, error) {
	panic("not implemented") // TODO: Implement
}

func (l *listService) Delete(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

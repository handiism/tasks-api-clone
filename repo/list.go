package repo

import (
	"github.com/handirachmawan/tasks-api-clone/model"

	"gorm.io/gorm"
)

type ListRepo interface {
	Store(model.List) (model.List, error)
	Update(model.List) (model.List, error)
	Delete(model.List) error
	Find(id uint) (model.List, error)
	Preload(list model.List) (model.List, error)
	PreloadAll(list model.List) (model.List, error)
}

type listRepo struct {
	db *gorm.DB
}

func NewListRepo(db *gorm.DB) ListRepo {
	return &listRepo{db}
}

func (l *listRepo) Store(list model.List) (model.List, error) {
	if err := l.db.Create(&list).Error; err != nil {
		return model.List{}, err
	}

	return list, nil
}

func (l *listRepo) Find(id uint) (model.List, error) {
	var list model.List

	if err := l.db.Where(model.List{ID: id}).First(&list).Error; err != nil {
		return model.List{}, err
	}

	return list, nil
}

func (l *listRepo) Update(list model.List) (model.List, error) {

	if err := l.db.Save(&list).Error; err != nil {
		return model.List{}, err
	}

	return list, nil
}

func (l *listRepo) Delete(list model.List) error {
	if err := l.db.Delete(&list).Error; err != nil {
		return err
	}

	return nil
}

func (l *listRepo) Preload(list model.List) (model.List, error) {
	if err := l.db.Preload("Tasks").
		First(&list, list.ID).Error; err != nil {
		return model.List{}, err
	}

	return list, nil
}

func (l *listRepo) PreloadAll(list model.List) (model.List, error) {
	if err := l.db.Preload("Tasks.Subtasks").
		First(&list, list.ID).Error; err != nil {
		return model.List{}, err
	}

	return list, nil
}

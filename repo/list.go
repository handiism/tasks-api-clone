package repo

import (
	"tasks/model"

	"gorm.io/gorm"
)

type List interface {
	Create(model.List) (model.List, error)
	Update(model.List) (model.List, error)
	Delete(model.List) error
	Find(id uint) (model.List, error)
}

type list struct {
	db *gorm.DB
}

func NewList(db *gorm.DB) List {
	return &list{db}
}

func (l *list) Create(list model.List) (model.List, error) {
	if err := l.db.Create(&list).Error; err != nil {
		return model.List{}, err
	}

	return list, nil
}

func (l *list) Find(id uint) (model.List, error) {
	var list model.List

	if err := l.db.Where(model.List{ID: id}).First(list).Error; err != nil {
		return model.List{}, err
	}

	return list, nil
}

func (l *list) Update(list model.List) (model.List, error) {
	if err := l.db.Save(&list).Error; err != nil {
		return model.List{}, err
	}

	return list, nil
}

func (l *list) Delete(list model.List) error {
	if err := l.db.Delete(list).Error; err != nil {
		return err
	}

	return nil
}

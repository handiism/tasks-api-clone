package repo

import (
	"tasks/model"

	"gorm.io/gorm"
)

type Subtask interface {
	Create(model.Subtask) (model.Subtask, error)
	Find(id uint) (model.Subtask, error)
	Update(model.Subtask) (model.Subtask, error)
	Delete(model.Subtask) error
}

type subtask struct {
	db *gorm.DB
}

func NewSubtask(db *gorm.DB) Subtask {
	return &subtask{db}
}

func (s *subtask) Create(subtask model.Subtask) (model.Subtask, error) {
	if err := s.db.Create(&subtask).Error; err != nil {
		return model.Subtask{}, err
	}

	return subtask, nil
}

func (s *subtask) Find(id uint) (model.Subtask, error) {
	var subtask model.Subtask

	if err := s.db.Where(model.Subtask{ID: id}).First(&subtask).Error; err != nil {
		return model.Subtask{}, err
	}

	return subtask, nil
}

func (s *subtask) Update(subtask model.Subtask) (model.Subtask, error) {
	if err := s.db.Save(&subtask).Error; err != nil {
		return model.Subtask{}, err
	}

	return subtask, nil
}

func (s *subtask) Delete(subtask model.Subtask) error {
	if err := s.db.Delete(&subtask).Error; err != nil {
		return err
	}

	return nil
}

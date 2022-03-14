package repo

import (
	"tasks/model"

	"gorm.io/gorm"
)

type SubtaskRepo interface {
	Create(model.Subtask) (model.Subtask, error)
	Find(id uint) (model.Subtask, error)
	Update(model.Subtask) (model.Subtask, error)
	Delete(model.Subtask) error
}

type subtaskRepo struct {
	db *gorm.DB
}

func NewSubtaskRepo(db *gorm.DB) SubtaskRepo {
	return &subtaskRepo{db}
}

func (s *subtaskRepo) Create(subtask model.Subtask) (model.Subtask, error) {
	if err := s.db.Create(&subtask).Error; err != nil {
		return model.Subtask{}, err
	}

	return subtask, nil
}

func (s *subtaskRepo) Find(id uint) (model.Subtask, error) {
	var subtask model.Subtask

	if err := s.db.Where(model.Subtask{ID: id}).First(&subtask).Error; err != nil {
		return model.Subtask{}, err
	}

	return subtask, nil
}

func (s *subtaskRepo) Update(subtask model.Subtask) (model.Subtask, error) {
	if err := s.db.Save(&subtask).Error; err != nil {
		return model.Subtask{}, err
	}

	return subtask, nil
}

func (s *subtaskRepo) Delete(subtask model.Subtask) error {
	if err := s.db.Delete(&subtask).Error; err != nil {
		return err
	}

	return nil
}

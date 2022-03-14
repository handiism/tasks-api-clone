package repo

import (
	"tasks/model"

	"gorm.io/gorm"
)

type TaskRepo interface {
	Create(model.Task) (model.Task, error)
	Find(id uint) (model.Task, error)
	Update(model.Task) (model.Task, error)
	Delete(model.Task) error
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepo {
	return &taskRepo{db}
}

func (t *taskRepo) Create(task model.Task) (model.Task, error) {
	if err := t.db.Create(&task).Error; err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (t *taskRepo) Find(id uint) (model.Task, error) {
	var task model.Task

	if err := t.db.Where(model.Task{ID: id}).First(&task).Error; err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (t *taskRepo) Update(task model.Task) (model.Task, error) {
	if err := t.db.Save(&task).Error; err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (t *taskRepo) Delete(task model.Task) error {
	if err := t.db.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}

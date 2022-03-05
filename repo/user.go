package repo

import (
	"tasks/model"

	"gorm.io/gorm"
)

type User interface {
	Create(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	IsEmailAvailable(email string) bool
	Update(user model.User) (model.User, error)
	Delete(user model.User) error
	Verify(email, password string) (model.User, error)
	Detail(user model.User) (model.User, error)
	DetailUsingEmail(email string) (model.User, error)
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return &user{
		db: db,
	}
}

func (u *user) Create(user model.User) (model.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return model.User{}, nil
	}

	return user, nil
}

func (u *user) FindByEmail(email string) (model.User, error) {
	var user model.User

	if err := u.db.Where(model.User{Email: email}).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *user) IsEmailAvailable(email string) bool {
	if err := u.db.Where(model.User{Email: email}).
		First(&model.User{}).Error; err != nil {
		return true
	}

	return false
}

func (u *user) Update(user model.User) (model.User, error) {
	if err := u.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *user) Delete(user model.User) error {
	if err := u.db.Delete(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *user) Verify(email, password string) (model.User, error) {
	var user model.User

	if err := u.db.Where(model.User{Email: email, Password: password}).
		First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *user) Detail(user model.User) (model.User, error) {
	if err := u.db.Preload("list").Preload("task").Preload("subtask").
		Where(model.User{ID: user.ID}).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *user) DetailUsingEmail(email string) (model.User, error) {
	var user model.User

	if err := u.db.Joins("list").Joins("task").Joins("subtask").
		Where(model.User{Email: email}).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

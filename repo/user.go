package repo

import (
	"github.com/google/uuid"
	"github.com/handirachmawan/tasks-api-clone/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	Store(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindByUUID(uuid uuid.UUID) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(user model.User) error
	Preload(user model.User) (model.User, error)
	PreloadAll(user model.User) (model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Store(user model.User) (model.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepo) FindByEmail(email string) (model.User, error) {
	var user model.User

	if err := u.db.Where(model.User{Email: email}).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepo) FindByUUID(id uuid.UUID) (model.User, error) {
	var user model.User

	if err := u.db.Where(model.User{ID: id}).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepo) Update(user model.User) (model.User, error) {
	if err := u.db.Save(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepo) Delete(user model.User) error {
	if err := u.db.Delete(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepo) Preload(user model.User) (model.User, error) {
	if err := u.db.Preload("Lists").
		First(&user, user.ID).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepo) PreloadAll(user model.User) (model.User, error) {
	if err := u.db.Preload("Lists.Tasks.Subtasks").
		First(&user, user.ID).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

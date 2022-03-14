package repo

import (
	"tasks/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	IsEmailAvailable(email string) bool
	Update(user model.User) (model.User, error)
	Delete(user model.User) error
	Verify(email, password string) (model.User, error)
	Detail(user model.User) (model.User, error)
	DetailUsingEmail(email string) (model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(user model.User) (model.User, error) {
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

func (u *userRepo) IsEmailAvailable(email string) bool {
	if err := u.db.Where(model.User{Email: email}).
		First(&model.User{}).Error; err != nil {
		return true
	}

	return false
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

func (u *userRepo) Verify(email, password string) (model.User, error) {
	var user model.User

	if err := u.db.Where(model.User{Email: email, Password: password}).
		First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepo) Detail(user model.User) (model.User, error) {
	if err := u.db.Preload("Lists.Tasks.Subtasks").
		First(&user, user.ID).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepo) DetailUsingEmail(email string) (model.User, error) {
	var user model.User

	if err := u.db.Where(model.User{Email: email}).Preload("Lists.Tasks.Subtasks").
		First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

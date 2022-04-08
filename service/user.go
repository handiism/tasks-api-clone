package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	dtoreq "github.com/handirachmawan/tasks-api-clone/dto/request"
	"github.com/handirachmawan/tasks-api-clone/model"
	"github.com/handirachmawan/tasks-api-clone/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(inp dtoreq.SaveUser) (model.User, error)
	Login(inp dtoreq.Login) (model.User, error)
	IsEmailAvailable(inp dtoreq.EmailChecking) bool
	Update(id uuid.UUID, inp dtoreq.SaveUser) (model.User, error)
	Fetch(id string) (model.User, error)
	Delete(id string) error
}

type userService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Register(inp dtoreq.SaveUser) (model.User, error) {
	avail := u.IsEmailAvailable(dtoreq.EmailChecking{Email: inp.Email})
	if !avail {
		return model.User{}, errors.New("cannot use registered email")
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(inp.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	now := time.Now()
	user := model.User{
		Name:      inp.Name,
		Email:     inp.Email,
		Password:  string(pw),
		CreatedAt: now,
		UpdatedAt: now,
	}

	user, err = u.userRepo.Store(user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userService) Login(inp dtoreq.Login) (model.User, error) {
	user, err := u.userRepo.FindByEmail(inp.Email)

	if err != nil {
		return model.User{}, errors.New("email not registered yet")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inp.Password)); err != nil {
		return model.User{}, err
	} else {

		return user, nil
	}
}

func (u *userService) IsEmailAvailable(inp dtoreq.EmailChecking) bool {
	if _, err := u.userRepo.FindByEmail(inp.Email); err != nil {
		return true
	} else {
		return false
	}
}

func (u *userService) Update(id uuid.UUID, inp dtoreq.SaveUser) (model.User, error) {
	user, err := u.userRepo.FindByUUID(id)

	if err != nil {
		return model.User{}, err
	}

	if _, err := u.userRepo.FindByEmail(inp.Email); err == nil {
		return model.User{}, err
	}

	bytePass, err := bcrypt.GenerateFromPassword([]byte(inp.Password), bcrypt.DefaultCost)

	if err != nil {
		return model.User{}, err
	}

	user.Name = inp.Name
	user.Email = inp.Email
	user.Password = string(bytePass)

	if user, err := u.userRepo.Update(user); err != nil {
		return model.User{}, err
	} else {
		return user, nil
	}
}

func (u *userService) Fetch(id string) (model.User, error) {
	uuid, err := uuid.Parse(id)

	if err != nil {
		return model.User{}, err
	}

	user, err := u.userRepo.FindByUUID(uuid)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userService) Delete(id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	user, err := u.userRepo.FindByUUID(uuid)
	if err != nil {
		return err
	}

	err = u.userRepo.Delete(user)
	if err != nil {
		return err
	}

	return nil
}

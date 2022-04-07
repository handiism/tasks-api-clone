package repo_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/handirachmawan/tasks-api-clone/model"
	"github.com/handirachmawan/tasks-api-clone/repo"
	"github.com/stretchr/testify/assert"
)

var (
	existingUser    model.User
	existingUUID, _        = uuid.Parse("8496679d-ce13-4074-b4af-f8cc37e1cbd8")
	existingEmail   string = "nxCjeWO@gItZUIK.com"
)

func TestUserCreateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewUserRepo(db)

	user, err := repo.Store(model.User{
		Name:     "Muhammad Handi Rachmawan",
		Email:    "email2@handiism.com",
		Password: "p4ssw*rd",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, user)

	existingUser = user
}

func TestUserCreateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewUserRepo(db)

	user, err := repo.Store(model.User{
		ID:        existingUUID,
		Name:      "",
		Email:     existingEmail,
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Token:     "",
		Lists:     []model.List{},
	})

	assert.NotNil(t, err)
	assert.Empty(t, user)
}

func TestUserFindByEmailSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewUserRepo(db)

	user, err := repo.FindByEmail(existingUser.Email)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func TestUserFindByEmailFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).FindByEmail(
		existingUser.Email + "adder",
	)

	assert.NotNil(t, err)
	assert.Empty(t, user)
}

func TestUserUpdateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	existingUser.Name = "change"
	existingUser.Email = "change"
	existingUser.Password = "change"

	user, err := repo.NewUserRepo(db).Update(existingUser)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)

	existingUser = user
}

func TestUserUpdateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).Update(model.User{
		ID:        [16]byte{},
		Name:      "",
		Email:     existingEmail,
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Token:     "",
		Lists:     []model.List{},
	})

	assert.NotNil(t, err)
	assert.Empty(t, user)
}

func TestUserPreloadSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).Preload(model.User{
		ID: existingUUID,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func TestUserPreloadFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).Preload(model.User{
		ID:        [16]byte{},
		Name:      "",
		Email:     "",
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Token:     "",
		Lists:     []model.List{},
	})

	assert.NotNil(t, err)
	assert.Empty(t, user)
}

func TestUserDeleteSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewUserRepo(db).Delete(existingUser)

	assert.Nil(t, err)
}

func TestUserDeleteFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewUserRepo(db).Delete(model.User{
		ID:        [16]byte{},
		Name:      "",
		Email:     "",
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Token:     "",
		Lists:     []model.List{},
	})

	assert.NotNil(t, err)
}

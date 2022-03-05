package test

import (
	"database/sql"
	"fmt"
	"tasks/model"
	"tasks/repo"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var realUser model.User

func TestRepoUserCreate(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewUser(db)

	user, err := repo.Create(model.User{
		ID:       uuid.NewString(),
		Name:     "Muhammad Handi Rachmawan",
		Email:    "email2@handiism.com",
		Password: "p4ssw*rd",
	})

	assert.NotNil(t, err)
	assert.NotEmpty(t, user)

	realUser = user
}

func TestRepoUserFindByEmailSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewUser(db)

	user, err := repo.FindByEmail(realUser.Email)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func TestRepoUserFindByEmailFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUser(db).FindByEmail(
		realUser.Email + "false",
	)

	assert.NotNil(t, err)
	assert.Empty(t, user)
}

func TestRepoUserIsEmailAvailabe(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	ok := repo.NewUser(db).IsEmailAvailable(
		realUser.Email + "false",
	)

	assert.True(t, ok)
}

func TestRepoUserIsEmailNotAvailabe(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	ok := repo.NewUser(db).IsEmailAvailable(realUser.Email)

	assert.False(t, ok)
}

func TestRepoUserUpdate(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	realUser, err := repo.NewUser(db).Update(model.User{
		ID:       realUser.ID,
		Name:     "GANTI",
		Email:    "GANTI",
		Password: "GANTI",
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, realUser)
}

func TestRepoUserVerifySuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUser(db).Verify(realUser.Email, realUser.Password)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func TestRepoUserVerifyFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUser(db).Verify(
		realUser.Email,
		realUser.Password+"false",
	)

	assert.NotNil(t, err)
	assert.Empty(t, user)
}

func TestRepoUserDetail(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	db.Session(&gorm.Session{
		Logger: logger.Default.LogMode(logger.Info),
	})

	user, err := repo.NewUser(db).Detail(realUser)

	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

func TestRepoUserDelete(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)
}

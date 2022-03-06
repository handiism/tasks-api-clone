package test

import (
	"database/sql"
	"tasks/model"
	"tasks/repo"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var existingUser model.User
var existingUUID string = "919e834b-1e72-4a59-a429-9e1aaf56245e"
var existingEmail string = "XutfjYv@XHGcFbH.org"

func TestRepoUserCreateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewUser(db)

	user, err := repo.Create(model.User{
		Name:     "Muhammad Handi Rachmawan",
		Email:    "email2@handiism.com",
		Password: "p4ssw*rd",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, user)

	existingUser = user
}

func TestRepoUserCreateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewUser(db)

	user, err := repo.Create(model.User{
		ID:        [16]byte{},
		Name:      "",
		Email:     existingEmail,
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: sql.NullTime{},
		Token:     "",
	})

	assert.NotNil(t, err)
	assert.Empty(t, user)
}

func TestRepoUserFindByEmailSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewUser(db)

	user, err := repo.FindByEmail(existingUser.Email)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func TestRepoUserFindByEmailFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUser(db).FindByEmail(
		existingUser.Email + "false",
	)

	assert.NotNil(t, err)
	assert.Empty(t, user)
}

func TestRepoUserIsEmailAvailabe(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	ok := repo.NewUser(db).IsEmailAvailable(
		existingUser.Email + "false",
	)

	assert.True(t, ok)
}

func TestRepoUserIsEmailNotAvailabe(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	ok := repo.NewUser(db).IsEmailAvailable(existingUser.Email)

	assert.False(t, ok)
}
func TestRepoUserUpdateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUser(db).Update(model.User{
		ID:        existingUser.ID,
		Name:      "GANTI" + existingUser.Name,
		Email:     "GANTI" + existingUser.Email,
		Password:  "GANTI" + existingUser.Password,
		CreatedAt: existingUser.CreatedAt,
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, user)

	existingUser = user
}

func TestRepoUserUpdateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUser(db).Update(model.User{
		ID:        [16]byte{},
		Name:      "",
		Email:     existingEmail,
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: sql.NullTime{},
		Token:     "",
		Lists:     []model.List{},
	})

	assert.NotNil(t, err)
	assert.Empty(t, user)
}

func TestRepoUserVerifySuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUser(db).Verify(
		existingUser.Email,
		existingUser.Password,
	)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func TestRepoUserVerifyFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUser(db).Verify(
		existingUser.Email,
		existingUser.Password+"false",
	)

	assert.NotNil(t, err)
	assert.Empty(t, user)
}

func TestRepoUserDetailSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	uuid, err := uuid.Parse(existingUUID)

	assert.Nil(t, err)

	user, err := repo.NewUser(db).Detail(model.User{
		ID: uuid,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func TestRepoUserDetailFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUser(db).Detail(model.User{
		ID:        [16]byte{},
		Name:      "",
		Email:     "",
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: sql.NullTime{},
		Token:     "",
		Lists:     []model.List{},
	})

	assert.NotNil(t, err)
	assert.Empty(t, user)
}

func TestRepoUserDetailUsingEmailSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUser(db).DetailUsingEmail(existingEmail)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func TestRepoUserDetailUsingEmailFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUser(db).DetailUsingEmail(existingEmail + "false")

	assert.NotNil(t, err)
	assert.Empty(t, user)
}

func TestRepoUserDeleteSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewUser(db).Delete(existingUser)

	assert.Nil(t, err)
}

func TestRepoUserDeleteFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewUser(db).Delete(model.User{
		ID:        [16]byte{},
		Name:      "",
		Email:     "",
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: sql.NullTime{},
		Token:     "",
		Lists:     []model.List{},
	})

	assert.NotNil(t, err)
}

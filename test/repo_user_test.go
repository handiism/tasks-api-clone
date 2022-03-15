package test

import (
	"github.com/handirachmawan/tasks-api-clone/helper"
	"github.com/handirachmawan/tasks-api-clone/model"
	"github.com/handirachmawan/tasks-api-clone/repo"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var (
	existingUser  model.User
	existingUUID  uuid.UUID = helper.ExistingUserUUID()
	existingEmail string    = helper.ExistingUserEmail()
)

func TestRepoUserCreateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewUserRepo(db)

	user, err := repo.Create(model.User{
		Name:     "Muhammad Handi Rachmawan",
		Email:    "email2@handiism.com",
		Password: "p4ssw*rd",
	})

	require.Nil(t, err)
	require.NotEmpty(t, user)

	existingUser = user
}

func TestRepoUserCreateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewUserRepo(db)

	user, err := repo.Create(model.User{
		ID:        existingUUID,
		Name:      "",
		Email:     existingEmail,
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Token:     "",
		Lists:     []model.List{},
	})

	require.NotNil(t, err)
	require.Empty(t, user)
}

func TestRepoUserFindByEmailSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewUserRepo(db)

	user, err := repo.FindByEmail(existingUser.Email)

	require.Nil(t, err)
	require.NotEmpty(t, user)
}

func TestRepoUserFindByEmailFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).FindByEmail(
		existingUser.Email + "adder",
	)

	require.NotNil(t, err)
	require.Empty(t, user)
}

func TestRepoUserIsEmailAvailabe(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	ok := repo.NewUserRepo(db).IsEmailAvailable(
		existingUser.Email + "adder",
	)

	require.True(t, ok)
}

func TestRepoUserIsEmailNotAvailabe(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	ok := repo.NewUserRepo(db).IsEmailAvailable(existingUser.Email)

	require.False(t, ok)
}
func TestRepoUserUpdateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	existingUser.Name = "change"
	existingUser.Email = "change"
	existingUser.Password = "change"

	user, err := repo.NewUserRepo(db).Update(existingUser)

	require.Nil(t, err)
	require.NotEmpty(t, user)

	existingUser = user
}

func TestRepoUserUpdateFailed(t *testing.T) {
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

	require.NotNil(t, err)
	require.Empty(t, user)
}

func TestRepoUserVerifySuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).Verify(
		existingUser.Email,
		existingUser.Password,
	)

	require.Nil(t, err)
	require.NotEmpty(t, user)
}

func TestRepoUserVerifyFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).Verify(
		existingUser.Email,
		existingUser.Password+"adder",
	)

	require.NotNil(t, err)
	require.Empty(t, user)
}

func TestRepoUserDetailSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).Detail(model.User{
		ID: existingUUID,
	})

	require.Nil(t, err)
	require.NotEmpty(t, user)
}

func TestRepoUserDetailFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).Detail(model.User{
		ID:        [16]byte{},
		Name:      "",
		Email:     "",
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Token:     "",
		Lists:     []model.List{},
	})

	require.NotNil(t, err)
	require.Empty(t, user)
}

func TestRepoUserDetailUsingEmailSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).DetailUsingEmail(existingEmail)

	require.Nil(t, err)
	require.NotEmpty(t, user)
}

func TestRepoUserDetailUsingEmailFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).DetailUsingEmail(existingEmail + "adder")

	require.NotNil(t, err)
	require.Empty(t, user)
}

func TestRepoUserDeleteSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewUserRepo(db).Delete(existingUser)

	require.Nil(t, err)
}

func TestRepoUserDeleteFailed(t *testing.T) {
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

	require.NotNil(t, err)
}

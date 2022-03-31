package repo_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/handirachmawan/tasks-api-clone/model"
	"github.com/handirachmawan/tasks-api-clone/repo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

var (
	existingUser    model.User
	existingUUID, _        = uuid.Parse(viper.GetString("USER_UUID"))
	existingEmail   string = viper.GetString("USER_EMAIL")
)

func TestUserCreateSuccess(t *testing.T) {
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

func TestUserCreateFailed(t *testing.T) {
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

func TestUserFindByEmailSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewUserRepo(db)

	user, err := repo.FindByEmail(existingUser.Email)

	require.Nil(t, err)
	require.NotEmpty(t, user)
}

func TestUserFindByEmailFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).FindByEmail(
		existingUser.Email + "adder",
	)

	require.NotNil(t, err)
	require.Empty(t, user)
}

func TestUserIsEmailAvailabe(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	ok := repo.NewUserRepo(db).IsEmailAvailable(
		existingUser.Email + "adder",
	)

	require.True(t, ok)
}

func TestUserIsEmailNotAvailabe(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	ok := repo.NewUserRepo(db).IsEmailAvailable(existingUser.Email)

	require.False(t, ok)
}
func TestUserUpdateSuccess(t *testing.T) {
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

	require.NotNil(t, err)
	require.Empty(t, user)
}

func TestUserVerifySuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).Verify(
		existingUser.Email,
		existingUser.Password,
	)

	require.Nil(t, err)
	require.NotEmpty(t, user)
}

func TestUserVerifyFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).Verify(
		existingUser.Email,
		existingUser.Password+"adder",
	)

	require.NotNil(t, err)
	require.Empty(t, user)
}

func TestUserDetailSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).Detail(model.User{
		ID: existingUUID,
	})

	require.Nil(t, err)
	require.NotEmpty(t, user)
}

func TestUserDetailFailed(t *testing.T) {
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

func TestUserDetailUsingEmailSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).DetailUsingEmail(existingEmail)

	require.Nil(t, err)
	require.NotEmpty(t, user)
}

func TestUserDetailUsingEmailFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	user, err := repo.NewUserRepo(db).DetailUsingEmail(existingEmail + "adder")

	require.NotNil(t, err)
	require.Empty(t, user)
}

func TestUserDeleteSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewUserRepo(db).Delete(existingUser)

	require.Nil(t, err)
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

	require.NotNil(t, err)
}

package repo_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/handirachmawan/tasks-api-clone/model"
	"github.com/handirachmawan/tasks-api-clone/repo"
	"github.com/stretchr/testify/require"
)

var list model.List

func TestListCreateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	createdList, err := repo.NewListRepo(db).Create(model.List{
		UserID: existingUUID,
		Title:  "New List",
	})

	require.Nil(t, err)
	require.NotEmpty(t, createdList)

	list = createdList
}

func TestListCreateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewListRepo(db).Create(model.List{
		UserID: uuid.New(),
		Title:  "New Failed List",
	})

	require.NotNil(t, err)
	require.Empty(t, list)
}

func TestListFindSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewListRepo(db).Find(uint(1 + rand.Intn(20)))

	require.Nil(t, err)
	require.NotEmpty(t, list)
}

func TestListFindFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewListRepo(db).Find(180)

	require.NotNil(t, err)
	require.Empty(t, list)
}

func TestListUpdateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list.Title = "change"

	list, err := repo.NewListRepo(db).Update(list)

	require.Nil(t, err)
	require.NotEmpty(t, list)
}
func TestListUpdateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewListRepo(db).Update(model.List{
		ID:        0,
		UserID:    uuid.UUID{},
		Title:     "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Tasks:     []model.Task{},
	})

	require.NotNil(t, err)
	require.Empty(t, list)
}

func TestListDeleteSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewListRepo(db).Delete(list)

	require.Nil(t, err)
}

func TestListDeleteFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewListRepo(db).Delete(model.List{
		ID:        0,
		UserID:    [16]byte{},
		Title:     "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Tasks:     []model.Task{},
	})

	require.NotNil(t, err)
}
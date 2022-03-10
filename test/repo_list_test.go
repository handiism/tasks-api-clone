package test

import (
	"math/rand"
	"tasks/model"
	"tasks/repo"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var list model.List

func TestRepoListCreateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	createdList, err := repo.NewList(db).Create(model.List{
		UserID: existingUUID,
		Title:  "New List",
	})

	require.Nil(t, err)
	require.NotEmpty(t, createdList)

	list = createdList
}

func TestRepoListCreateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewList(db).Create(model.List{
		UserID: uuid.New(),
		Title:  "New Failed List",
	})

	require.NotNil(t, err)
	require.Empty(t, list)
}

func TestRepoListFindSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewList(db).Find(uint(1 + rand.Intn(20)))

	require.Nil(t, err)
	require.NotEmpty(t, list)
}

func TestRepoListFindFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewList(db).Find(180)

	require.NotNil(t, err)
	require.Empty(t, list)
}

func TestRepoListUpdateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list.Title = "change"

	list, err := repo.NewList(db).Update(list)

	require.Nil(t, err)
	require.NotEmpty(t, list)
}
func TestRepoListUpdateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewList(db).Update(model.List{
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

func TestRepoListDeleteSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewList(db).Delete(list)

	require.Nil(t, err)
}

func TestRepoListDeleteFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewList(db).Delete(model.List{
		ID:        0,
		UserID:    [16]byte{},
		Title:     "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Tasks:     []model.Task{},
	})

	require.NotNil(t, err)
}

package repo_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/handirachmawan/tasks-api-clone/model"
	"github.com/handirachmawan/tasks-api-clone/repo"
	"github.com/stretchr/testify/assert"
)

var list model.List

func TestListCreateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	createdList, err := repo.NewListRepo(db).Create(model.List{
		UserID: existingUUID,
		Title:  "New List",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, createdList)

	list = createdList
}

func TestListCreateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewListRepo(db).Create(model.List{
		UserID: uuid.New(),
		Title:  "New Failed List",
	})

	assert.NotNil(t, err)
	assert.Empty(t, list)
}

func TestListFindSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewListRepo(db).Find(uint(1 + rand.Intn(20)))

	assert.Nil(t, err)
	assert.NotEmpty(t, list)
}

func TestListFindFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewListRepo(db).Find(180)

	assert.NotNil(t, err)
	assert.Empty(t, list)
}

func TestListUpdateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list.Title = "change"

	list, err := repo.NewListRepo(db).Update(list)

	assert.Nil(t, err)
	assert.NotEmpty(t, list)
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

	assert.NotNil(t, err)
	assert.Empty(t, list)
}

func TestListDeleteSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewListRepo(db).Delete(list)

	assert.Nil(t, err)
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

	assert.NotNil(t, err)
}

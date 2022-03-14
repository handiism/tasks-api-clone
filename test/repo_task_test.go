package test

import (
	"database/sql"
	"math/rand"
	"tasks/model"
	"tasks/repo"
	"testing"

	"github.com/stretchr/testify/require"
)

var createdTask model.Task

func TestRepoTaskCreateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	task, err := repo.NewTaskRepo(db).Create(model.Task{
		ListID: uint(rand.Intn(20) + 1),
		Name:   "New Task",
	})

	require.Nil(t, err)
	require.NotEmpty(t, task)

	createdTask = task
}

func TestRepoTaskCreateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	task, err := repo.NewTaskRepo(db).Create(model.Task{})

	require.NotNil(t, err)
	require.Empty(t, task)
}

func TestRepoTaskFindSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	task, err := repo.NewTaskRepo(db).Find(uint(rand.Intn(20) + 1))

	require.Nil(t, err)
	require.NotEmpty(t, task)
}

func TestRepoTaskFindFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	task, err := repo.NewTaskRepo(db).Find(99)

	require.NotNil(t, err)
	require.Empty(t, task)
}

func TestRepoTaskUpdateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	createdTask.Detail = sql.NullString{
		String: "Change detail",
		Valid:  true,
	}

	createdTask.DueDate = sql.NullTime{
		Time:  createdTask.CreatedAt.AddDate(0, 2, 4),
		Valid: false,
	}

	task, err := repo.NewTaskRepo(db).Update(createdTask)

	require.Nil(t, err)
	require.NotEmpty(t, task)

	createdTask = task
}

func TestRepoTaskUpdateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	task, err := repo.NewTaskRepo(db).Update(model.Task{})

	require.NotNil(t, err)
	require.Empty(t, task)
}

func TestRepoTaskDeleteSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewTaskRepo(db).Delete(createdTask)

	require.Nil(t, err)
}

func TestRepoTaskDeleteFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewTaskRepo(db).Delete(model.Task{})

	require.NotNil(t, err)
}

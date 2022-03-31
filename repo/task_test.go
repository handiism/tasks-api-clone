package repo_test

import (
	"database/sql"
	"math/rand"
	"testing"

	"github.com/handirachmawan/tasks-api-clone/model"
	"github.com/handirachmawan/tasks-api-clone/repo"
	"github.com/stretchr/testify/require"
)

var createdTask model.Task

func TestTaskCreateSuccess(t *testing.T) {
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

func TestTaskCreateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	task, err := repo.NewTaskRepo(db).Create(model.Task{})

	require.NotNil(t, err)
	require.Empty(t, task)
}

func TestTaskFindSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	task, err := repo.NewTaskRepo(db).Find(uint(rand.Intn(20) + 1))

	require.Nil(t, err)
	require.NotEmpty(t, task)
}

func TestTaskFindFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	task, err := repo.NewTaskRepo(db).Find(99)

	require.NotNil(t, err)
	require.Empty(t, task)
}

func TestTaskUpdateSuccess(t *testing.T) {
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

func TestTaskUpdateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	task, err := repo.NewTaskRepo(db).Update(model.Task{})

	require.NotNil(t, err)
	require.Empty(t, task)
}

func TestTaskDeleteSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewTaskRepo(db).Delete(createdTask)

	require.Nil(t, err)
}

func TestTaskDeleteFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewTaskRepo(db).Delete(model.Task{})

	require.NotNil(t, err)
}

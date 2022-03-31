package repo_test

import (
	"math/rand"
	"testing"

	"github.com/handirachmawan/tasks-api-clone/model"
	"github.com/handirachmawan/tasks-api-clone/repo"

	"github.com/stretchr/testify/require"
)

var createdSubtask model.Subtask

func TestSubtaskCreateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	subtask, err := repo.NewSubtaskRepo(db).Create(model.Subtask{
		TaskID: uint(rand.Intn(20) + 1),
		Name:   "New Subtask",
		IsDone: false,
	})

	require.Nil(t, err)
	require.NotEmpty(t, subtask)

	createdSubtask = subtask
}

func TestSubtaskCreateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	subtask, err := repo.NewSubtaskRepo(db).Create(model.Subtask{})

	require.NotNil(t, err)
	require.Empty(t, subtask)
}

func TestSubtaskFindSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	subtask, err := repo.NewSubtaskRepo(db).Find(uint(rand.Intn(20) + 1))

	require.Nil(t, err)
	require.NotEmpty(t, subtask)
}

func TestSubtaskFindFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	subtask, err := repo.NewSubtaskRepo(db).Find(999)

	require.NotNil(t, err)
	require.Empty(t, subtask)
}

func TestSubtaskUpdate(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	createdSubtask.IsDone = true
	createdSubtask.Name = "Change subtask"

	subtask, err := repo.NewSubtaskRepo(db).Update(createdSubtask)

	require.Nil(t, err)
	require.NotEmpty(t, subtask)

	createdSubtask = subtask
}

func TestSubtaskUpdateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	createdSubtask.IsDone = true
	createdSubtask.Name = "Change subtask"

	subtask, err := repo.NewSubtaskRepo(db).Update(model.Subtask{})

	require.NotNil(t, err)
	require.Empty(t, subtask)
}

func TestSubtaskDeleteSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewSubtaskRepo(db).Delete(createdSubtask)

	require.Nil(t, err)
}

func TestSubtaskDeleteFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewSubtaskRepo(db).Delete(model.Subtask{})

	require.NotNil(t, err)
}

package test

import (
	"math/rand"
	"tasks/model"
	"tasks/repo"
	"testing"

	"github.com/stretchr/testify/require"
)

var createdSubtask model.Subtask

func TestRepoSubtaskCreateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	subtask, err := repo.NewSubtask(db).Create(model.Subtask{
		TaskID: uint(rand.Intn(20) + 1),
		Name:   "New Subtask",
		IsDone: false,
	})

	require.Nil(t, err)
	require.NotEmpty(t, subtask)

	createdSubtask = subtask
}

func TestRepoSubtaskCreateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	subtask, err := repo.NewSubtask(db).Create(model.Subtask{})

	require.NotNil(t, err)
	require.Empty(t, subtask)
}

func TestRepoSubtaskFindSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	subtask, err := repo.NewSubtask(db).Find(uint(rand.Intn(20) + 1))

	require.Nil(t, err)
	require.NotEmpty(t, subtask)
}

func TestRepoSubtaskFindFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	subtask, err := repo.NewSubtask(db).Find(999)

	require.NotNil(t, err)
	require.Empty(t, subtask)
}

func TestRepoSubtaskUpdate(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	createdSubtask.IsDone = true
	createdSubtask.Name = "Change subtask"

	subtask, err := repo.NewSubtask(db).Update(createdSubtask)

	require.Nil(t, err)
	require.NotEmpty(t, subtask)

	createdSubtask = subtask
}

func TestRepoSubtaskUpdateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	createdSubtask.IsDone = true
	createdSubtask.Name = "Change subtask"

	subtask, err := repo.NewSubtask(db).Update(model.Subtask{})

	require.NotNil(t, err)
	require.Empty(t, subtask)
}

func TestRepoSubtaskDeleteSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewSubtask(db).Delete(createdSubtask)

	require.Nil(t, err)
}

func TestRepoSubtaskDeleteFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewSubtask(db).Delete(model.Subtask{})

	require.NotNil(t, err)
}

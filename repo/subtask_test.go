package repo_test

import (
	"math/rand"
	"testing"

	"github.com/handirachmawan/tasks-api-clone/model"
	"github.com/handirachmawan/tasks-api-clone/repo"
	"github.com/stretchr/testify/assert"
)

var createdSubtask model.Subtask

func TestSubtaskCreateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	subtask, err := repo.NewSubtaskRepo(db).Store(model.Subtask{
		TaskID: uint(rand.Intn(20) + 1),
		Name:   "New Subtask",
		IsDone: false,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, subtask)

	createdSubtask = subtask
}

func TestSubtaskCreateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	subtask, err := repo.NewSubtaskRepo(db).Store(model.Subtask{})

	assert.NotNil(t, err)
	assert.Empty(t, subtask)
}

func TestSubtaskFindSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	subtask, err := repo.NewSubtaskRepo(db).Find(uint(rand.Intn(20) + 1))

	assert.Nil(t, err)
	assert.NotEmpty(t, subtask)
}

func TestSubtaskFindFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	subtask, err := repo.NewSubtaskRepo(db).Find(999)

	assert.NotNil(t, err)
	assert.Empty(t, subtask)
}

func TestSubtaskUpdate(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	createdSubtask.IsDone = true
	createdSubtask.Name = "Change subtask"

	subtask, err := repo.NewSubtaskRepo(db).Update(createdSubtask)

	assert.Nil(t, err)
	assert.NotEmpty(t, subtask)

	createdSubtask = subtask
}

func TestSubtaskUpdateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	createdSubtask.IsDone = true
	createdSubtask.Name = "Change subtask"

	subtask, err := repo.NewSubtaskRepo(db).Update(model.Subtask{})

	assert.NotNil(t, err)
	assert.Empty(t, subtask)
}

func TestSubtaskDeleteSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewSubtaskRepo(db).Delete(createdSubtask)

	assert.Nil(t, err)
}

func TestSubtaskDeleteFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := repo.NewSubtaskRepo(db).Delete(model.Subtask{})

	assert.NotNil(t, err)
}

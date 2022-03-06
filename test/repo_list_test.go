package test

import (
	"math/rand"
	"tasks/model"
	"tasks/repo"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRepoListCreateSuccess(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewList(db).Create(model.List{
		UserID:    existingUUID,
		Title:     "New List",
		CreatedAt: time.Now(),
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, list)
}

func TestRepoListCreateFailed(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	list, err := repo.NewList(db).Create(model.List{
		UserID: uuid.New(),
		Title:  "New Failed List",
	})

	assert.NotNil(t, err)
	assert.Empty(t, list)
}
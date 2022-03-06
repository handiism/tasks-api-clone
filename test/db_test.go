package test

import (
	"database/sql"
	"math/rand"
	"reflect"
	"tasks/model"
	"tasks/repo"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var randUUID []uuid.UUID = []uuid.UUID{
	uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New(),
	uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New(),
	uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New(),
	uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New(),
}

func TestFillDB(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	if err := db.AutoMigrate(
		&model.User{}, &model.List{}, &model.Task{}, &model.Subtask{},
	); err != nil {
		panic(err)
	}

	err := faker.AddProvider("ref", func(v reflect.Value) (interface{}, error) {
		return uint(1 + rand.Intn(20)), nil
	})

	assert.Nil(t, err)

	for i := 0; i < 20; i++ {
		var user model.User
		err := faker.FakeData(&user)
		assert.Nil(t, err)

		user.ID = randUUID[i]
		user.CreatedAt = time.Now()

		if rand.Intn(4) == i%4 {
			user.UpdatedAt = sql.NullTime{
				Time:  user.CreatedAt.AddDate(0, 0, i+1),
				Valid: true,
			}
		}

		repo := repo.NewUser(db)

		user, err = repo.Create(user)

		assert.Nil(t, err)
		assert.NotEmpty(t, user)
	}

	for i := 0; i < 20; i++ {
		var list model.List

		err := faker.FakeData(&list)

		assert.Nil(t, err)

		list.UserID = randUUID[rand.Intn(20)]
		list.CreatedAt = time.Now()

		if rand.Intn(4) == i%4 {
			list.UpdatedAt = sql.NullTime{
				Time:  list.CreatedAt.AddDate(0, 0, i+1),
				Valid: true,
			}
		}

		repo := repo.NewList(db)

		list, err = repo.Create(list)

		assert.Nil(t, err)
		assert.NotEmpty(t, list)
	}

	for i := 0; i < 20; i++ {
		var task model.Task

		if err := faker.FakeData(&task); err != nil {
			panic(err)
		}

		task.ID = uint(i + 1)
		task.CreatedAt = time.Now()

		sentence, err := faker.GetLorem().Sentence(reflect.Value{})

		assert.Nil(t, err)

		switch s := sentence.(type) {
		case string:
			if i%3 == 0 {
				task.Detail = sql.NullString{
					String: s,
					Valid:  true,
				}
			}
		}

		if rand.Intn(4) == i%4 {
			task.UpdatedAt = sql.NullTime{
				Time:  task.CreatedAt.AddDate(0, 0, i+1),
				Valid: true,
			}
		}

		if !task.UpdatedAt.Valid {
			task.DueDate = sql.NullTime{
				Time:  task.UpdatedAt.Time.AddDate(0, 0, i+1),
				Valid: true,
			}
		}

		if rand.Intn(6) < 4 && task.DueDate.Valid {
			task.IsDone = true
		} else {
			task.IsDone = false
		}

		repo := repo.NewTask(db)

		task, err = repo.Create(task)

		assert.Nil(t, err)
		assert.NotEmpty(t, task)
	}

	for i := 0; i < 20; i++ {
		var subtask model.Subtask

		err := faker.FakeData(&subtask)

		assert.Nil(t, err)

		subtask.ID = uint(i + 1)
		subtask.CreatedAt = time.Now()

		if rand.Intn(4) == i%4 {
			subtask.UpdatedAt = sql.NullTime{
				Time:  subtask.CreatedAt.AddDate(0, 0, i+1),
				Valid: true,
			}
			subtask.IsDone = true
		} else {
			subtask.IsDone = false
		}

		repo := repo.NewSubtask(db)

		subtask, err = repo.Create(subtask)

		assert.Nil(t, err)
		assert.NotEmpty(t, subtask)
	}
}

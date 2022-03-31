package repo_test

import (
	"database/sql"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/handirachmawan/tasks-api-clone/model"
	"github.com/handirachmawan/tasks-api-clone/repo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var randUUID []uuid.UUID = []uuid.UUID{
	uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New(),
	uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New(),
	uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New(),
	uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New(),
}

func init() {
	viper.SetConfigFile(`../.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func openDBConn() *gorm.DB {
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	hostname := viper.GetString("DB_HOSTNAME")
	port := viper.GetString("DB_PORT")

	database := viper.GetString("DB_DATABASE_TEST")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		hostname, user, password, port, database,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	return db
}

func closeDBConn(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic(err)
	}
	conn.Close()
}

func TestMigrate(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := db.AutoMigrate(
		&model.User{}, &model.List{}, &model.Task{}, &model.Subtask{},
	)

	require.Nil(t, err)
}

func TestFillDB(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	err := faker.AddProvider("ref", func(v reflect.Value) (interface{}, error) {
		return uint(1 + rand.Intn(20)), nil
	})

	require.Nil(t, err)

	t.Run("User", func(t *testing.T) {
		for i := 0; i < 20; i++ {
			var user model.User
			err := faker.FakeData(&user)
			require.Nil(t, err)

			user.ID = randUUID[i]
			user.CreatedAt = time.Now()

			if rand.Intn(4) == i%4 {
				user.UpdatedAt = time.Now()
			}

			repo := repo.NewUserRepo(db)

			user, err = repo.Create(user)

			require.Nil(t, err)
			require.NotEmpty(t, user)
		}
	})

	t.Run("List", func(t *testing.T) {
		for i := 0; i < 20; i++ {
			var list model.List

			err := faker.FakeData(&list)

			require.Nil(t, err)

			list.UserID = randUUID[rand.Intn(20)]
			list.CreatedAt = time.Now()

			if rand.Intn(4) == i%4 {
				list.UpdatedAt = time.Now()
			}

			repo := repo.NewListRepo(db)

			list, err = repo.Create(list)

			if err != nil {
				t.FailNow()
			}

			require.Nil(t, err)
			require.NotEmpty(t, list)
		}
	})

	t.Run("Task", func(t *testing.T) {
		for i := 0; i < 20; i++ {
			var task model.Task

			if err := faker.FakeData(&task); err != nil {
				panic(err)
			}

			task.CreatedAt = time.Now()

			sentence, err := faker.GetLorem().Sentence(reflect.Value{})

			require.Nil(t, err)

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
				task.UpdatedAt = time.Now()
				task.DueDate = sql.NullTime{
					Time:  task.UpdatedAt.AddDate(0, 0, i+1),
					Valid: true,
				}
			}

			if rand.Intn(6) < 4 && task.DueDate.Valid {
				task.IsDone = true
			} else {
				task.IsDone = false
			}

			repo := repo.NewTaskRepo(db)

			task, err = repo.Create(task)

			require.Nil(t, err)
			require.NotEmpty(t, task)
		}
	})

	t.Run("Subtask", func(t *testing.T) {
		for i := 0; i < 20; i++ {
			var subtask model.Subtask

			err := faker.FakeData(&subtask)

			require.Nil(t, err)

			subtask.CreatedAt = time.Now()

			if rand.Intn(4) == i%4 {
				subtask.UpdatedAt = time.Now()
				subtask.IsDone = true
			} else {
				subtask.IsDone = false
			}

			repo := repo.NewSubtaskRepo(db)

			subtask, err = repo.Create(subtask)

			require.Nil(t, err)
			require.NotEmpty(t, subtask)
		}
	})
}
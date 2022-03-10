package helper

import (
	"os"
	"strconv"

	"github.com/google/uuid"
)

func ExistingUserUUID() uuid.UUID {
	GetEnv()

	stringUUID := os.Getenv("USER_UUID")

	if stringUUID == "" {
		panic("value of user_uuid is empty")
	}

	uuid, err := uuid.Parse(stringUUID)

	if err != nil {
		panic(err)
	}

	return uuid
}

func ExistingUserEmail() string {
	GetEnv()

	email := os.Getenv("USER_EMAIL")

	if email == "" {
		panic("value of user_email is empty")
	}

	return email
}

func ExistingListID() uint {
	GetEnv()

	id := os.Getenv("LIST_ID")

	if id == "" {
		panic("value of list_id is empty")
	}

	parsedID, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	return uint(parsedID)
}

func ExistingListUserID() uuid.UUID {
	GetEnv()

	stringUUID := os.Getenv("LIST_USER_UUID")

	if stringUUID == "" {
		panic(stringUUID)
	}

	uuid, err := uuid.Parse(stringUUID)

	if err != nil {
		panic(err)
	}

	return uuid
}

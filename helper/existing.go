package helper

import (
	"os"

	"github.com/google/uuid"
)

func ExistingUUID() uuid.UUID {
	GetEnv()

	stringUUID := os.Getenv("EXISTING_UUID")

	if stringUUID == "" {
		panic("value of exisiting_uuid is empty")
	}

	uuid, err := uuid.Parse(stringUUID)

	if err != nil {
		panic(err)
	}

	return uuid
}

func ExistingEmail() string {
	GetEnv()

	email := os.Getenv("EXISTING_EMAIL")
	
	if email == "" {
		panic("value of existing_email is empty")
	}

	return email
}

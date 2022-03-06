package helper

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func GetEnv() {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	base := filepath.Base(dir)
	env := strings.Replace(dir, base, ".env", 1)

	if err := godotenv.Load(env); err != nil {
		panic(err)
	}
}

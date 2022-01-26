package config

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDBConnectionImpl_GetDB(t *testing.T) {
	err := godotenv.Load("../../.env")
	assert.Nil(t, err)

	db := NewDBConnection(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	client, err := db.GetDB()
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

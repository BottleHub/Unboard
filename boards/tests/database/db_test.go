package database_test

import (
	"testing"

	"github.com/bottlehub/unboard/boards/internal/database"
	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	_, err := database.ConnectDB()
	assert.NoError(t, err)
}

package mq_test

import (
	"testing"

	"github.com/bottlehub/unboard/boards/internal/mq"
	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	err := mq.Publish("TestQueue", "just testing...")
	assert.NoError(t, err)
}

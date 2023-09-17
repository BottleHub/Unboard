package mq_test

import (
	"testing"

	"github.com/bottlehub/unboard/boards/internals/mq"
	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	code, err := mq.Publish("TestQueue", "testing...")
	assert.NoError(t, err)
	assert.Equal(t, 0, code)
}

package mq_test

import (
	"testing"

	"github.com/bottlehub/unboard/users/internal/mq"
	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	err := mq.Publish("TestQueue", "testing...")
	assert.NoError(t, err)
}

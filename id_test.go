package letter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncrementIdStartingAt(t *testing.T) {
	id := []byte("XZZ")
	result := string(incrementIdStartingAt(2, id))
	assert.Equal(t, "XZAA", result)
}

func TestNextIdInvalid(t *testing.T) {
	id := "a1"
	result, err := NextId(id)
	assert.NotNil(t, err)
	assert.Equal(t, "a1", result)
}

func TestNextIdLowercase(t *testing.T) {
	id := "a"
	result, err := NextId(id)
	assert.Nil(t, err)
	assert.Equal(t, "B", result)
}

func TestNextIdSimple(t *testing.T) {
	id := "A"
	result, err := NextId(id)
	assert.Nil(t, err)
	assert.Equal(t, "B", result)
}

func TestNextIdRolloverAppend(t *testing.T) {
	id := "Z"
	result, err := NextId(id)
	assert.Nil(t, err)
	assert.Equal(t, "AA", result)
}

func TestNextIdRolloverAppendLarge(t *testing.T) {
	id := "ZZZZZZZZZZZZZZZZZZ" /* 18 chars*/
	result, err := NextId(id)
	assert.Nil(t, err)
	// result should be 19 A's
	assert.Equal(t, "AAAAAAAAAAAAAAAAAAA", result)
}

func TestNextIdRolloverTwo(t *testing.T) {
	id := "ZA"
	result, err := NextId(id)
	assert.Nil(t, err)
	assert.Equal(t, "AB", result)
}

func TestNextIdRolloverThree(t *testing.T) {
	id := "ZZL"
	result, err := NextId(id)
	assert.Nil(t, err)
	assert.Equal(t, "AAM", result)
}

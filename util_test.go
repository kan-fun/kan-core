package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAllWhiteChar(t *testing.T) {
	assert.Equal(t, true, IsAllWhiteChar(" 	 "))
	assert.Equal(t, false, IsAllWhiteChar(" 	6 "))
}

package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAllWhiteChar(t *testing.T) {
	assert.Equal(t, true, isAllWhiteChar(" 	 "))
	assert.Equal(t, false, isAllWhiteChar(" 	6 "))
}

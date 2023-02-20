package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRentalAPI(t *testing.T) {
	res := GetRentalData("Boston", "MA", 1)
	assert.Equal(t, 1, len(res))
}

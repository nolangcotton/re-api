package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSalesAPI(t *testing.T) {
	res := GetSalesData("Boston", "MA", 1)
	assert.Equal(t, 1, len(res))
}

package p765_couples_holding_hands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	row := []int{0, 2, 1, 3}
	assert.Equal(t, 1, minSwapsCouples(row))
}

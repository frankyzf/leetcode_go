package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	want := 5
	got := maxProfit(prices)
	assert.Equal(t, got, want, "equal")
}

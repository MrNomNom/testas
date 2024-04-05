package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChild1(t *testing.T) {
	child := newChildService()
	assert.Equal(t, "c1", child.Child1())
}

func TestChild2(t *testing.T) {
	child := newChildService()
	assert.Equal(t, "c2", child.Child2())
}

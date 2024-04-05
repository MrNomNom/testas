package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ChildServiceMockImpl struct {
}

func (cs ChildServiceMockImpl) Child1() string {
	return "test-c1"
}

func (cs ChildServiceMockImpl) Child2() string {
	return "test-c2"
}

func newChildServiceMock() ChildService {
	return ChildServiceMockImpl{}
}

func TestMainCase1(t *testing.T) {
	mock := newChildServiceMock()
	main := newMainService(mock)
	assert.Equal(t, "test-c1", main.Main(true))
}

func TestMainCase2(t *testing.T) {
	mock := newChildServiceMock()
	main := newMainService(mock)
	assert.Equal(t, "test-c2", main.Main(false))
}

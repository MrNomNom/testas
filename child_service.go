package main

import "fmt"

type ChildService interface {
	Child1() string
	Child2() string
}

type ChildServiceImpl struct {
}

func (cs ChildServiceImpl) Child1() string {
	fmt.Printf("c1")
	return "c1"
}

func (cs ChildServiceImpl) Child2() string {
	fmt.Printf("c2")
	return "c2"
}

func newChildService() ChildService {
	return ChildServiceImpl{}
}

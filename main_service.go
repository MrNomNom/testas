package main

type MainService interface {
	Main(d bool) string
}

type MainServiceImpl struct {
	c ChildService
}

func (ms MainServiceImpl) Main(d bool) string {
	if d {
		return ms.c.Child1()
	} else {

		return ms.c.Child2()
	}
}

func newMainService(c ChildService) MainService {
	return MainServiceImpl{c}
}

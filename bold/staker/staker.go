package staker

import "fmt"

type Staker struct{}

func New() *Staker {
	return &Staker{}
}

func (s *Staker) Foo() {
	fmt.Println("Foo")
}

func (s *Staker) Bar() string {
	return "Bar"
}

func (s *Staker) Baz() string {
	return "Baz"
}

package golang

import "fmt"

type Foo struct {
	A, B chan struct{}
}

func (f *Foo) First() {
	fmt.Println("First")
	f.A <- struct{}{}
}

func (f *Foo) Second() {
	<-f.A
	fmt.Println("Second")
	f.B <- struct{}{}
}

func (f *Foo) Third() {
	<-f.B
	fmt.Println("Third")
}

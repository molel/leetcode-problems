package golang

import (
	"fmt"
	"sync"
)

type FooBar struct {
	N       int
	FooChan chan struct{}
	BarChan chan struct{}
}

func (fb *FooBar) PrintFoo() {
	<-fb.FooChan
	fmt.Print("foo")
	fb.BarChan <- struct{}{}
}

func (fb *FooBar) PrintBar() {
	<-fb.BarChan
	fmt.Print("bar\n")
	fb.FooChan <- struct{}{}
}

func (fb *FooBar) Start() {
	wait := sync.WaitGroup{}
	for i := 0; i < fb.N*2; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			fb.PrintFoo()
		}()
		go func() {
			defer wait.Done()
			fb.PrintBar()
		}()
	}
	fb.FooChan <- struct{}{}
	wait.Wait()
}

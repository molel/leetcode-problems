package medium

import (
	"fmt"
	"sync"
)

type ZeroEvenOdd struct {
	N        int
	ZeroChan chan struct{}
	EvenChan chan struct{}
	OddChan  chan struct{}
	WG       sync.WaitGroup
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	return &ZeroEvenOdd{N: n, ZeroChan: make(chan struct{}), EvenChan: make(chan struct{}), OddChan: make(chan struct{}), WG: sync.WaitGroup{}}
}

func (zeo *ZeroEvenOdd) Zero(PrintNumber func(number int)) {
	defer zeo.WG.Done()
	for n := 1; n <= zeo.N; n++ {
		<-zeo.ZeroChan
		PrintNumber(0)
		if n%2 == 0 {
			zeo.EvenChan <- struct{}{}
		} else {
			zeo.OddChan <- struct{}{}
		}
	}
}

func (zeo *ZeroEvenOdd) Even(PrintNumber func(number int)) {
	defer zeo.WG.Done()
	for n := 2; n <= zeo.N; n += 2 {
		<-zeo.EvenChan
		PrintNumber(n)
		if n != zeo.N {
			zeo.ZeroChan <- struct{}{}
		}
	}
}

func (zeo *ZeroEvenOdd) Odd(PrintNumber func(number int)) {
	defer zeo.WG.Done()
	for n := 1; n <= zeo.N; n += 2 {
		<-zeo.OddChan
		PrintNumber(n)
		if n != zeo.N {
			zeo.ZeroChan <- struct{}{}
		}
	}
}

func (zeo *ZeroEvenOdd) Start() {
	printNumber := func(number int) { fmt.Print(number) }
	zeo.WG.Add(3)
	go zeo.Zero(printNumber)
	zeo.ZeroChan <- struct{}{}
	go zeo.Even(printNumber)
	go zeo.Odd(printNumber)
	zeo.WG.Wait()
}

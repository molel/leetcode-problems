package medium

import (
	"fmt"
	"sync"
)

type H2O struct {
	HydrogenChan chan struct{}
	OxygenChan   chan struct{}
	WG           sync.WaitGroup
}

func NewH2O() *H2O {
	return &H2O{
		HydrogenChan: make(chan struct{}, 2),
		OxygenChan:   make(chan struct{}, 1),
		WG:           sync.WaitGroup{},
	}
}

func (h2o *H2O) Hydrogen(wg *sync.WaitGroup) {
	defer wg.Done()
	defer h2o.WG.Done()
	<-h2o.HydrogenChan
	fmt.Print("H")
}

func (h2o *H2O) Oxygen() {
	for {
		<-h2o.OxygenChan
		wg := sync.WaitGroup{}
		fmt.Print("O")
		h2o.WG.Done()
		wg.Add(2)
		h2o.HydrogenChan <- struct{}{}
		go h2o.Hydrogen(&wg)
		h2o.HydrogenChan <- struct{}{}
		go h2o.Hydrogen(&wg)
		wg.Wait()
	}
}

func (h2o *H2O) Start() {
	go h2o.Oxygen()
	var input string
	fmt.Scan(&input)
	h2o.WG.Add(len(input))
	for _, v := range input {
		if v == 'O' {
			h2o.OxygenChan <- struct{}{}
		}
	}
	h2o.WG.Wait()
}

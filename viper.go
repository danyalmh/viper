package viper

import "time"

// InStream is Type of Input
type InStream string

// Viper is struct of pool
type Viper struct {
	Gate       chan InStream
	nbabies    byte
	diesignals chan byte
}

func babyViper(nest <-chan InStream, hunt func(stream InStream), dieSignal <-chan byte) {
	for {
		select {
		case prey := <-nest:
			hunt(prey)
		case <-dieSignal:
			return
		}
	}
}

// NewNest This is create new Nest/Pool
func NewNest(nbuff int, nbabies byte) *Viper {
	return &Viper{
		Gate:       make(chan InStream, nbuff),
		nbabies:    nbabies,
		diesignals: make(chan byte, nbabies)}
}

// Spawn make babies/worker
func (v *Viper) Spawn(hunt func(stream InStream)) {
	var i byte
	for i = 0; i < v.nbabies; i++ {
		go babyViper(v.Gate, hunt, v.diesignals)
	}

}

// Kill send dieSignal to N babies
func (v *Viper) Kill(num byte) {
	var i byte
	for i = 0; i < num; i++ {
		v.diesignals <- 1
	}
}

// CloseNest send dieSignal to all babies
func (v *Viper) CloseNest() {
	var i byte
	for i = 0; i < v.nbabies; i++ {
		v.diesignals <- 1
	}

	time.Sleep(1 * time.Second)
}

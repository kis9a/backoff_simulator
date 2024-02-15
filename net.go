package main

import (
	"math"
	"math/rand"
)

type Net struct {
	mean int
	sd   int
}

func newNet(mean int, sd int) *Net {
	return &Net{
		mean: mean,
		sd:   sd,
	}
}

func (self *Net) delay() float64 {
	return math.Abs(rand.Float64()*float64(self.sd-self.mean) + float64(self.mean))
}

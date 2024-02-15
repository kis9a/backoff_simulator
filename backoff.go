package main

import (
	"math"
	"math/rand"
)

type Backoff struct {
	base int
	capa int
	name string
}

type IBackoff interface {
	backoff(n int) float64
	getName() string
}

func (b *Backoff) init(name string, base int, capa int) {
	b.name = name
	b.base = base
	b.capa = capa
}

func (self *Backoff) expo(n int) float64 {
	return math.Min(float64(self.capa), math.Pow(2, float64(n))*float64(self.base))
}

// NoBackoff
type NoBackoff struct {
	Backoff
}

func newNoBackoff(name string, base int, capa int) *NoBackoff {
	b := new(NoBackoff)
	b.init(name, base, capa)
	return b
}

func (self *NoBackoff) backoff(n int) float64 {
	return 0
}

func (self *NoBackoff) getName() string {
	return self.name
}

// ExpoBackoff
type ExpoBackoff struct {
	Backoff
}

func newExpoBackoff(name string, base int, capa int) *ExpoBackoff {
	b := new(ExpoBackoff)
	b.init(name, base, capa)
	return b
}

func (self *ExpoBackoff) backoff(n int) float64 {
	return self.expo(n)
}

func (self *ExpoBackoff) getName() string {
	return self.name
}

// ExpoBackoffEqualJitter
type ExpoBackoffEqualJitter struct {
	Backoff
}

func newExpoBackoffEqualJitter(name string, base int, capa int) *ExpoBackoffEqualJitter {
	b := new(ExpoBackoffEqualJitter)
	b.init(name, base, capa)
	return b
}

func (self *ExpoBackoffEqualJitter) backoff(n int) float64 {
	v := self.expo(n)
	return v/2 + rand.Float64()*(v/2)
}

func (self *ExpoBackoffEqualJitter) getName() string {
	return self.name
}

// ExpoBackoffFullJitter
type ExpoBackoffFullJitter struct {
	Backoff
}

func newExpoBackoffFullJitter(name string, base int, capa int) *ExpoBackoffFullJitter {
	b := new(ExpoBackoffFullJitter)
	b.init(name, base, capa)
	return b
}

func (self *ExpoBackoffFullJitter) backoff(n int) float64 {
	v := self.expo(n)
	return rand.Float64() * v
}

func (self *ExpoBackoffFullJitter) getName() string {
	return self.name
}

// ExpoBackoffDecorr
type ExpoBackoffDecorr struct {
	Backoff
	sleep float64
}

func newExpoBackoffDecorr(name string, base int, capa int) *ExpoBackoffDecorr {
	b := new(ExpoBackoffDecorr)
	b.init(name, base, capa)
	b.sleep = float64(base)
	return b
}

func (self *ExpoBackoffDecorr) backoff(n int) float64 {
	self.sleep = math.Min(float64(self.capa), float64(self.base)+rand.Float64()*float64(self.sleep*2-float64(self.base)))
	return self.sleep
}

func (self *ExpoBackoffDecorr) getName() string {
	return self.name
}

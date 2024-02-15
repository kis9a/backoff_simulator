package main

type Stats struct {
	calls    int
	failures int
}

func newStats() *Stats {
	return &Stats{
		calls:    0,
		failures: 0,
	}
}

func (s *Stats) incrementCalls() {
	s.calls++
}

func (s *Stats) incrementFailures() {
	s.failures++
}

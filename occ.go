package main

type OccRequest struct {
	time    float64
	version int
	client  *OccClient
	succ    bool
	sendTo  func(time float64, data *OccRequest) *OccRequest
}

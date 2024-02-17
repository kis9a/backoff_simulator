package main

type OccRequest struct {
	time    float64
	version int
	succ    bool
	replyTo func(time float64, data *OccRequest) *OccRequest
	sendTo  func(time float64, data *OccRequest) *OccRequest
}

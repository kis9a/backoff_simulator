package main

import (
	"fmt"
	"os"
)

type OccServer struct {
	version int
	net     *Net
	stats   *Stats
	tsFile  *os.File
}

func newOccServer(net *Net, stats *Stats, tsFile *os.File) *OccServer {
	return &OccServer{
		version: 0,
		net:     net,
		stats:   stats,
		tsFile:  tsFile,
	}
}

func (self *OccServer) read(tm float64, req *OccRequest) *OccRequest {
	return &OccRequest{
		time:    tm + self.net.delay(),
		version: self.version,
		succ:    req.succ,
		sendTo:  req.replyTo,
		replyTo: nil,
	}
}

func (self *OccServer) write(tm float64, req *OccRequest) *OccRequest {
	self.tsFile.WriteString(fmt.Sprintf("%d\n", int(tm)))
	self.stats.incrementCalls()
	success := false
	if req.version == self.version {
		self.version++
		success = true
	} else {
		self.stats.incrementFailures()
	}
	return &OccRequest{
		time:    tm + self.net.delay(),
		version: self.version,
		succ:    success,
		sendTo:  req.replyTo,
		replyTo: nil,
	}
}

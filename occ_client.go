package main

type OccClient struct {
	server  *OccServer
	net     *Net
	attempt int
	backoff IBackoff
}

func newOccClient(server *OccServer, net *Net, backoff IBackoff) *OccClient {
	return &OccClient{
		server:  server,
		net:     net,
		attempt: 0,
		backoff: backoff,
	}
}

func (self *OccClient) start(tm float64) *OccRequest {
	return &OccRequest{
		time:    tm + self.net.delay(),
		sendTo:  self.server.read,
		replyTo: self.readRsp,
		version: 0,
		succ:    false,
	}
}

func (self *OccClient) readRsp(tm float64, req *OccRequest) *OccRequest {
	return &OccRequest{
		time:    tm + self.net.delay(),
		sendTo:  self.server.write,
		replyTo: self.writeRsp,
		version: req.version,
		succ:    req.succ,
	}
}

func (self *OccClient) writeRsp(tm float64, req *OccRequest) *OccRequest {
	if req.succ {
		return nil
	} else {
		self.attempt++
		return &OccRequest{
			time:    tm + self.net.delay() + self.backoff.backoff(self.attempt),
			sendTo:  self.server.read,
			replyTo: self.readRsp,
			version: req.version,
			succ:    req.succ,
		}
	}
}

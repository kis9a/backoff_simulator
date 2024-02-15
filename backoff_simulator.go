package main

import (
	"container/heap"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func runSim(queue *Heapq) float64 {
	var tm float64
	for queue.Len() > 0 {
		msg := heap.Pop(queue).(*OccRequest)
		if msg.time < tm {
			panic("Time must move forward")
		}
		tm = msg.time
		nextMsg := msg.sendTo(tm, msg)
		if nextMsg != nil {
			heap.Push(queue, nextMsg)
		}
	}
	return tm
}

func setupSim(
	cleints int,
	backoff_cls IBackoff,
	tsf *os.File,
	stats *Stats,
) (*Heapq, *Stats) {
	net := newNet(10, 2)
	queue := newHeapq()
	server := newOccServer(net, stats, tsf)
	for i := 1; i < cleints+1; i++ {
		client := newOccClient(server, net, backoff_cls)
		heap.Push(queue, client.start(0))
	}
	return queue, stats
}

var backoffTypes = []IBackoff{
	newExpoBackoff("Exponential", 5, 2000),
	newExpoBackoffDecorr("Decorr", 5, 2000),
	newExpoBackoffEqualJitter("EqualJitter", 5, 2000),
	newExpoBackoffFullJitter("FullJitter", 5, 2000),
	newNoBackoff("None", 5, 2000),
}

func main() {
	backoffResults := "backoff_results.csv"
	_, err := os.Stat(backoffResults)
	if os.IsNotExist(err) {
		os.Create(backoffResults)
	}
	f, err := os.OpenFile(backoffResults, os.O_WRONLY|os.O_TRUNC, 0o777)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString("clients,time,calls,Algorithm\n")
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < 20; i++ {
		clients := i * 10
		for _, backoff := range backoffTypes {
			tsFile := "ts_" + backoff.getName()
			_, err := os.Stat(tsFile)
			if os.IsNotExist(err) {
				os.Create(tsFile)
			}

			tsf, err := os.OpenFile(tsFile, os.O_WRONLY|os.O_TRUNC, 0o777)
			defer tsf.Close()
			if err != nil {
				log.Fatal(err)
			}

			var tm float64
			stats := newStats()
			queue := newHeapq()
			for i := 0; i < 100; i++ {
				queue, stats = setupSim(clients, backoff, tsf, stats)
				tm += runSim(queue)
			}
			_, err = f.WriteString(fmt.Sprintf("%d,%d,%d,%s\n", clients, int(float64(tm)/float64(100)), int(float64(stats.calls)/float64(100)), backoff.getName()))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

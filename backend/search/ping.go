package search

import (
	"github.com/sparrc/go-ping"
)

func PingServer(url string) bool {
	pinger, err := ping.NewPinger(url)
	if err != nil {
		panic(err)
	}

	pinger.Count = 3
	pinger.Timeout = 2500000000
	pinger.Run() // blocks until finished
	if pinger.PacketsRecv == 0 {
		return true
	}
	return false
}

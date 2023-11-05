package scan

import (
	"fmt"
	"github.com/go-ping/ping"
	"net"
	"time"
)

type PortResult struct {
	Address string
	Status  bool
}

type PingResult struct {
	Address  string
	PingTime int
	Status   bool
}

func PingPort(address string) PortResult {
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return PortResult{
			Address: address,
			Status:  false,
		}
	}
	defer conn.Close()

	return PortResult{
		Address: address,
		Status:  true,
	}
}

func PingDevice(ip string, timeout int) PingResult {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		fmt.Printf("Error creating pinger: %v\n", err)
		return PingResult{
			Address:  ip,
			PingTime: 0,
			Status:   false,
		}
	}

	pinger.SetPrivileged(true)
	pinger.Count = 1
	pinger.Timeout = time.Duration(timeout) * time.Millisecond

	pinger.Run()
	stats := pinger.Statistics()
	pingTime := int(stats.AvgRtt.Milliseconds())
	portStatus := true

	if stats.PacketLoss == 100 {
		pingTime = 0
		portStatus = false
	}

	return PingResult{
		Address:  ip,
		PingTime: pingTime,
		Status:   portStatus,
	}
}

func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

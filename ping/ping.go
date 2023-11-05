package ping

type PingResult struct {
	Address    string
	PingTime   int
	PortStatus bool
}

func pingAndCheckPorts(address string) PingResult {
	// Реализация проверки портов с использованием go-ping
}

func pingDevice(ip string, timeout int) PingResult {
	// Реализация проверки доступности IP-адреса с использованием go-ping
}

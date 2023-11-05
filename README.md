# LANscan
LANscan is a Go application that performs network-related checks, including ICMP ping to IP addresses and port availability checks. The application provides a simple HTTP API for users to submit IP addresses and ports for testing. It leverages the "go-ping" library to ensure the availability of IP addresses and ports.

Key Features:

1. ICMP Ping: The application performs ICMP ping checks to verify the availability of specified IP addresses.
2. Port Check: It checks the availability of specified ports on the provided IP addresses.
3. HTTP API: NetworkChecker exposes a straightforward HTTP API for easy integration.
4. JSON Responses: Results are presented in JSON format, including ping times and port status.

Usage:

1. Submit IP addresses and ports using HTTP requests to the /ping/ endpoint.
2. Receive JSON responses with ping times and port status.

Getting Started:

1. Clone the repository.
2. Build and run the Go application.
3. Use HTTP requests to interact with the API.

Dependencies:

- Go programming language
- "github.com/go-ping/ping" library

Contributions:

Contributions to the project are welcome. Feel free to submit issues, feature requests, or pull requests.
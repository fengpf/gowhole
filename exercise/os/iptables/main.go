package main

import (
	"bytes"
	"os"
	"os/exec"
)

// Protocol to differentiate between IPv4 and IPv6
type Protocol byte

const (
	ProtocolIPv4 Protocol = iota
	ProtocolIPv6
)

func main() {
	path, err := exec.LookPath(getIptablesCommand(ProtocolIPv4))
	if err != nil {
		panic(err)
	}
	vstring, err := getIptablesVersionString(path)
	if err != nil {
		panic(err)
	}
	println(path, vstring, os.Getenv("PATH"))
}

// getIptablesCommand returns the correct command for the given protocol, either "iptables" or "ip6tables".
func getIptablesCommand(proto Protocol) string {
	if proto == ProtocolIPv6 {
		return "ip6tables"
	} else {
		return "iptables"
	}
}

// Runs "iptables --version" to get the version string
func getIptablesVersionString(path string) (string, error) {
	cmd := exec.Command(path, "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

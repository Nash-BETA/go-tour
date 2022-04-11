package main

import "fmt"

type IPAddr [4]byte

func (I IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", I[0], I[1], I[2], I[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

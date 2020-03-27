package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ip := "docker.io"
	ip2, _ := dnsLookup(ip)
	fmt.Println("docker.io resolves to --> ", ip2)

	//Print a blank line for a space in output
	fmt.Println()

	// Now lookup a fake hostname
	l33t := "badHOstN@me.com"
	dnsResp, _ := dnsLookup(l33t)
	fmt.Println("badHOstN@me.com doesnt resolve and prints a nil pointer, thats uglies --> ", dnsResp)

	fmt.Printf("\n")

	// Now properly handle the nil response for the failed DNS lookup
	dnsResp2, err := dnsLookup(l33t)
	if err != nil {
		log.Printf("lookup failed for [ %s ]", l33t)
	} else {
		fmt.Println("gOOgl3.cOm resolves to --> ", dnsResp2)
	}

	//Print a blank line for a space in output print format
	fmt.Printf("\n")

	//Now properly handle a good response from a successful lookup
	ovs := "openvswitch.org"
	dnsResp3, err1 := dnsLookup(ovs)
	if err1 != nil {
		log.Println("lookup failed")
	} else {
		fmt.Printf("%s resolves to --> %s", ovs, dnsResp3)
	}
}

func dnsLookup(ipStr string) (string, error) {
	ipAddr, err := net.ResolveIPAddr("ip", ipStr)
	return ipAddr.String(), err
}

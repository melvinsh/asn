package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/melvinsh/asn/pkg/asn"
	"github.com/melvinsh/asn/pkg/subnet"
)

func main() {
	var ips bool
	flag.BoolVar(&ips, "ips", false, "Print IPs instead of subnets")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go [-ips] <ASN>")
		return
	}

	subnets, err := asn.FindSubnets(flag.Arg(0))

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if ips {
		printIps(subnets)
	} else {
		printSubnets(subnets)
	}
}

func printIps(subnets []string) {
	for _, net := range subnets {
		ips, err := subnet.GetIPAddressesInSubnet(net)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting IP addresses for subnet %s: %v\n", net, err)
		}
		for _, ip := range ips {
			fmt.Println(ip)
		}
	}
}

func printSubnets(subnets []string) {
	for _, subnet := range subnets {
		fmt.Println(subnet)
	}
}

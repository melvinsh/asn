package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	var ips bool
	flag.BoolVar(&ips, "ips", false, "Print IPs instead of ASN")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go [-ips] <ASN>")
		return
	}

	asn := flag.Args()[0]

	// normalize to correct format
	asn = strings.ToLower(asn)
	if !strings.HasPrefix(asn, "as") {
		asn = "as" + asn
	}

	// make the request to ip2location.com
	resp, err := http.Get("https://www.ip2location.com/" + asn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error making request:", err)
		return
	}
	if resp.StatusCode == 404 {
		fmt.Fprintln(os.Stderr, "Error making request: ASN not found (404)")
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading response body:", err)
		return
	}

	// extract the subnets from the HTML
	re := regexp.MustCompile(`href="/demo/(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})">(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\/\d{1,2})</a>`)
	subnets := re.FindAllStringSubmatch(string(body), -1)

	if ips {
		printIps(subnets)
	} else {
		printSubnets(subnets)
	}
}

func printIps(subnets [][]string) {
	for _, subnet := range subnets {
		ips, _ := getIPAddressesInSubnet(subnet[2])
		for _, ip := range ips {
			fmt.Println(ip)
		}
	}
}

func printSubnets(subnets [][]string) {
	for _, subnet := range subnets {
		fmt.Println(subnet[2])
	}
}

func getIPAddressesInSubnet(subnet string) ([]string, error) {
	ip, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, err
	}
	var ips []string
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	return ips, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

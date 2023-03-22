package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	// get the ASN from the first command-line argument
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <ASN>")
		return
	}
	asn := os.Args[1]

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

	// print the subnets
	for _, subnet := range subnets {
		fmt.Println(subnet[2])
	}
}

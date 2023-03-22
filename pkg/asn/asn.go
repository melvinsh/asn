package asn

import (
	"errors"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func FindSubnets(asn string) ([]string, error) {
	var subnets []string
	asn = normalize(asn)

	resp, err := http.Get("https://www.ip2location.com/" + asn)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, errors.New("ASN not found (404)")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`href="/demo/(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})">(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\/\d{1,2})</a>`)
	matches := re.FindAllStringSubmatch(string(body), -1)
	for _, match := range matches {
		subnets = append(subnets, match[2])
	}

	return subnets, nil
}

func normalize(asn string) string {
	asn = strings.ToLower(asn)
	if !strings.HasPrefix(asn, "as") {
		asn = "as" + asn
	}
	return asn
}

# asn

This is a simple Go program that extracts the subnets associated with a given Autonomous System Number (ASN) from the IP2Location website. 

The program takes the ASN as a command-line argument and outputs a list of subnets associated with that ASN.

## Installation

To install this program, you must have Go installed on your system. You can install the program using the following command:

``` bash
go install github.com/melvinsh/asn@latest
```

This will download the source code, build the program, and install it in your `$GOPATH/bin` directory. Make sure that `$GOPATH/bin` is in your `$PATH` environment variable so that you can run the program from anywhere.

## Usage

To run the program, simply type `asn` followed by the ASN as the argument. For example, to get the subnets for ASN 15169 (Google), you would run:

``` bash
asn 15169
```

## How it works

The program uses the Go `net/http` package to make a GET request to the IP2Location website with the specified ASN. It then reads the response and extracts the subnets using regular expressions. Finally, it outputs the subnets to the console.

Note that the program normalizes the ASN to ensure that it has the correct format (i.e., starts with "AS"). If the program cannot find any subnets associated with the specified ASN, it will output an error message.

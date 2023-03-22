# asn

This is a simple Go program that extracts the subnets associated with a given Autonomous System Number (ASN) from the IP2Location website. 

The program takes the ASN as a command-line argument and outputs a list of subnets associated with that ASN.

## Usage

To use this program, you must have Go installed on your system. To run the program, follow these steps:

1. Clone the repository: `git clone https://github.com/melvinsh/asn.git`
2. Navigate to the repository directory: `cd asn`
3. Build the program: `go build`
4. Run the program with an ASN as the argument: `./asn <ASN>`

For example, to get the subnets for ASN 15169 (Google):

``` bash
$ ./asn 15169
8.8.4.0/24
8.8.8.0/24
8.35.200.0/21
34.0.96.0/19
34.0.160.0/19
34.3.3.0/24
34.4.4.0/24
34.143.64.0/19
34.144.0.0/20
34.144.128.0/17
34.149.0.0/16
```

## How it works

The program uses the Go `net/http` package to make a GET request to the IP2Location website with the specified ASN. It then reads the response and extracts the subnets using regular expressions. Finally, it outputs the subnets to the console.

Note that the program normalizes the ASN to ensure that it has the correct format (i.e., starts with "AS"). If the program cannot find any subnets associated with the specified ASN, it will output an error message.

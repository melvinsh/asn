# asn

This is a simple Go program that extracts the subnets associated with a given Autonomous System Number (ASN) from the IP2Location website. The program takes the ASN as a command-line argument and outputs a list of subnets associated with that ASN.

## Usage

To use this program, you must have Go installed on your system. To run the program, follow these steps:

1. Clone the repository: `git clone https://github.com/melvinsh/asn.git`
2. Navigate to the repository directory: `cd asn`
3. Build the program: `go build`
4. Run the program with an ASN as the argument: `./asn <ASN>`

For example, to get the subnets for ASN 15169 (Google), you would run:

```
./asn 15169
```

## How it works

The program uses the Go `net/http` package to make a GET request to the IP2Location website with the specified ASN. It then reads the response body using the `io` package and extracts the subnets using regular expressions from the `regexp` package. Finally, it outputs the subnets to the console using the `fmt` package.

Note that the program normalizes the ASN to ensure that it has the correct format (i.e., starts with "AS"). If the program cannot find any subnets associated with the specified ASN, it will output an error message to the console.

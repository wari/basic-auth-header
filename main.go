// Package main provides just a simple auth string generator. Nothing much really, just a utility.
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

var (
	msg         = flag.String("msg", "", "String to decode into base64")
	auth_header = flag.Bool("auth", false, "Prints out the Authorization header for Basic Auth")
)

func init() {
	flag.Parse()
	if *msg == "" {
		flag.Usage()
		os.Exit(2)
	}
}

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte(*msg))
	if *auth_header {
		fmt.Println("Authorization: Basic", encoded)
	} else {
		fmt.Println(encoded)
	}
}

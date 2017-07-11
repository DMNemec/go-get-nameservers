package main

import (
	"fmt"
	"net"
)

func main() {
	site := "lsbia.bank"
	servers, err := net.LookupNS(site)
	if err != nil {
		panic(err)
	}
	for _, name := range servers {
		fmt.Println(name.Host)
	}
	// our name servers are *.jackhenry.com and *.bancinternetgroup.com
}

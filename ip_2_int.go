package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	var valu int32
	var ip_ string
	fmt.Print("Enter the ip address:")
	fmt.Scan(&ip_)
	bits := strings.Split(net.ParseIP(ip_).String(), ".")
	l_swift := []int{24, 16, 8, 0}
	for i := 0; len(bits) > i; i++ {
		a, _ := strconv.Atoi(bits[i])
		valu += int32(a << l_swift[i])
	}
	fmt.Print(valu)
}

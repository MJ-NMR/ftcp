package clint

import (
	"fmt"
	"net"
	"os"
)


func Connect(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	conn.Write([]byte{1,0,10})
	buff := make([]byte, 1024)

	for {
		n, err := os.Stdin.Read(buff)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(n)
		n, err = conn.Write(buff[:n-1])
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

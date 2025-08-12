package server

import (
	"encoding/binary"
	"fmt"
	"net"
)


func Start(address string) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		l.Close()
		return
	}
	defer l.Close()
	fmt.Printf("starting server on %v\n", l.Addr().String())

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("connection from %v\n", conn.RemoteAddr().String())

		go handelConnect(&clint{conn: conn})
	}
}

const headerlenght  = 3

type clint struct{
	data []byte
	conn net.Conn
	body []byte
	command int
}

func handelConnect(cl *clint) {
	err := cl.Read()
	if err != nil {
		return
	}

	if cl.command == 1 {
		fmt.Println(string(cl.body))
	}
}


func (c *clint) Read() error {
	buff := make([]byte, 1024)
	for {
		n, err := c.conn.Read(buff)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		c.data = append(c.data, buff[:n]...)
		println("data after: ", len(c.data))
		if err := c.parseData(); err == nil {
			c.conn.Close()
			break
		} else {
			fmt.Println(err.Error())
		}
	}

	return nil
}

func (c *clint) parseData() error {
	dlength := len(c.data)
	if dlength < headerlenght {
		return fmt.Errorf("keep going")
	}
	
	bodyLength := int(binary.BigEndian.Uint16(c.data[1:3]))
	//fmt.Println("boody: ",bodyLength,"daaata: ", dlength)
	if dlength < headerlenght+bodyLength {
		return fmt.Errorf("keep going")
	}

	c.body = c.data[headerlenght:]
	c.command = int(c.data[0])

	return  nil
}

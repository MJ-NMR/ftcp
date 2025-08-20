package ftcp

import (
	"encoding/binary"
	"fmt"
	"io"
)

const headerlenght = 3

func Decode(reader io.Reader) ( int, []byte, error ) {
	buff := make([]byte, 1024)
	var data []byte
	for {
		n, err := reader.Read(buff)
		if err != nil {
			return 0, []byte{}, err
		}

		data = append(data, buff[:n]...)
		if command, body, err := parseData(data); err == nil {
			return command, body, nil
		} else {
		}
	}
}

func parseData(data []byte) ( int, []byte, error ) {
	dlength := len(data)
	if dlength < headerlenght {
		return 0, []byte{}, fmt.Errorf("keep going")
	}
	
	bodyLength := int(binary.BigEndian.Uint16(data[1:3]))
	if dlength < headerlenght+bodyLength {
		return 0, []byte{}, fmt.Errorf("keep going")
	}

	body := data[headerlenght:]
	command := int(data[0])

	return  command, body, nil
}

func Incode(command int8, body []byte) []byte {
	data := []byte{}
	data = append(data, byte(command))
	bodyLen := len(body)
	data = append(data, byte(bodyLen))
	data = append(data, body...)
	return data
}

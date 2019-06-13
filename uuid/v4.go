package uuid

import (
	"fmt"
	"io"
)

func CreateV4(source io.Reader) (string, error) {
	data := make([]byte, 16)
	if _, err := source.Read(data); err != nil {
		return "", err
	}
	data[8] = (data[8] | 0x80) & 0xbf
	data[6] = (data[6] | 0x40) & 0x4f

	uuid := fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", data[0:4], data[4:6], data[6:8], data[8:10],  data[10:])
	return uuid, nil
}
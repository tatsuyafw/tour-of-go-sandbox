package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (reader MyReader) Read(buffer []byte) (int, error) {
	length := len(buffer)
	for i := 0; i < length; i++ {
		buffer[i] = 'A'
	}
	return length, nil
}

func main() {
	reader.Validate(MyReader{})
}

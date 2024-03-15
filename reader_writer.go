package main

import (
	"errors"
	"fmt"
	"io"
	"math"
)

type writerReaderImpl struct {
	data string
}

func (w *writerReaderImpl) Write(p []byte) (int, error) { // this method writes the slice of bytes to the w which is an object of type writerReaderImpl
	w.data = string(p) //p is converted to string because this string is a type of byte slice
	return len(p), io.EOF
}

func (w *writerReaderImpl) Read(p []byte) (int, error) { // this method reads the w object data into p which is a slice of bytes
	temp := []byte(w.data)
	threshold := int(math.Min(float64(cap(p)), float64(len(temp))))
	copy(p, temp[:threshold])
	return len(p), io.EOF
}

func main() {
	myWriterReader := writerReaderImpl{}
	_, err := fmt.Fprintln(&myWriterReader, "Later please add this to your object")
	if err != nil && !errors.Is(err, io.EOF) {
		fmt.Println(err)
		return
	}
	fmt.Printf("the object value is %v\n", myWriterReader)

	buffer := make([]byte, 15)
	_, err = myWriterReader.Read(buffer)
	if err != nil && !errors.Is(err, io.EOF) {
		fmt.Println(err)
		return
	}
	fmt.Printf("the object value is %v\n", myWriterReader)
	fmt.Println(string(buffer[:]))
}

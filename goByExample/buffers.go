package gobyexample

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func reader() {
	msg := "1 Hello buffer friends!"
	// fmt.Println(msg)

	reader := strings.NewReader(msg)
	var first_word string
	var second_word string
	var third_word string

	n, err := fmt.Fscan(reader, &first_word, &second_word, &third_word)
	fmt.Println(err, n)

	fmt.Println(first_word)
	fmt.Println(second_word)
	fmt.Println(third_word)
}

func writer() {
	write_buffer := new(bytes.Buffer)

	first_word := "hi!"
	second_word := 2
	third_word := 3

	n, err := fmt.Fprint(write_buffer, first_word, second_word, third_word)

	fmt.Println(n, err)
	fmt.Println(write_buffer.String())

}

func compress() {
	src := []byte("squish me please!")

	//compress src
	buf := new(bytes.Buffer) // Create empty buffer
	w := gzip.NewWriter(buf) // Create a write with a buffer as output
	w.Write(src)             // Write src to writer
	w.Close()                // close writer

	gzipReader, _ := gzip.NewReader(buf)
	gzipReader.Close()

	// output, err := ioutil.ReadAll(gzipReader)
	// fmt.Println(string(output), err)
	io.Copy(os.Stdout, gzipReader)
}

func Buffers() {
	compress()
}

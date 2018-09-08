// io package samples
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func sampleCopy() {
	// io.Copy
	r := strings.NewReader("with NewReader\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	// with file read
	file, err := os.Open("io/some.txt")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(os.Stdout, file); err != nil {
		log.Fatal(err)
	}
}

func sampleCopyBuffer() {
	// io.CopyBuffer
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}

	// if buf is zero length, CopyBuffer panics.
	// buf = make([]byte, 0)
	// io.CopyBuffer(os.Stdout, r1, buf)
}

func sampleCopyN() {
	// io.CopyN
	buf := &bytes.Buffer{}
	r := strings.NewReader("まるちばいともじれつ") // with multi byte string

	if _, err := io.CopyN(buf, r, 6); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buf.Bytes()) // output "まる"
}

func sampleReadAtLeast() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 10)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	buf2 := make([]byte, 23)
	if _, err := io.ReadAtLeast(r, buf2, 4); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf2)

	// buffer smaller than minimal read size.
	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
		fmt.Println("error:", err) // short buffer
	}

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 23)
	if _, err := io.ReadAtLeast(r, longBuf, 23); err != nil {
		fmt.Println("error:", err) // EOF
	}

	r2 := strings.NewReader("shortmsg")

	buf = make([]byte, 9)
	if _, err := io.ReadAtLeast(r2, buf, 9); err != nil {
		fmt.Println("error:", err) // unexpected EOF
	}
}

func main() {
	sampleCopy()
	sampleCopyBuffer()
	sampleCopyN()
	sampleReadAtLeast()
}

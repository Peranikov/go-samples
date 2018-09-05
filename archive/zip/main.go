// archive/zip package sample
package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	flag.Parse()

	sub := flag.Arg(0)

	switch sub {
	case "archive":
		archive(flag.Args()[1], flag.Args()[2:])
	case "unarchive":
		unarchive(flag.Arg(1))
	}

	fmt.Printf("Unsupported commnd: %s", sub)
}

func archive(zipName string, files []string) {
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	zw.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	for _, file := range files {
		func() {
			f, err := os.Open(file)
			if err != nil {
				panic(err) // not recommend

			}
			defer f.Close()

			w, err := zw.Create(f.Name())
			if err != nil {
				panic(err)
			}

			b, err := ioutil.ReadAll(f)
			if err != nil {
				panic(err)
			}

			if _, err := w.Write(b); err != nil {
				panic(err)
			}
		}()
	}

	if err := zw.Close(); err != nil {
		panic(err)
	}

	ioutil.WriteFile(zipName, buf.Bytes(), 0600)
}

func unarchive(zipFile string) {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	for _, f := range r.File {
		func(f *zip.File) {
			rc, err := f.Open()
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(os.Stdout, rc)
			if err != nil {
				panic(err)
			}

			rc.Close()
		}(f)

	}
}

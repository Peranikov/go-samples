// archive/tar package sample
package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
)

func unarchive() {
	tr := tar.NewReader(os.Stdin)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}

		func(hdr *tar.Header) {
			file, err := os.Create(hdr.Name)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			io.Copy(file, tr)
			file.Chmod((os.FileMode)(hdr.Mode))
			file.Chown(hdr.Uid, hdr.Gid)
		}(hdr)
	}
}

func main() {
	unarchive()
}

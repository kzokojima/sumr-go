package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func md5String(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func main() {
	files, err := filepath.Glob("*")
	if err != nil {
		panic(err)
	}
	for _, v := range files {
		if fileinfo, err := os.Stat(v); err != nil {
			panic(err)
		} else if fileinfo.IsDir() {
			//
		} else {
			if sum, err := md5String(v); err != nil {
				panic(err)
			} else {
				fmt.Printf("%s\t%s\n", v, sum)
			}
		}
	}
}

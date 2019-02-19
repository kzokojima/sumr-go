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

func writeSumRecursive(file *os.File, dir string) error {
	files, err := filepath.Glob(fmt.Sprintf("%s/*", dir))
	if err != nil {
		return err
	}
	for _, v := range files {
		if fileinfo, err := os.Stat(v); err != nil {
			return err
		} else if fileinfo.IsDir() {
			//
		} else {
			if sum, err := md5String(v); err != nil {
				return err
			} else {
				file.Write([]byte(fmt.Sprintf("%s/%s\t%s\n", dir, v, sum)))
			}
		}
	}

	return nil
}

func main() {
	writeSumRecursive(os.Stdout, "PATH\tCHECKSUM (md5)\n")
	writeSumRecursive(os.Stdout, ".")
}

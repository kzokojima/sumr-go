package main

import (
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const usage = "usage: sumr [-a algo] [dir]"

var ignore = []string{"desktop.ini", "Thumbs.db", ".DS_Store"}

func exit(message string, code int) {
	fmt.Println(message)
	os.Exit(code)
}

func newHash(algo string) hash.Hash {
	var h hash.Hash
	switch algo {
	case "md5":
		h = md5.New()
	case "sha256":
		h = sha256.New()
	default:
		exit(usage, 1)
	}
	return h
}

func hashString(path string, algo string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := newHash(algo)
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func writeSumRecursive(file *os.File, algo string) error {
	err := filepath.Walk(".", func(path string, fileinfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fileinfo.IsDir() {
			for _, each := range ignore {
				if each == fileinfo.Name() {
					return nil
				}
			}
			if sum, err := hashString(path, algo); err != nil {
				return err
			} else {
				cd, _ := os.Getwd()
				path = strings.Replace(path, cd, "", 1)
				file.Write([]byte(fmt.Sprintf("./%s\t%s\n", path, sum)))
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	const defaultAlgo = "md5"
	var algo string
	var dir = string('.')

	// command line options
	flag.StringVar(&algo, "a", defaultAlgo, usage)
	flag.Parse()
	switch len(flag.Args()) {
	case 0:
	case 1:
		dir = flag.Arg(0)
	default:
		exit(usage, 1)
	}
	if fileinfo, err := os.Stat(dir); err != nil {
		exit(usage, 1)
	} else if !fileinfo.IsDir() {
		exit(usage, 1)
	}
	newHash(algo)

	os.Stdout.Write([]byte(fmt.Sprintf("PATH\tCHECKSUM (%s)\n", algo)))
	os.Chdir(dir)
	writeSumRecursive(os.Stdout, algo)
}

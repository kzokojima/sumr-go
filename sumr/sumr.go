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

func writeSumRecursive(file *os.File, dir string, algo string) error {
	// escape
	replacer := strings.NewReplacer(
		"[", "\\[",
		"]", "\\]",
	)
	dir = replacer.Replace(dir)

	files, err := filepath.Glob(fmt.Sprintf("%s/*", dir))
	if err != nil {
		return err
	}
OUTER:
	for _, v := range files {
		if fileinfo, err := os.Stat(v); err != nil {
			return err
		} else if fileinfo.IsDir() {
			writeSumRecursive(file, fmt.Sprintf("%s/%s", dir, v), algo)
		} else {
			for _, each := range ignore {
				if each == v {
					continue OUTER
				}
			}
			if sum, err := hashString(v, algo); err != nil {
				return err
			} else {
				file.Write([]byte(fmt.Sprintf("./%s\t%s\n", v, sum)))
			}
		}
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
	writeSumRecursive(os.Stdout, dir, algo)
}

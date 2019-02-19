package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func checkMd5String(t *testing.T, path string, data string, expected string) {
	if err := ioutil.WriteFile(path, []byte(data), 0644); err != nil {
		panic(err)
	}
	if actual, err := md5String(path); err != nil {
		panic(err)
	} else if actual != expected {
		t.Errorf("got: %s\nwant: %s", actual, expected)
	}
}

func TestMd5String(t *testing.T) {
	defer os.RemoveAll("tmp")
	os.Mkdir("tmp", 0755)
	checkMd5String(t, "tmp/1.txt", "", "d41d8cd98f00b204e9800998ecf8427e")
	checkMd5String(t, "tmp/2.txt", "foo", "acbd18db4cc2f85cedef654fccc4a4d8")
}

func TestMd5StringR(t *testing.T) {
	// expected
	defer os.Remove("expected.txt")
	if file, err := os.Create("expected.txt"); err != nil {
		panic(err)
	} else {
		file.Write([]byte("./1.txt\td41d8cd98f00b204e9800998ecf8427e\n"))
		file.Write([]byte("./2.txt\tacbd18db4cc2f85cedef654fccc4a4d8\n"))
		file.Write([]byte("./sub/1.txt\td41d8cd98f00b204e9800998ecf8427e\n"))
		file.Write([]byte("./sub/2.txt\tacbd18db4cc2f85cedef654fccc4a4d8\n"))
	}

	// actual
	defer os.Remove("actual.txt")
	if file, err := os.Create("actual.txt"); err != nil {
		panic(err)
	} else {
		defer os.RemoveAll("tmp")
		os.Mkdir("tmp", 0755)
		os.Mkdir("tmp/sub", 0755)
		os.Chdir("tmp")
		ioutil.WriteFile("1.txt", []byte(""), 0644)
		ioutil.WriteFile("2.txt", []byte("foo"), 0644)
		ioutil.WriteFile("sub/1.txt", []byte(""), 0644)
		ioutil.WriteFile("sub/2.txt", []byte("foo"), 0644)
		writeSumRecursive(file, ".")
		os.Chdir("..")
	}

	// assert
	if expected, err := ioutil.ReadFile("expected.txt"); err != nil {
		panic(err)
	} else if actual, err := ioutil.ReadFile("actual.txt"); err != nil {
		panic(err)
	} else {
		if !bytes.Equal(expected, actual) {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	}
}

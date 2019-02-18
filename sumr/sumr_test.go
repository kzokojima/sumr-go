package main

import (
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

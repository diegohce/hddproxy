package main

import (
	"io/ioutil"
	"path"
	"testing"
)

func TestFileWriterTo(t *testing.T) {

	var dir = "/home/diego/tmp/stuff"
	var rawFile = []byte(`{"dir":"/home/diego/tmp/stuff", "name":"outputfile.txt", "body":"Hola mundo file content"}`)

	f, err := hddFileFromJSON(rawFile)
	if err != nil {
		t.Error(err)
	}

	if err := f.writeTo(dir); err != nil {
		t.Error(err)
	}

	fullname := path.Join(dir, f.Name)

	b, err := ioutil.ReadFile(fullname)
	if err != nil {
		t.Error(err, "reading", fullname)
	}

	if string(b) != "Hola mundo file content" {
		t.Error("File content mismatch")
	}

}

func TestFileWriter(t *testing.T) {

	var rawFile = []byte(`{"dir":"/home/diego/tmp/stuff", "name":"outputfile2.txt", "body":"Hola mundo file content"}`)

	f, err := hddFileFromJSON(rawFile)
	if err != nil {
		t.Error(err)
	}

	if err := f.write(); err != nil {
		t.Error(err)
	}

	fullname := path.Join(f.Dir, f.Name)

	b, err := ioutil.ReadFile(fullname)
	if err != nil {
		t.Error(err, "reading", fullname)
	}

	if string(b) != "Hola mundo file content" {
		t.Error("File content mismatch")
	}

}

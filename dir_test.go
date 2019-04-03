package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"testing"
	"time"

	"github.com/diegohce/logger"
)

var (
	dir string = "/home/diego/tmp/stuff"
)

func TestDirScanner(t *testing.T) {

	scanner := &dirScanner {
		dir:           dir,
		shortPollWait: time.Duration(1*time.Second),
		longPollWait:  time.Duration(5*time.Second),
	}

	ch := scanner.start()

	i := 0
	for f := range ch {
		fmt.Printf(f)
		if i == 2 {
			scanner.stop()
		}
		i++
	}
}


func TestMain(m *testing.M) {

	log = logger.New("kpm-proxy-test - ")

	ioutil.WriteFile(path.Join(dir, "file1.txt"), []byte("hola mundo1"), 0600)
	ioutil.WriteFile(path.Join(dir, "file2.txt"), []byte("hola mundo2"), 0600)
	ioutil.WriteFile(path.Join(dir, "file3.txt"), []byte("hola mundo3"), 0600)

	m.Run()
}


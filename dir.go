package main

import (
	"io/ioutil"
	"os"
	"path"
	"time"
)

type dirScanner struct {
	dir           string
	shortPollWait time.Duration
	longPollWait  time.Duration
	quit          bool
	wait          chan bool
}

func (ds *dirScanner) start() chan string {

	ch := make(chan string)
	ds.wait = make(chan bool)
	ds.quit = false

	go ds.dirWorker(ch)

	return ch
}

func (ds *dirScanner) startWithChannel(ch chan string) chan string {

	//ch := make(chan string)
	ds.wait = make(chan bool)
	ds.quit = false

	go ds.dirWorker(ch)

	return ch
}

func (ds *dirScanner) stop() {
	ds.quit = true
	<-ds.wait
}

func (ds *dirScanner) dirWorker(ch chan string) {

	defer func() {
		if r := recover(); r != nil {
			log.Debug().Println(r)
		}
	}()

	for {
		if ds.quit {
			close(ch)
			ds.wait <- true
			return
		}
		files, err := ioutil.ReadDir(ds.dir)
		if err != nil {
			log.Error().Println(err, "scanning", ds.dir)
			close(ch)
			return
		}

		for _, f := range files {
			if ds.quit {
				close(ch)
				ds.wait <- true
				return
			}

			fullName := path.Join(ds.dir, f.Name())

			if f.IsDir() {
				log.Info().Println("Ignoring dir", fullName)
				continue
			}

			log.Info().Println("Reading", fullName, "content")

			b, err := ioutil.ReadFile(fullName)
			if err != nil {
				log.Error().Println(err, "reading", fullName)
			}

			f := &hddFile{
				Dir:  ds.dir,
				Name: f.Name(),
				Body: string(b),
			}

			if j, err := f.Marshal(); err != nil {
				log.Error().Println(err, "Marshaling", fullName)
			} else {
				ch <- string(j)

				if err := os.Remove(fullName); err != nil {
					log.Error().Println(err, "deleting", fullName)
				}
			}
		}

		if len(files) > 0 {
			log.Info().Println("Will wait for", ds.shortPollWait)
			time.Sleep(ds.shortPollWait)
		} else {
			log.Info().Println("Will wait for", ds.longPollWait)
			time.Sleep(ds.longPollWait)
		}
	}
}

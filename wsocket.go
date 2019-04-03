package main

import (
	"net/http"
	//"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


func wsocketHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error().Println(err)
		return
	}

	/*for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		log.Info().Printf("%s type: %d sent: %s\n", conn.RemoteAddr(), msgType, string(msg))

		 Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}*/

	var ch chan string
	scanners := make([]*dirScanner, len(config.dirs))

	for i, d := range config.dirs {
		log.Info().Println("Starting scanner for", d)

		scanners[i] = &dirScanner {
			dir:           d,
			longPollWait:  config.longPoll,
			shortPollWait: config.shortPoll,
		}
		if i == 0 {
			ch = scanners[i].start()
		} else {
			scanners[i].startWithChannel(ch)
		}
	}

	go func() {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				for _, s := range scanners {
					log.Info().Println("Stopping scanner for", s.dir)
					s.stop()
				}
				break
			}
		}
	}()

	for f := range ch {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(f)); err != nil {
			log.Error().Println(err)
		}
	}

	/*for _, s := range scanners {
		log.Info().Println("Stopping scanner for", s.dir)
		s.stop()
	}*/


/*	scanner := &dirScanner {
		dir:           "/home/diego/tmp/stuff",
		longPollWait:  time.Duration(5*time.Second),
		shortPollWait: time.Duration(1*time.Second),
	}

	for f := range scanner.start() {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(f)); err != nil {
			log.Error().Println(err)
		}
	}
	scanner.stop()*/
}


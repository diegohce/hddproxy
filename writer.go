package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func responseWithError(w http.ResponseWriter, err error) {
	w.WriteHeader(400)
	fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
}

func writeRequest(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseWithError(w, err)
		return
	}
	// no need because i'm a server? defer r.Body.Close()

	log.Info().Println("Received write request len:", len(b))

	f, err := hddFileFromJSON(b)
	if err != nil {
		responseWithError(w, err)
		return
	}

	if err := f.write(); err != nil {
		responseWithError(w, err)
		return
	}

	log.Info().Println(f.Name, "written into", f.Dir)
}

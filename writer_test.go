package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
)

func TestWriterRequest(t *testing.T) {

	/*f := &kmpFile{
		Dir:  dir,
		Name: "testwriter.txt",
		Body: "Hello, World!",
	}*/

	var rawFile = []byte(`{"dir":"/home/diego/tmp/stuff", "name":"testwriter.txt", "body":"Hello, World!"}`)

	f, _ := kpmFileFromJson(rawFile)

	payload := bytes.NewReader(rawFile)

	req, err := http.NewRequest("POST", "/kpmproxy/write", payload)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	writeRequest(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Response code was %v: want 200", res.Code)
	}


	b, err := ioutil.ReadFile(path.Join(dir, f.Name))
	if err != nil {
		t.Error(err)
	}

	if string(b) != f.Body {
		t.Errorf("File content %s: want %s", string(b), f.Body)
	}

}



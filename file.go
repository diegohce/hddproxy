package main

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

type hddFile struct {
	Dir  string `json:"dir"`
	Name string `json:"name"`
	Body string `json:"body"`
}

func (f *hddFile) write() error {
	fullname := path.Join(f.Dir, f.Name)
	return ioutil.WriteFile(fullname, []byte(f.Body), 0644)
}

func (f *hddFile) writeTo(dir string) error {
	fullname := path.Join(dir, f.Name)
	return ioutil.WriteFile(fullname, []byte(f.Body), 0644)
}

func (f *hddFile) Marshal() ([]byte, error) {
	return json.Marshal(f)
}

func hddFileFromJSON(data []byte) (*hddFile, error) {

	f := &hddFile{}

	err := json.Unmarshal(data, f)

	return f, err
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

type kpmFile struct {
	Dir  string `json:"dir"`
	Name string `json:"name"`
	Body string `json:"body"`
}

func (f *kpmFile) write() error {
	fullname := path.Join(f.Dir, f.Name)
	return ioutil.WriteFile(fullname, []byte(f.Body), 0644)
}

func (f *kpmFile) writeTo(dir string) error {
	fullname := path.Join(dir, f.Name)
	return ioutil.WriteFile(fullname, []byte(f.Body), 0644)
}

func (f *kpmFile) Marshal() ([]byte, error) {
	return json.Marshal(f)
}

func kpmFileFromJson(data []byte) (*kpmFile, error) {

	f := &kpmFile{}

	err := json.Unmarshal(data, f)

	return f, err
}


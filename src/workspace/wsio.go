package workspace

import (
	"encoding/json"
	"io/ioutil"
	"os"
	//    "errors"
	//    "fmt"
)

func (s *Scene) SaveAs(path string) error {
	p := path + s.Name + ".scn"
	//f, err := os.OpenFile(p, os.O_CREATE|os.O_RDWR, 0600)
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	//be careful with defer when using goroutine
	defer f.Close()

	j, err := json.Marshal(s)
	if err != nil {
		return err
	}

	_, err = f.Write(j)
	if err != nil {
		return err
	}

	return nil
}

func Load(path string) (*Scene, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b := make([]byte, 1024)
	b, err = ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var s Scene
	err = json.Unmarshal(b, &s)

	return &s, nil
}

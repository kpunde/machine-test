package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"machine_test/entity"
	"os"
)

type fsHandler struct{}

func (f fsHandler) GetAllFilesFromDir(path string) ([]string, error) {
	fi, err := ioutil.ReadDir(path)
	var fileNames []string
	if err != nil {
		return nil, err
	}

	for _, item := range fi {
		if !item.IsDir() {
			if item.Name()[len(item.Name())-5:] == ".json" {
				fileNames = append(fileNames, item.Name())
			}
		}
	}

	return fileNames, nil
}

func (f fsHandler) GetPortEntityFromFile(path string, data chan entity.PortEntity, errChannel chan error) {
	defer close(data)
	defer close(errChannel)

	fileObj, err := os.Open(path)
	if err != nil {
		errChannel <- err
		return
	}

	r := bufio.NewReader(fileObj)
	dec := json.NewDecoder(r)

	t, err := dec.Token()
	if err != nil {
		errChannel <- err
		return
	}

	if t != json.Delim('{') {
		errChannel <- fmt.Errorf("expected {, got %v", t)
		return
	}

	for dec.More() {
		t, err = dec.Token()
		if err != nil {
			errChannel <- err
		}
		key := t.(string)

		var value entity.Port
		if err = dec.Decode(&value); err != nil {
			errChannel <- err
		}

		var portEntity entity.PortEntity
		portEntity.Name = key
		portEntity.PortObj = value

		data <- portEntity
	}
}

type FsHandlerService interface {
	GetAllFilesFromDir(path string) ([]string, error)
	GetPortEntityFromFile(path string, data chan entity.PortEntity, errChannel chan error)
}

func NewBulkFileHandler() FsHandlerService {
	return &fsHandler{}
}

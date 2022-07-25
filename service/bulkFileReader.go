package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"machine_test/entity"
	"os"
	"path/filepath"
	"strings"
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
				p, _ := filepath.Abs(filepath.Join(path, item.Name()))
				fileNames = append(fileNames, p)
			}
		}
	}

	return fileNames, nil
}

func (f fsHandler) GetPortEntityFromFile(path string, data chan entity.PortEntity, errChannel chan error) {
	log.Println("GetPortEntityFromFile called path: " + path)

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
		key := strings.TrimSpace(t.(string))
		if len(key) == 0 {
			continue
		}

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

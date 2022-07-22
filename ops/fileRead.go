package ops

import (
	"bufio"
	"encoding/json"
	"fmt"
	"machine_test/entity"
	"os"
)

func GetPortEntityFromFile(path string, data chan entity.PortEntity, errChannel chan error) {
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

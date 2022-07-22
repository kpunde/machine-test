package test

import (
	"machine_test/entity"
	"machine_test/ops"
	"testing"
)

func TestGetPortEntityFromFileSmall(t *testing.T) {
	dataChannel := make(chan entity.PortEntity)
	errorChannel := make(chan error)
	go ops.GetPortEntityFromFile("../resources/smallFile.json", dataChannel, errorChannel)

	count := 0

	for _ = range dataChannel {
		count++
	}

	for err := range errorChannel {
		t.Errorf("got error %v\n", err)
	}

	if count == 2 {
		t.Logf("Test executed successfuly")
	}
}

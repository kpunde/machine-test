package service

import (
	"fmt"
	"testing"
)

func TestGetAllFilesFromDir(t *testing.T) {
	iFace := NewBulkFileHandler()
	files, _ := iFace.GetAllFilesFromDir("../resources")
	fmt.Println(files)
}

package handler

import (
	"fmt"
	"log"
	"machine_test/entity"
	"machine_test/service"
	"machine_test/utlities/logger"
)

type portHandler struct {
	logger      logger.LoggerService
	fsService   service.FsHandlerService
	portService service.PostsService
	fsPath      string
}

func (p portHandler) HandleFSFiles() {
	p.logger.Info(fmt.Sprintf("HandleFSFiles Called !\n"))
	filesInDir, err := p.fsService.GetAllFilesFromDir(p.fsPath)
	if err != nil {
		p.logger.Error(err)
		return
	}

	for _, item := range filesInDir {
		p.logger.Debug(item)
		dataChannel := make(chan entity.PortEntity)
		errorChannel := make(chan error)
		go p.fsService.GetPortEntityFromFile(item, dataChannel, errorChannel)

		for {
			select {
			case msg, ok := <-dataChannel:
				if !ok {
					dataChannel = nil
				}
				//log.Println(msg)
				err = p.portService.InsertPost(msg)
				if err != nil {
					log.Println(err)
					p.logger.Error(err)
				}
			case err2, ok := <-errorChannel:
				if !ok {
					errorChannel = nil
				} else {
					log.Println(err2)
					p.logger.Error(err)
				}
			}

			if dataChannel == nil && errorChannel == nil {
				break
			}
		}
	}
}

func (p portHandler) OutputAllEntries() {
	p.logger.Info(fmt.Sprintf("OutputAllEntries Called !\n"))
	db := p.portService.GetAll()
	for key, value := range db {
		log.Println(fmt.Sprintf("PortName %v : PortDetails %v", key, value))
		p.logger.Info(fmt.Sprintf("PortName %v : PortDetails %v", key, value))
	}
}

type PortHandler interface {
	HandleFSFiles()
	OutputAllEntries()
}

func NewPortHandler(fileDir string) PortHandler {
	getLogger, _ := logger.GetLogger()
	fsService := service.NewBulkFileHandler()
	portService := service.NewPostsService()

	return &portHandler{
		logger:      getLogger,
		fsService:   fsService,
		portService: portService,
		fsPath:      fileDir,
	}
}

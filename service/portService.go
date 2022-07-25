package service

import (
	"machine_test/dao"
	"machine_test/entity"
)

type postService struct {
	dao dao.PostsDao
}

func (p postService) GetAll() map[string]entity.Port {
	return p.dao.GetAll()
}

func (p postService) InsertPost(port entity.PortEntity) error {
	return p.dao.InsertPost(port)
}

func (p postService) GetPost(portName string) (*entity.Port, error) {
	return p.dao.GetPost(portName)
}

func (p postService) DeletePost(portName string) error {
	return p.dao.DeletePost(portName)
}

func (p postService) UpdatePost(port entity.PortEntity) error {
	return p.dao.UpdatePost(port)
}

type PostsService interface {
	InsertPost(port entity.PortEntity) error
	GetPost(portName string) (*entity.Port, error)
	DeletePost(portName string) error
	UpdatePost(port entity.PortEntity) error
	GetAll() map[string]entity.Port
}

func NewPostsService() PostsService {
	return &postService{
		dao: dao.NewPostsDao(),
	}
}

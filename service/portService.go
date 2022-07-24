package service

import (
	"machine_test/dao"
	"machine_test/entity"
)

type postService struct {
	dao PostsDao
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

type PostsDao interface {
	InsertPost(port entity.PortEntity) error
	GetPost(portName string) (*entity.Port, error)
	DeletePost(portName string) error
	UpdatePost(port entity.PortEntity) error
}

func NewPostsDao() PostsDao {
	return &postService{
		dao: dao.NewPostsDao(),
	}
}

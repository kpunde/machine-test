package dao

import (
	"fmt"
	"machine_test/entity"
)

var dataBase = make(map[string]entity.Port)

type postDao struct {
	db map[string]entity.Port
}

func (p postDao) InsertPost(port entity.PortEntity) error {
	portName := port.Name
	if _, ok := p.db[portName]; ok {
		return p.UpdatePost(port)
	}

	p.db[portName] = port.PortObj
	return nil
}

func (p postDao) GetPost(portName string) (*entity.Port, error) {
	if val, ok := p.db[portName]; ok {
		return &val, nil
	} else {
		return nil, fmt.Errorf("value not found in datastore")
	}
}

func (p postDao) DeletePost(portName string) error {
	if _, ok := p.db[portName]; ok {
		delete(p.db, portName)
		return nil
	} else {
		return fmt.Errorf("value not found in datastore")
	}
}

func (p postDao) UpdatePost(port entity.PortEntity) error {
	portName := port.Name
	p.db[portName] = port.PortObj
	return nil
}

type PostsDao interface {
	InsertPost(port entity.PortEntity) error
	GetPost(portName string) (*entity.Port, error)
	DeletePost(portName string) error
	UpdatePost(port entity.PortEntity) error
}

func NewPostsDao() PostsDao {
	return &postDao{
		db: dataBase,
	}
}

package dao

import (
	"fmt"
	"machine_test/entity"
	"strings"
)

var dataBase = make(map[string]entity.Port)

type postDao struct {
	db map[string]entity.Port
}

func (p postDao) GetAll() map[string]entity.Port {
	return p.db
}

func (p postDao) InsertPost(port entity.PortEntity) error {
	portName := port.Name
	portName = strings.TrimSpace(portName)
	if len(portName) == 0 {
		return nil
	}
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
	GetAll() map[string]entity.Port
}

func NewPostsDao() PostsDao {
	return &postDao{
		db: dataBase,
	}
}

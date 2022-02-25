package handler

import (
	"github.com/rs/zerolog"
	"meduse-server/pkg/domain/model"
	"net"
	"sync"
)

var (
	connections = &sync.Map{}
)

type Connections struct {
	logger *zerolog.Logger
}

func NewConnections(logger *zerolog.Logger) *Connections {
	return &Connections{
		logger: logger,
	}
}

func (c *Connections) Save(userID string, conn *net.Conn) error {
	_, ok := connections.Load(userID)
	if ok {
		return model.ErrUserConnectionAlreadyExisted
	}
	connections.Store(userID, conn)
	return nil
}

func (c *Connections) Delete(userID string) error {
	_, ok := connections.Load(userID)
	if !ok {
		return model.ErrUserConnectionNotFound
	}
	connections.Delete(userID)
	return nil
}

func (c *Connections) Find(userID string) (*net.Conn, error) {
	conn, ok := connections.Load(userID)
	if !ok {
		return nil, model.ErrUserConnectionNotFound
	}
	v, ok := conn.(*net.Conn)
	if !ok {
		return nil, model.ErrUserConnectionNotFound
	}
	return v, nil
}

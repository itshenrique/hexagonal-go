package repositories

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"nito/api/internals/core/domain"
	"nito/api/internals/core/ports"
	"time"

	"github.com/gomodule/redigo/redis"
)

// DI

type SessionRepository struct {
	conn redis.Conn
}

func NewSessionRepository(address string) ports.ISessionRepository {
	conn, err := redis.Dial(
		"tcp",
		address,
	)

	if err != nil {
		panic("failed to connect redis database")
	}

	return &SessionRepository{
		conn: conn,
	}
}

// Functions

func (r *SessionRepository) Set(key string, value domain.Session, maxAge int) error {
	now := time.Now()
	value.LastModified = &now
	var valueBuffer bytes.Buffer
	enc := gob.NewEncoder(&valueBuffer)
	enc.Encode(value)
	_, err := r.conn.Do("SET", key, []byte(valueBuffer.Bytes()))
	if err != nil {
		return errors.New("error setting session")
	}
	_, err = r.conn.Do("EXPIRE", key, maxAge)
	if err != nil {
		return errors.New("error setting expire session")
	}
	return nil
}

func (r *SessionRepository) Get(key string) (*domain.Session, error) {
	var data []byte
	data, err := redis.Bytes(r.conn.Do("GET", key))
	var valueBuffer bytes.Buffer
	dec := gob.NewDecoder(&valueBuffer)
	dec.Decode(&data)
	result := domain.Session{}
	binary.Read(bytes.NewBuffer(data), binary.BigEndian, result)
	return &result, err
}

func (r *SessionRepository) Delete(key string) error {
	_, err := redis.Bytes(r.conn.Do("DEL", key))
	return err
}

package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidUUID = errors.New("id inválido")
)

type ID string

func NewId() (ID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return ID(id.String()), nil
}

func ParseID(id string) (ID, error) {
	_id, err := uuid.Parse(id)
	if err != nil {
		return "", errors.Join(err, ErrInvalidUUID)
	}

	return ID(_id.String()), nil
}

func (i ID) String() string {
	return string(i)
}

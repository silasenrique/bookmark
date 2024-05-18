package service

import (
	"errors"
	command "go-read-list/src/business/bookmark/application/command"
	"go-read-list/src/business/bookmark/domain/entity"
)

func (l *BookmarkService) GetById(id string) (*command.ReadResponse, error) {
	if id == "" {
		return nil, errors.New("o id não pode ser nulo")
	}

	_id, err := entity.ParseID(id)
	if err != nil {
		return nil, err
	}

	read, err := l.rep.GetById(_id)
	if err != nil {
		return nil, err
	}

	return command.ToReadResponde(read), nil
}

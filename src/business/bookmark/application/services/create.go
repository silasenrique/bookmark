package service

import (
	"errors"
	"go-read-list/src/business/bookmark/application/command"
	"go-read-list/src/business/bookmark/domain/entity"
)

func (l *BookmarkService) Create(cmd *command.BookmarkCreateCommand) (*command.ReadResponse, error) {
	if cmd.Url == "" {
		return nil, errors.New("precisa ser informado uma url")
	}

	if cmd.Title == "" {
		return nil, errors.New("é necessário informar um titulo")
	}

	id, err := entity.NewId()
	if err != nil {
		return nil, err
	}

	read := entity.NewRead(id, cmd.Url, cmd.Title)
	read.AddCover(cmd.Cover)

	if cmd.Path != "" {
		parseId, err := entity.ParseID(cmd.Path)
		if err != nil {
			return nil, err
		}

		read.AddPath(parseId)
	}

	err = l.rep.Create(read)

	return command.ToReadResponde(read), err
}

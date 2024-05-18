package service

import (
	"errors"
	"go-read-list/src/business/collection/application/command"
	"go-read-list/src/business/collection/application/mapper"
	"time"
)

func (f *CollectionService) Rename(cmd *command.CollectionRenameCommand) (*command.CollectionResponse, error) {
	if cmd.Id == 0 {
		return nil, errors.New("o id deve ser informado")
	}

	if cmd.Name == "" {
		return nil, errors.New("o nome deve ser informado")
	}

	colle, err := f.rep.GetFolderById(cmd.Id)
	if err != nil {
		return nil, err
	}

	colle.AddName(cmd.Name)
	colle.AddUpdateAt(time.Now().Unix())

	err = f.rep.Update(colle)
	if err != nil {
		return nil, err
	}

	return mapper.CollectionToResponse(colle), nil
}

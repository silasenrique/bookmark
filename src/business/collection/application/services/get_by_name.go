package service

import (
	"errors"
	"go-read-list/src/business/collection/application/command"
	"go-read-list/src/business/collection/application/mapper"
)

func (f *CollectionService) GetByName(name string) (*command.CollectionResponse, error) {
	if name == "" {
		return nil, errors.New("o nome deve ser informado")
	}

	colle, err := f.rep.GetByName(name)
	if err != nil {
		return nil, err
	}

	if colle.GetParentID() != 0 {
		parent, err := f.rep.GetById(colle.GetParentID())
		if err != nil {
			return nil, err
		}

		colle.AddParent(parent)
	}

	return mapper.CollectionToResponse(colle), nil
}

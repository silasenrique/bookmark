package service

import (
	"errors"
	"go-read-list/src/business/collection/application/command"
	"go-read-list/src/business/collection/application/mapper"
)

func (f *CollectionService) GetById(id int64) (*command.CollectionResponse, error) {
	if id == 0 {
		return nil, errors.New("id deve ser informado")
	}

	colle, err := f.rep.GetFolderByInternalId(id)
	if err != nil {
		return nil, err
	}

	if colle.GetInternalParentID() != 0 {
		parent, err := f.rep.GetFolderByInternalId(colle.GetInternalParentID())
		if err != nil {
			return nil, err
		}

		colle.AddParent(parent)
	}

	return mapper.CollectionToResponse(colle), nil
}

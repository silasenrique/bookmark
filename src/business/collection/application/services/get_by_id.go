package service

import (
	"database/sql"
	"errors"
	"go-read-list/src/business/collection/application/command"
	"go-read-list/src/business/collection/application/mapper"
)

func (f *CollectionService) GetById(id int64) (*command.CollectionResponse, error) {
	if id == 0 {
		return nil, ErrIdEqualZero
	}

	colle, err := f.rep.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Join(err, ErrNotExist)
		}

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

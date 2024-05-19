package service

import (
	"errors"
	"go-read-list/src/business/collection/application/command"
	"go-read-list/src/business/collection/application/mapper"
	"time"
)

func (f *CollectionService) Update(colle *command.CollectionUpdateCommand) (*command.CollectionResponse, error) {
	if colle.Id == 0 {
		return nil, errors.New("e necessário informar um identificador da colecao")
	}

	_colle, err := f.rep.GetById(colle.Id)
	if err != nil {
		return nil, err
	}

	err = f.parseParent(_colle, _colle.GetInternalParentID())
	if err != nil {
		return nil, err
	}

	if colle.Name != "" {
		_colle.AddName(colle.Name)
	}

	if colle.ParentId != 0 && !colle.RemoveParent {
		err = f.parseParent(_colle, colle.ParentId)
		if err != nil {
			return nil, err
		}
	} else if colle.RemoveParent {
		_colle.RemoveParent()
	}

	_colle.AddUpdateAt(time.Now().Unix())

	err = f.rep.Update(_colle)
	if err != nil {
		return nil, err
	}

	return mapper.CollectionToResponse(_colle), nil
}

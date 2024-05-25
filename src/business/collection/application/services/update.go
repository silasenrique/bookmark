package service

import (
	"database/sql"
	"errors"
	"go-read-list/src/business/collection/application/command"
	"go-read-list/src/business/collection/application/mapper"
	"strings"
	"time"
)

func (f *CollectionService) Update(colle *command.CollectionUpdateCommand) (*command.CollectionResponse, error) {
	if colle.Id == 0 {
		return nil, ErrIdEqualZero
	}

	_colle, err := f.rep.GetById(colle.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Join(err, ErrNotExist)
		}

		return nil, err
	}

	if strings.Trim(colle.Name, " ") != "" {
		_colle.AddName(colle.Name)
	}

	if strings.Trim(colle.Icon, " ") != "" && colle.Icon != _colle.GetIcon() {
		_colle.AddIcon(colle.Icon)
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

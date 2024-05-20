package service

import (
	"database/sql"
	"errors"
	"go-read-list/src/business/collection/application/command"
	"go-read-list/src/business/collection/application/mapper"
	"strings"
	"time"
)

func (f *CollectionService) Rename(cmd *command.CollectionRenameCommand) (*command.CollectionResponse, error) {
	if cmd.Id == 0 {
		return nil, ErrIdEqualZero
	}

	if strings.Trim(cmd.Name, " ") == "" {
		return nil, ErrNameIsNull
	}

	colle, err := f.rep.GetById(cmd.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExist
		}

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

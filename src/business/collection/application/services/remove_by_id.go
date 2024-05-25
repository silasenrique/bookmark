package service

import (
	"database/sql"
	"errors"
)

func (f *CollectionService) DeleteById(id int64, recusive bool) error {
	if id == 0 {
		return ErrIdEqualZero
	}

	_, err := f.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotExist
		}

		return err
	}

	count, err := f.rep.CountDependencies(id)
	if err != nil {
		return err
	}

	if count > 0 && !recusive {
		return ErrCollectionHasChildrens
	}

	return f.rep.RemevoFolderById(id)
}

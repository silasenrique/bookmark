package service

import (
	"database/sql"
	"errors"
	"go-read-list/src/business/collection/application/mapper"
)

func (f *CollectionService) GetByParentId(id int64) (*mapper.CollectionsResponse, error) {
	if id == 0 {
		return nil, ErrIdEqualZero
	}

	colle, err := f.rep.GetByParentId(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExist
		}

		return nil, err
	}

	return mapper.ToCollectionsResponse(colle), nil
}

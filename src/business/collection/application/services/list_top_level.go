package service

import (
	"database/sql"
	"errors"
	"go-read-list/src/business/collection/application/mapper"
)

func (f *CollectionService) ListAllTopCollections() (*mapper.CollectionsResponse, error) {
	colle, err := f.rep.GetTopFolder()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExist
		}

		return nil, err
	}

	return mapper.ToCollectionsResponse(colle), nil
}

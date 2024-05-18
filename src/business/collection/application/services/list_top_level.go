package service

import (
	"go-read-list/src/business/collection/application/mapper"
)

func (f *CollectionService) ListAllTopCollections() (*mapper.CollectionsResponse, error) {
	colle, err := f.rep.GetTopFolder()
	if err != nil {
		return nil, err
	}

	return mapper.ToCollectionsResponse(colle), nil
}

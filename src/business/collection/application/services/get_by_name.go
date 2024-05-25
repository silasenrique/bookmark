package service

import (
	"go-read-list/src/business/collection/application/mapper"
	"strings"
)

// Search a collection by its name.
// If there are other collections with the same name, they will be returned as well.
func (f *CollectionService) GetByName(name string) (*mapper.CollectionsResponse, error) {
	if strings.Trim(name, " ") == "" {
		return nil, ErrNameIsNull
	}

	colle, err := f.rep.GetByName(name)
	if err != nil {
		return nil, err
	}

	return mapper.ToCollectionsResponse(colle), nil
}

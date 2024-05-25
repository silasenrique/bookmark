package service

import "go-read-list/src/business/collection/application/mapper"

func (s *CollectionService) List() (*mapper.CollectionsResponse, error) {
	coll, err := s.rep.List()
	if err != nil {
		return nil, err
	}

	return mapper.ToCollectionsResponse(coll), err
}

package service

import (
	"go-read-list/src/business/collection/domain/entity"
	"go-read-list/src/business/collection/domain/repository"
)

type CollectionService struct {
	rep repository.Collection
}

func NewFolderService(rep repository.Collection) *CollectionService {
	return &CollectionService{rep}
}

func (f *CollectionService) parseParent(dir *entity.Collection, parentId int64) error {
	if parentId == 0 {
		return nil
	}

	parent, err := f.rep.GetById(parentId)
	if err != nil {
		return err
	}

	dir.AddParent(parent)

	return nil
}

package service

import (
	"go-read-list/src/business/collection/application/command"
	"go-read-list/src/business/collection/application/mapper"
	"go-read-list/src/business/collection/domain/entity"
)

func (f *CollectionService) Create(colle *command.CollectionCreateCommand) (*command.CollectionResponse, error) {
	id, err := f.rep.GetNextId()
	if err != nil {
		return nil, err
	}

	dir := entity.NewCollection(id, colle.Name).AddIcon(colle.Icon)

	err = f.parseParent(dir, colle.ParentId)
	if err != nil {
		return nil, err
	}

	err = f.rep.CreateFolder(dir)
	if err != nil {
		return nil, err
	}

	return mapper.CollectionToResponse(dir), nil
}

package infrastructure

import (
	"go-read-list/src/business/collection/application/mapper"
	"go-read-list/src/business/collection/domain/entity"
)

func (f *collectionPersistence) GetFolderById(id int64) (*entity.Collection, error) {
	colle := &mapper.CollectionDto{}

	err := f.QueryRow("select id, name, ifnull(icon, ''), ifnull(collection_id, 0), create_at, update_at from collection where id = ?", id).
		Scan(
			&colle.ID,
			&colle.Name,
			&colle.Icon,
			&colle.IntParentId,
			&colle.CreateAt,
			&colle.UpdateAt,
		)

	if err != nil {
		return nil, err
	}

	return mapper.DtoToCollection(colle), nil
}

func (f *collectionPersistence) GetFolderByInternalId(id int64) (*entity.Collection, error) {
	colle := &mapper.CollectionDto{}

	err := f.QueryRow("select id, name, ifnull(icon, ''), ifnull(collection_id, 0), create_at, update_at from collection where id = ?", id).
		Scan(
			&colle.ID,
			&colle.Name,
			&colle.Icon,
			&colle.IntParentId,
			&colle.CreateAt,
			&colle.UpdateAt,
		)

	if err != nil {
		return nil, err
	}

	return mapper.DtoToCollection(colle), nil
}

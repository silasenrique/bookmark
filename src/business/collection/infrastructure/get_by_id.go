package infrastructure

import (
	"go-read-list/src/business/collection/application/mapper"
	"go-read-list/src/business/collection/domain/entity"
)

func (f *collectionPersistence) GetById(id int64) (*entity.Collection, error) {
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

func (f *collectionPersistence) GetByParentId(id int64) ([]*entity.Collection, error) {
	collections := []*entity.Collection{}

	rows, err := f.Query("select id, name, ifnull(icon, ''), ifnull(collection_id, 0), create_at, update_at from collection where collection_id = ?", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		colle := &mapper.CollectionDto{}

		err = rows.Scan(
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

		collections = append(collections, mapper.DtoToCollection(colle))
	}

	return collections, nil
}

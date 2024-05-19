package infrastructure

import (
	"go-read-list/src/business/collection/application/mapper"
	"go-read-list/src/business/collection/domain/entity"
)

func (f *collectionPersistence) GetByName(name string) ([]*entity.Collection, error) {
	colle := []*entity.Collection{}

	rows, err := f.Query("select id, name, ifnull(icon, ''), ifnull(collection_id, 0), create_at, update_at from collection where name = ?", name)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		col := &mapper.CollectionDto{}
		err = rows.Scan(
			&col.ID,
			&col.Name,
			&col.Icon,
			&col.IntParentId,
			&col.CreateAt,
			&col.UpdateAt,
		)

		if err != nil {
			return nil, err
		}

		colle = append(colle, mapper.DtoToCollection(col))
	}

	return colle, nil
}

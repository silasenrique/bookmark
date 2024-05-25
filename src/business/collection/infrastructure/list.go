package infrastructure

import (
	"go-read-list/src/business/collection/application/mapper"
	"go-read-list/src/business/collection/domain/entity"
)

func (f *collectionPersistence) List() ([]*entity.Collection, error) {
	colls := []*entity.Collection{}

	rows, err := f.Query("select id, name, ifnull(icon, ''), ifnull(collection_id, 0), create_at, update_at from collection")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

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

		colls = append(colls, mapper.DtoToCollection(col))
	}

	return colls, nil
}

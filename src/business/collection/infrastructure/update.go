package infrastructure

import (
	"database/sql"
	"go-read-list/src/business/collection/domain/entity"
)

func (f *collectionPersistence) Update(colle *entity.Collection) error {
	ist, err := f.Prepare(" update collection set name = ?, icon = ?, collection_id = ?, update_at = ? where id = ?")
	if err != nil {
		return err
	}

	defer ist.Close()

	parent := sql.NullInt64{Int64: colle.GetInternalParentID()}

	_, err = ist.Exec(
		colle.GetName(),
		colle.GetIcon(),
		parent,
		colle.GetUpdateAt(),
		colle.GetID(),
	)

	return err
}

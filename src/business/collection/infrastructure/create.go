package infrastructure

import (
	"database/sql"
	"go-read-list/src/business/collection/domain/entity"
)

func (f *collectionPersistence) CreateFolder(colle *entity.Collection) error {
	ist, err := f.Prepare("insert into collection (id, name, icon, collection_id, create_at, update_at) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer ist.Close()

	parent := sql.NullInt64{Int64: colle.GetParentID()}

	_, err = ist.Exec(
		colle.GetID(),
		colle.GetName(),
		colle.GetIcon(),
		parent,
		colle.GetUnixCreateAt(),
		colle.GetUnixUpdateAt(),
	)

	return err
}

package infrastructure

import (
	"go-read-list/src/business/collection/domain/entity"
)

func (f *collectionPersistence) CreateFolder(colle *entity.Collection) error {
	ist, err := f.Prepare("insert into collection (id, name, icon, collection_id, create_at, update_at) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer ist.Close()

	// var parent interface{}
	// if colle.GetParentID() != 0 {
	// 	parent = colle.GetParentID()
	// }

	_, err = ist.Exec(
		colle.GetID(),
		colle.GetName(),
		colle.GetIcon(),
		colle.GetParentID(),
		colle.GetUnixCreateAt(),
		colle.GetUnixUpdateAt(),
	)

	return err
}

package infrastructure

import "go-read-list/src/business/bookmark/domain/entity"

func (r *bookmarkPersistence) DeleteById(id entity.ID) error {
	sel, err := r.Prepare("delete from read where _id = ?")
	if err != nil {
		return err
	}

	defer sel.Close()

	_, err = sel.Exec(id.String())

	return err
}

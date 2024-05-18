package infrastructure

import "go-read-list/src/business/bookmark/domain/entity"

func (r *bookmarkPersistence) GetByPath(id entity.ID) ([]*entity.Bookmark, error) {
	read := []*entity.Bookmark{}

	sel, err := r.Prepare("select _id, url, title, cover, create_at, update_at from read where path = ?")
	if err != nil {
		return nil, err
	}

	defer sel.Close()

	rows, err := sel.Query(id.String())
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		dto := &entity.BookmarkDTO{}

		err = rows.Scan(
			&dto.Id,
			&dto.Url,
			&dto.Title,
			&dto.Cover,
			&dto.CreateAt,
			&dto.UpdateAt,
		)

		if err != nil {
			return nil, err
		}

		read = append(read, entity.DtoToReader(dto))
	}

	return read, nil
}

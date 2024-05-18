package infrastructure

import "go-read-list/src/business/bookmark/domain/entity"

func (r *bookmarkPersistence) GetById(id entity.ID) (*entity.Bookmark, error) {
	dto := &entity.BookmarkDTO{}

	sel, err := r.Prepare("select _id, url, title, cover, create_at, update_at from read where _id = ?")
	if err != nil {
		return nil, err
	}

	err = sel.QueryRow(id.String()).Scan(
		&dto.Id,
		&dto.Url,
		&dto.Title,
		&dto.Cover,
		&dto.CreateAt,
		&dto.UpdateAt,
	)

	return entity.DtoToReader(dto), err
}

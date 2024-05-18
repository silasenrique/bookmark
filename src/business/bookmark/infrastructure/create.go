package infrastructure

import "go-read-list/src/business/bookmark/domain/entity"

func (r *bookmarkPersistence) Create(read *entity.Bookmark) error {
	ist, err := r.Prepare("insert into read (_id, url, title, cover, create_at, update_at) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer ist.Close()

	_, err = ist.Exec(
		read.GetId(),
		read.GetUrl(),
		read.GetTitle(),
		read.GetCover(),
		read.GetCreateAt(),
		read.GetUpdateAt(),
	)

	return err
}

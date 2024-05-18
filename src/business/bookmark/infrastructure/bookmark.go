package infrastructure

import (
	"database/sql"
	"go-read-list/src/business/bookmark/domain/repository"
)

type bookmarkPersistence struct {
	*sql.DB
}

func NewBookmarkRepository(db *sql.DB) repository.BookmarkRepository {
	return &bookmarkPersistence{db}
}

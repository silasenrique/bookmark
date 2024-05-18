package cli

import (
	"database/sql"
	books "go-read-list/src/business/bookmark/application/services"
	booki "go-read-list/src/business/bookmark/infrastructure"
	colls "go-read-list/src/business/collection/application/services"
	colln "go-read-list/src/business/collection/infrastructure"
)

type wrapper struct {
	collection *colls.CollectionService
	bookmark   *books.BookmarkService
}

func NewWrapperRepository(db *sql.DB) *wrapper {
	return &wrapper{
		collection: colls.NewFolderService(colln.NewCollectionRepository(db)),
		bookmark:   books.NewBookmarkService(booki.NewBookmarkRepository(db)),
	}
}

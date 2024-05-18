package repository

import "go-read-list/src/business/bookmark/domain/entity"

type BookmarkWriter interface {
	Create(read *entity.Bookmark) error
	DeleteById(id entity.ID) error
}

type BookmarkReader interface {
	GetByPath(id entity.ID) ([]*entity.Bookmark, error)
	GetById(id entity.ID) (*entity.Bookmark, error)
}

type BookmarkRepository interface {
	BookmarkWriter
	BookmarkReader
}

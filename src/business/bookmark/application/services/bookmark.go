package service

import (
	"go-read-list/src/business/bookmark/domain/repository"
)

type BookmarkService struct {
	rep repository.BookmarkRepository
}

func NewBookmarkService(rep repository.BookmarkRepository) *BookmarkService {
	return &BookmarkService{rep}
}

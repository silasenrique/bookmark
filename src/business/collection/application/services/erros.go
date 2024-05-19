package service

import "errors"

var (
	ErrParentNotFound = errors.New("collecttion parent not exist")
	ErrNotExist       = errors.New("collection not exist")
)

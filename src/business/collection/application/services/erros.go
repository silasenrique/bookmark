package service

import "errors"

var (
	ErrParentNotFound = errors.New("collecttion parent not exist")
	ErrNotExist       = errors.New("collection not exist")
	ErrIdEqualZero    = errors.New("id não pode ser menor ou igual que zero")
	ErrNameIsNull     = errors.New("o nome não pode ser nulo ou vazio")
)

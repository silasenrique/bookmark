package repository

import "go-read-list/src/business/collection/domain/entity"

type Collection interface {
	CreateFolder(folder *entity.Collection) error
	RemevoFolderById(id int64) error
	Update(colle *entity.Collection) error
	GetById(id int64) (*entity.Collection, error)
	GetTopFolder() ([]*entity.Collection, error)
	GetByName(name string) ([]*entity.Collection, error)
	CountDependencies(id int64) (int64, error)
	GetNextId() (int64, error)
	List() ([]*entity.Collection, error)
	GetByParentId(id int64) ([]*entity.Collection, error)
}

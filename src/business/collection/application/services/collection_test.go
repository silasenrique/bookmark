package service_test

import (
	"go-read-list/src/business/collection/domain/entity"
)

var seq int64 = 121
var mockTable = map[int64]*entity.Collection{}

type collectionMock struct {
	CreateFunc            func(folder *entity.Collection) error
	RemevoByIdFunc        func(id int64) error
	UpdateFunc            func(colle *entity.Collection) error
	GetByIdFunc           func(id int64) (*entity.Collection, error)
	GetTopFunc            func() ([]*entity.Collection, error)
	GetByNameFunc         func(name string) ([]*entity.Collection, error)
	CountDependenciesFunc func(id int64) (int64, error)
	GetNextIdFunc         func() (int64, error)
	ListFunc              func() ([]*entity.Collection, error)
}

func (f collectionMock) CountDependencies(id int64) (int64, error) {
	return f.CountDependenciesFunc(id)
}

func (f collectionMock) CreateFolder(folder *entity.Collection) error {
	return f.CreateFunc(folder)
}

func (f collectionMock) GetByName(name string) ([]*entity.Collection, error) {
	return f.GetByNameFunc(name)
}

func (f collectionMock) GetById(id int64) (*entity.Collection, error) {
	return f.GetByIdFunc(id)
}

func (f collectionMock) GetNextId() (int64, error) {
	return f.GetNextIdFunc()
}

func (f collectionMock) GetTopFolder() ([]*entity.Collection, error) {
	return f.GetTopFunc()
}

func (f collectionMock) RemevoFolderById(id int64) error {
	return f.RemevoByIdFunc(id)
}

func (f collectionMock) Update(colle *entity.Collection) error {
	return f.UpdateFunc(colle)
}

func (f collectionMock) List() ([]*entity.Collection, error) {
	return f.ListFunc()
}

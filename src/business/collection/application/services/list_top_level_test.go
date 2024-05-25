package service_test

import (
	"database/sql"
	service "go-read-list/src/business/collection/application/services"
	"go-read-list/src/business/collection/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCollectionWithoutParent(t *testing.T) {
	collections := map[int64]*entity.Collection{
		1: entity.NewCollection(1, "Books"),
		2: entity.NewCollection(2, "Bookstore").AddParent(entity.NewCollection(1, "Books")),
		3: entity.NewCollection(3, "Mangas").AddParent(entity.NewCollection(1, "Books")),
		4: entity.NewCollection(4, "Movies"),
	}

	mock := &collectionMock{
		GetTopFunc: func() ([]*entity.Collection, error) {
			coll := []*entity.Collection{}

			if len(collections) == 0 {
				return nil, sql.ErrNoRows
			}

			for _, c := range collections {
				if c.GetParentID() != nil {
					continue
				}

				coll = append(coll, c)
			}

			return coll, nil
		},
	}

	serv := service.NewFolderService(mock)
	resp, _ := serv.ListAllTopCollections()

	assert.Equal(t, 2, len(*resp), "não foi retornado a quantidade correta de collections sem pai")
}

func TestGetZeroCollectionWithoutParent(t *testing.T) {
	collections := map[int64]*entity.Collection{}

	mock := &collectionMock{
		GetTopFunc: func() ([]*entity.Collection, error) {
			coll := []*entity.Collection{}

			if len(collections) == 0 {
				return coll, sql.ErrNoRows
			}

			for _, c := range collections {
				if c.GetParentID() == nil {
					continue
				}

				coll = append(coll, c)
			}

			return coll, nil
		},
	}

	serv := service.NewFolderService(mock)
	_, err := serv.ListAllTopCollections()
	if err == nil {
		t.Errorf("deverira ter retornado um erro")
		t.FailNow()
	}

	assert.ErrorIs(t, err, service.ErrNotExist, "deve retornar o erro de nome nao pode ser vazio")
}

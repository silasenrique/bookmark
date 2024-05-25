package service_test

import (
	"database/sql"
	service "go-read-list/src/business/collection/application/services"
	"go-read-list/src/business/collection/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveExistingId(t *testing.T) {
	collections := map[int64]*entity.Collection{
		1: entity.NewCollection(1, "Movies"),
		2: entity.NewCollection(2, "Books"),
	}

	mock := &collectionMock{
		RemevoByIdFunc: func(id int64) error {
			delete(collections, id)
			return nil
		},
		GetByIdFunc: func(id int64) (*entity.Collection, error) {
			collection := collections[id]
			if collection == nil {
				return nil, sql.ErrNoRows
			}

			return collection, nil
		},
		CountDependenciesFunc: func(id int64) (int64, error) {
			return 0, nil
		},
	}

	serv := service.NewFolderService(mock)

	err := serv.DeleteById(2, false)
	if err != nil {
		t.Errorf("erro nao previsto retornado. Err: %s", err)
		t.FailNow()
	}

	assert.Equal(t, 1, len(collections), "a quantidade esperada nao bate")
}

func TestRemoveExistingIdWithDependencies(t *testing.T) {
	collections := map[int64]*entity.Collection{
		1: entity.NewCollection(1, "Movies"),
		2: entity.NewCollection(2, "Books"),
		3: entity.NewCollection(3, "Fiction").AddParent(entity.NewCollection(2, "Books")),
		4: entity.NewCollection(4, "Fantasy").AddParent(entity.NewCollection(2, "Books")),
		5: entity.NewCollection(5, "Work"),
		6: entity.NewCollection(6, "Read Later").AddParent(entity.NewCollection(5, "Work")),
	}

	mock := &collectionMock{
		RemevoByIdFunc: func(id int64) error {
			for _, c := range collections {
				_id := c.GetParentID()
				if _id == nil {
					continue
				}

				if *_id == id {
					delete(collections, c.GetID())
				}
			}

			delete(collections, id)
			return nil
		},
		GetByIdFunc: func(id int64) (*entity.Collection, error) {
			collection := collections[id]
			if collection == nil {
				return nil, sql.ErrNoRows
			}

			return collection, nil
		},
		CountDependenciesFunc: func(id int64) (int64, error) {
			var count int64 = 0

			for _, c := range collections {
				_id := c.GetParentID()
				if _id == nil {
					continue
				}

				if *_id == id {
					count++
				}
			}

			return count, nil
		},
	}

	serv := service.NewFolderService(mock)

	t.Run("recursive", func(t *testing.T) {
		err := serv.DeleteById(2, true)
		if err != nil {
			t.Errorf("erro nao previsto retornado. Err: %s", err)
			t.FailNow()
		}

		assert.Equal(t, 3, len(collections), "a quantidade esperada nao bate")
	})

	t.Run("not recursive", func(t *testing.T) {
		err := serv.DeleteById(5, false)
		if err == nil {
			t.Errorf("deveria ser retornado um erro.")
			t.FailNow()
		}

		assert.ErrorIs(t, service.ErrCollectionHasChildrens, err, "os erros não batem")
	})

}

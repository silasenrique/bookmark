package service_test

import (
	"database/sql"
	"go-read-list/src/business/collection/application/command"
	service "go-read-list/src/business/collection/application/services"
	"go-read-list/src/business/collection/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionUpdateErrors(t *testing.T) {
	collections := map[int64]*entity.Collection{}

	mock := &collectionMock{
		UpdateFunc: func(colle *entity.Collection) error {
			collections[colle.GetID()] = colle
			return nil
		},
		GetByIdFunc: func(id int64) (*entity.Collection, error) {
			collection := collections[id]
			if collection == nil {
				return nil, sql.ErrNoRows
			}

			return collection, nil
		},
	}

	serv := service.NewFolderService(mock)

	t.Run("invalid id", func(t *testing.T) {
		_, err := serv.Update(&command.CollectionUpdateCommand{Id: 0})
		if err == nil {
			t.Errorf("deveria retornar um erro")
			t.FailNow()
		}

		assert.ErrorIs(t, service.ErrIdEqualZero, err, "o erro retornando não era o esperado")
	})

	t.Run("not found collection", func(t *testing.T) {
		_, err := serv.Update(&command.CollectionUpdateCommand{Id: 123})
		if err == nil {
			t.Errorf("deveria retornar um erro")
			t.FailNow()
		}

		assert.ErrorIsf(t, err, service.ErrNotExist, "erro não esperado retornando")
	})
}

func TestCollectionUpdateName(t *testing.T) {
	collections := map[int64]*entity.Collection{
		1: entity.NewCollection(1, "movie").AddIcon("path_to_icon"),
		2: entity.NewCollection(2, "mangas").AddParent(
			entity.NewCollection(1, "movie").AddIcon("path_to_icon"),
		).AddIcon("url_path_to_icon"),
	}

	mock := &collectionMock{
		UpdateFunc: func(colle *entity.Collection) error {
			collections[colle.GetID()] = colle
			return nil
		},
		GetByIdFunc: func(id int64) (*entity.Collection, error) {
			collection := collections[id]
			if collection == nil {
				return nil, sql.ErrNoRows
			}

			return collection, nil
		},
	}

	serv := service.NewFolderService(mock)

	cmd := &command.CollectionUpdateCommand{Id: 1, Name: "watch later"}
	resp, err := serv.Update(cmd)
	if err != nil {
		t.Errorf("deveria retornar um erro")
		t.FailNow()
	}

	assert.Equal(t, cmd.Name, resp.Name, "os nomes deveriam ser iguais")
}

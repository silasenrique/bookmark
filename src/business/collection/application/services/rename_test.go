package service_test

import (
	"database/sql"
	"go-read-list/src/business/collection/application/command"
	service "go-read-list/src/business/collection/application/services"
	"go-read-list/src/business/collection/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRename(t *testing.T) {
	collections := map[int64]*entity.Collection{
		1: entity.NewCollection(1, "Movies").AddCreateAt(1716201655),
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
	t.Run("existing collection", func(t *testing.T) {
		cmd := &command.CollectionRenameCommand{Id: 1, Name: "Filmes"}
		resp, err := serv.Rename(cmd)
		if err != nil {
			t.Errorf("erro nao previsto retornado. Err: %s", err)
			t.FailNow()
		}

		assert.Equal(t, cmd.Name, resp.Name, "os nomes não batem")
		assert.NotEqual(t, resp.CreateAt, resp.UpdateAt, "os valores não deviriam ser iguais")
	})

	t.Run("existing collection without name", func(t *testing.T) {
		cmd := &command.CollectionRenameCommand{Id: 1, Name: ""}
		_, err := serv.Rename(cmd)
		if err == nil {
			t.Errorf("deveria retornar um erro")
			t.FailNow()
		}

		assert.ErrorIs(t, service.ErrNameIsNull, err, "os erros não batem")
	})

	t.Run("existing collection without name", func(t *testing.T) {
		cmd := &command.CollectionRenameCommand{Id: 2, Name: "Filmes"}
		_, err := serv.Rename(cmd)
		if err == nil {
			t.Errorf("deveria retornar um erro")
			t.FailNow()
		}

		assert.ErrorIs(t, service.ErrNotExist, err, "os erros não batem")
	})
}

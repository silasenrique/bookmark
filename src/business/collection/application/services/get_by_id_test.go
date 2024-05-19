package service_test

import (
	"database/sql"
	"errors"
	service "go-read-list/src/business/collection/application/services"
	"go-read-list/src/business/collection/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetById(t *testing.T) {
	mock := &collectionMock{
		GetByIdFunc: func(id int64) (*entity.Collection, error) {
			collection := mockTable[id]
			if collection == nil {
				return nil, sql.ErrNoRows
			}

			return collection, nil
		},
	}

	mockTable[130] = entity.NewCollection(130, "Read Later")

	serv := service.NewFolderService(mock)

	t.Run("not found", func(t *testing.T) {
		_, err := serv.GetById(100)
		if err == nil {
			t.Errorf("deve retornar um erro, collection nao deve existir! Err: %s", err)
			t.FailNow()
		}

		if errors.Is(err, service.ErrNotExist) {
			return
		}

		t.Errorf("não foi retornado o erro esperado! Err: %s", err)
		t.FailNow()
	})

	t.Run("found", func(t *testing.T) {
		resp, err := serv.GetById(130)
		if err != nil {
			t.Errorf("erro não previsto! Err: %s", err)
			t.FailNow()
		}

		collection := mockTable[130]

		assert.Equal(t, collection.GetID(), resp.Id, "os ids retornados não batem")
		assert.Equal(t, collection.GetName(), resp.Name, "os nomes não batem")
		assert.Equal(t, collection.GetStringCreateAt(), resp.CreateAt, "As datas de criação não batem")
		assert.Equal(t, collection.GetStringUpdateAt(), resp.UpdateAt, "As datas de atualização não batem")
		assert.Equal(t, collection.GetIcon(), resp.Icon, "os icones não batem")
	})
}

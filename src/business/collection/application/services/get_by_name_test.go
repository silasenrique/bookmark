package service_test

import (
	service "go-read-list/src/business/collection/application/services"
	"go-read-list/src/business/collection/domain/entity"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByName(t *testing.T) {
	collections := map[int64]*entity.Collection{
		1: entity.NewCollection(1, "Books"),
		2: entity.NewCollection(2, "Bookstore"),
		3: entity.NewCollection(3, "Mangas"),
	}

	mock := &collectionMock{
		GetByNameFunc: func(name string) ([]*entity.Collection, error) {
			coll := []*entity.Collection{}

			for _, c := range collections {
				if strings.Contains(strings.ToLower(c.GetName()), strings.ToLower(name)) {
					coll = append(coll, c)
				}
			}

			return coll, nil
		},
	}

	serv := service.NewFolderService(mock)

	t.Run("corret name", func(t *testing.T) {
		resp, err := serv.GetByName("book")
		if err != nil {
			t.Errorf("erro não previsto! Err: %s", err)
			t.FailNow()
		}

		assert.Equal(t, 2, len(*resp), "deveria retornar 2 registros")
	})

	t.Run("without name", func(t *testing.T) {
		_, err := serv.GetByName("")
		if err == nil {
			t.Errorf("deverira ter retornado um erro")
			t.FailNow()
		}

		assert.ErrorIs(t, err, service.ErrNameIsNull, "deve retornar o erro de nome nao pode ser vazio")
	})
}

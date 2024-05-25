package service_test

import (
	"database/sql"
	"errors"
	"go-read-list/src/business/collection/application/command"
	service "go-read-list/src/business/collection/application/services"
	"go-read-list/src/business/collection/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCollectionWithoutParent(t *testing.T) {
	tt := &collectionMock{
		CreateFunc: func(folder *entity.Collection) error {
			mockTable[folder.GetID()] = folder
			return nil
		},
		GetNextIdFunc: func() (int64, error) {
			seq = seq + 1
			return seq, nil
		},
	}

	serv := service.NewFolderService(tt)

	t.Run("cool collections", func(t *testing.T) {
		commands := []struct {
			cmd *command.CollectionCreateCommand
		}{
			{&command.CollectionCreateCommand{Name: "Move", Icon: "path_to_icon"}},
			{&command.CollectionCreateCommand{Name: "Golang"}},
			{&command.CollectionCreateCommand{Name: "  .  "}},
		}

		for _, i := range commands {
			resp, err := serv.Create(i.cmd)
			if err != nil {
				t.Errorf("não foi retornado o erro esperado! Err: %s", err)
				t.FailNow()
			}

			assert.Equal(t, i.cmd.Name, resp.Name, "O nome não retornou como esperado")
			assert.Equal(t, i.cmd.Icon, resp.Icon, "O icone não retornou como esperado")
			assert.NotEmpty(t, resp.CreateAt, "A data do cadastro do registro deveria estar preenchida")
			assert.NotEmpty(t, resp.UpdateAt, "A data de alteração deveria estar preenchida")
		}
	})

	t.Run("empty name", func(t *testing.T) {
		commands := []struct {
			cmd *command.CollectionCreateCommand
		}{
			{&command.CollectionCreateCommand{Name: ""}},
			{&command.CollectionCreateCommand{Name: "   "}},
		}

		for _, i := range commands {
			_, err := serv.Create(i.cmd)
			if err == nil {
				t.Errorf("era esperado um erro! nao pode ser cadastrado nulo")
				t.FailNow()
			}

			assert.ErrorIs(t, err, service.ErrNameIsNull, "a string nao pode estar vazia")
		}
	})

}

func TestCreateCollectionWithParent(t *testing.T) {
	tt := &collectionMock{
		CreateFunc: func(folder *entity.Collection) error {
			mockTable[folder.GetID()] = folder
			return nil
		},
		GetNextIdFunc: func() (int64, error) {
			seq = seq + 1
			return seq, nil
		},
		GetByIdFunc: func(id int64) (*entity.Collection, error) {
			collection := mockTable[id]
			if collection == nil {
				return nil, sql.ErrNoRows
			}

			return collection, nil
		},
	}

	serv := service.NewFolderService(tt)
	t.Run("parent not exist", func(t *testing.T) {
		cmd := &command.CollectionCreateCommand{Name: "Fantasy", ParentId: 120, Icon: "path_to_icon"}
		_, err := serv.Create(cmd)
		if err == nil {
			t.Errorf("deve retornar um erro, o pai nao exise! Err: %s", err)
			t.FailNow()
		}

		if errors.Is(err, service.ErrParentNotFound) {
			return
		}

		t.Errorf("não foi retornado o erro esperado! Err: %s", err)
		t.FailNow()
	})

	t.Run("valid parent", func(t *testing.T) {
		mockTable[123] = entity.NewCollection(123, "book")

		cmd := &command.CollectionCreateCommand{Name: "Fantasy", ParentId: 123, Icon: "path_to_icon"}
		resp, err := serv.Create(cmd)
		if err != nil {
			t.Errorf("erro inesperado! Err: %s", err)
			t.FailNow()
		}

		assert.Equal(t, cmd.ParentId, resp.Parent.Id, "O id retornado do pai não era o esperado")
	})

}

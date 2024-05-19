package service_test

import (
	"errors"
	"go-read-list/src/business/collection/application/command"
	service "go-read-list/src/business/collection/application/services"
	"go-read-list/src/business/collection/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

var seq int64 = 121
var mockTable = map[int64]*entity.Collection{}

func TestCreateCollectionWithoutParent(t *testing.T) {
	tt := &collectionMock{
		CreateFunc: func(folder *entity.Collection) error {
			if folder.GetID() == 0 {
				return errors.New("nao pode ser 0")
			}

			mockTable[folder.GetID()] = folder
			return nil
		},
		GetNextIdFunc: func() (int64, error) {
			seq = seq + 1
			return seq, nil
		},
	}

	serv := service.NewFolderService(tt)

	commands := []struct {
		cmd *command.CollectionCreateCommand
	}{
		{&command.CollectionCreateCommand{Name: "Move", Icon: "path_to_icon"}},
		{&command.CollectionCreateCommand{Name: "Golang"}},
	}

	for _, i := range commands {
		resp, err := serv.Create(i.cmd)
		if err != nil {
			t.Errorf("não foi retornado o erro esperado! Err: %s", err)
		}

		assert.Equal(t, i.cmd.Name, resp.Name, "O nome não retornou como esperado")
		assert.Equal(t, i.cmd.Icon, resp.Icon, "O icone não retornou como esperado")
		assert.NotEmpty(t, resp.CreateAt, "A data do cadastro do registro deveria estar preenchida")
		assert.NotEmpty(t, resp.UpdateAt, "A data de alteração deveria estar preenchida")
	}
}

func TestCreateCollectionParent(t *testing.T) {
	tt := &collectionMock{
		CreateFunc: func(folder *entity.Collection) error {
			if folder.GetID() == 0 {
				return errors.New("nao pode ser 0")
			}

			mockTable[folder.GetID()] = folder
			return nil
		},
		GetNextIdFunc: func() (int64, error) {
			seq = seq + 1
			return seq, nil
		},
	}

	serv := service.NewFolderService(tt)

	commands := []struct {
		cmd *command.CollectionCreateCommand
	}{
		{&command.CollectionCreateCommand{Name: "Move", ParentId: 122, Icon: "path_to_icon"}},
	}

	for _, i := range commands {
		resp, err := serv.Create(i.cmd)
		if err != nil {
			t.Errorf("não foi retornado o erro esperado! Err: %s", err)
		}

		assert.Equal(t, i.cmd.Name, resp.Name, "O nome não retornou como esperado")
		assert.Equal(t, i.cmd.Icon, resp.Icon, "O icone não retornou como esperado")
		assert.NotEmpty(t, resp.CreateAt, "A data do cadastro do registro deveria estar preenchida")
		assert.NotEmpty(t, resp.UpdateAt, "A data de alteração deveria estar preenchida")
	}
}

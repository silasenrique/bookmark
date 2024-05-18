package service

import (
	"errors"
)

func (f *CollectionService) DeleteById(id int64, recusive bool) error {
	if id == 0 {
		return errors.New("é necessário informar um id para ser excluido")
	}

	count, err := f.rep.CountDependencies(id)
	if err != nil {
		return err
	}

	if count > 0 && !recusive {
		return errors.New("a exclusão não pode ser realizada! existe dependentes")
	}

	return f.rep.RemevoFolderById(id)
}

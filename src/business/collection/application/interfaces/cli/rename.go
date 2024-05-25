package cli

import (
	"errors"
	"go-read-list/src/business/collection/application/command"
	service "go-read-list/src/business/collection/application/services"

	"github.com/urfave/cli/v2"
)

func Rename(serv *service.CollectionService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		id := ctx.Int64("id")
		title := ctx.String("title")

		if id == 0 {
			return errors.New("é necessário informar o id")
		}

		if title == "" {
			return errors.New("é necessário informar o novo nome da coleção")
		}

		response, err := serv.Rename(&command.CollectionRenameCommand{
			Id:   id,
			Name: title,
		})

		if err != nil {
			return err
		}

		printOne(response)

		return nil
	}
}

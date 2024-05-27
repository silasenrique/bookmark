package cli

import (
	"go-read-list/src/business/collection/application/command"
	service "go-read-list/src/business/collection/application/services"

	"github.com/urfave/cli/v2"
)

func Rename(serv *service.CollectionService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		id := ctx.Int64("id")
		title := ctx.String("title")

		if id == 0 {
			return cli.Exit("é necessário informar o id", 0)
		}

		if title == "" {
			return cli.Exit("é necessário informar o novo nome da coleção", 0)
		}

		response, err := serv.Rename(&command.CollectionRenameCommand{
			Id:   id,
			Name: title,
		})

		if err != nil {
			return cli.Exit(err, 0)
		}

		printOneOmint(response, nil)

		return nil
	}
}

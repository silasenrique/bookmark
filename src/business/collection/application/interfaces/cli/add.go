package cli

import (
	"go-read-list/src/business/collection/application/command"
	service "go-read-list/src/business/collection/application/services"

	"github.com/urfave/cli/v2"
)

func Add(serv *service.CollectionService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		folderCommand := &command.CollectionCreateCommand{
			Name:     ctx.String("name"),
			ParentId: ctx.Int64("parentId"),
			Icon:     ctx.String("icon"),
		}

		response, err := serv.Create(folderCommand)
		if err != nil {
			return err
		}

		printOne(response)

		return nil
	}
}

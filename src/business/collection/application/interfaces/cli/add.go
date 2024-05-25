package cli

import (
	"fmt"
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

		fmt.Printf("%-5v %-5v %-10v %-19v %-30v\n",
			response.Icon,
			response.Id,
			response.Name,
			response.CreateAt,
			response.UpdateAt)

		if response.Parent != nil {
			fmt.Printf("|--> Parent: %d - %s\n",
				response.Parent.Id,
				response.Parent.Name)
		}

		return nil
	}
}

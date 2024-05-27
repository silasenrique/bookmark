package cli

import (
	"go-read-list/src/business/collection/application/command"
	service "go-read-list/src/business/collection/application/services"

	"github.com/urfave/cli/v2"
)

func Update(serv *service.CollectionService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		id := ctx.Int64("id")
		name := ctx.String("name")
		icon := ctx.String("icon")
		parentId := ctx.Int64("parentId")
		removeParent := ctx.Bool("removeParent")

		if parentId != 0 && removeParent {
			return cli.Exit("deve ser escolhido apenas uma das opções: parentId ou removeParent", 0)
		}

		if name == "" && icon == "" && parentId == 0 && !removeParent {
			return cli.Exit("é necessário informar algum paremetro alem do id", 0)
		}

		resp, err := serv.Update(
			&command.CollectionUpdateCommand{
				Id:           id,
				Name:         name,
				ParentId:     parentId,
				Icon:         icon,
				RemoveParent: removeParent,
			},
		)

		if err != nil {
			return cli.Exit(err, 0)
		}

		printOneOmint(resp, nil)

		return nil
	}
}

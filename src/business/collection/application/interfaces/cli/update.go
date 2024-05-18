package cli

import (
	"errors"
	"fmt"
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
			return errors.New("deve ser escolhido apenas uma das opções: parentId ou removeParent")
		}

		if name == "" && icon == "" && parentId == 0 && !removeParent {
			return errors.New("é necessário informar algum paremetro alem do id")
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
			return err
		}

		fmt.Printf("%-35v | %-19v | %-19v | %-30v\n",
			resp.Id,
			resp.CreateAt,
			resp.UpdateAt,
			resp.Name)

		return nil
	}
}

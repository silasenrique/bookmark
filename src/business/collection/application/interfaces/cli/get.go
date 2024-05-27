package cli

import (
	service "go-read-list/src/business/collection/application/services"

	"github.com/urfave/cli/v2"
)

func GetById(serv *service.CollectionService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		id := ctx.Int64("id")

		colle, err := serv.GetById(id)
		if err != nil {
			return cli.Exit(err, 0)
		}

		printOneOmint(colle, ctx.StringSlice("omit"))

		return nil
	}
}

func GetByName(serv *service.CollectionService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		name := ctx.String("name")

		colle, err := serv.GetByName(name)
		if err != nil {
			return cli.Exit(err, 0)
		}

		printManyOmit(colle, ctx.StringSlice("omit"))

		return nil
	}
}

func GetByParentId(serv *service.CollectionService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		id := ctx.Int64("parentId")

		colle, err := serv.GetByParentId(id)
		if err != nil {
			return cli.Exit(err, 0)
		}

		printManyOmit(colle, ctx.StringSlice("omit"))

		return nil
	}
}

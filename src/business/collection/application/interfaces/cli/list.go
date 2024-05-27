package cli

import (
	service "go-read-list/src/business/collection/application/services"

	"github.com/urfave/cli/v2"
)

func List(serv *service.CollectionService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		col, err := serv.List()
		if err != nil {
			return cli.Exit(err, 0)
		}

		printManyOmit(col, ctx.StringSlice("omit"))

		return nil
	}
}

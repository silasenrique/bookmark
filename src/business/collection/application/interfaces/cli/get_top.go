package cli

import (
	service "go-read-list/src/business/collection/application/services"

	"github.com/urfave/cli/v2"
)

func GetTopLevel(serv *service.CollectionService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		collection, err := serv.ListAllTopCollections()
		if err != nil {
			return cli.Exit(err, 0)
		}

		printManyOmit(collection, ctx.StringSlice("omit"))

		return nil
	}
}

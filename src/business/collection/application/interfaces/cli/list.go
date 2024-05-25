package cli

import (
	service "go-read-list/src/business/collection/application/services"

	"github.com/urfave/cli/v2"
)

func List(serv *service.CollectionService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		col, err := serv.List()
		if err != nil {
			return err
		}

		printMany(col)

		return nil
	}
}

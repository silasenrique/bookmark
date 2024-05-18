package cli

import (
	service "go-read-list/src/business/collection/application/services"

	"github.com/urfave/cli/v2"
)

func RemoveById(serv *service.CollectionService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		return serv.DeleteById(ctx.Int64("id"), ctx.Bool("recursive"))
	}
}

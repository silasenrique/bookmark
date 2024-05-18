package read

import (
	service "go-read-list/src/business/bookmark/application/services"

	"github.com/urfave/cli/v2"
)

func RemoveById(serv *service.BookmarkService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		return nil
	}
}

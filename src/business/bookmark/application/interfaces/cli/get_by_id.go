package read

import (
	"fmt"
	service "go-read-list/src/business/bookmark/application/services"

	"github.com/savioxavier/termlink"
	"github.com/urfave/cli/v2"
)

func GetById(serv *service.BookmarkService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		id := ctx.String("id")

		read, err := serv.GetById(id)
		if err != nil {
			return err
		}

		hiperlink := termlink.Link(read.Title, read.Url)
		fmt.Printf("%-20v%-20v%-37v%-30v%-37v\n",
			read.CreateAt,
			read.UpdateAt,
			read.Id,
			hiperlink,
			read.Cover,
		)

		return nil
	}
}

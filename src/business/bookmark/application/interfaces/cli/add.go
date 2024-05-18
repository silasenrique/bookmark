package read

import (
	"fmt"
	"go-read-list/src/business/bookmark/application/command"
	service "go-read-list/src/business/bookmark/application/services"

	"github.com/savioxavier/termlink"
	"github.com/urfave/cli/v2"
)

func Add(serv *service.BookmarkService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		read, err := serv.Create(&command.BookmarkCreateCommand{
			Url:   ctx.String("uri"),
			Title: ctx.String("title"),
			Cover: ctx.String("cover"),
		})

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

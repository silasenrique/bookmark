package cli

import (
	"fmt"
	service "go-read-list/src/business/collection/application/services"

	"github.com/urfave/cli/v2"
)

func GetTopLevel(serv *service.CollectionService) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		collection, err := serv.ListAllTopCollections()
		if err != nil {
			return err
		}

		for _, r := range *collection {

			fmt.Printf("%-5v | %-10v | %-19v | %-30v\n",
				r.Id,
				r.Name,
				r.CreateAt,
				r.UpdateAt)
		}

		return nil
	}
}

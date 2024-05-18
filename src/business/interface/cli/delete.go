package cli

import (
	booki "go-read-list/src/business/bookmark/application/interfaces/cli"
	colli "go-read-list/src/business/collection/application/interfaces/cli"

	"github.com/urfave/cli/v2"
)

func deleteCmds(wrapperRepository *wrapper) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "collection",
			Aliases: []string{"c"},
			Usage:   "deleta uma collection",
			Subcommands: []*cli.Command{
				{
					Name:    "byId",
					Aliases: []string{"id"},
					Usage:   "excluir uma coleção utilizando o id",
					Flags: []cli.Flag{
						&cli.Int64Flag{Name: "id", Aliases: []string{"i"}, Usage: "id da pasta", Required: true},
						&cli.BoolFlag{Name: "recursive", Aliases: []string{"r"}, Usage: "indica se pode excluir quem depende dessa collection", Value: false},
					},
					Action: colli.RemoveById(wrapperRepository.collection),
				},
			},
		},
		{
			Name:    "bookmark",
			Aliases: []string{"b"},
			Usage:   "deleta um favorito",
			Subcommands: []*cli.Command{
				{
					Name:    "byId",
					Aliases: []string{"id"},
					Usage:   "excluir um bookmark",
					Flags: []cli.Flag{
						&cli.Int64Flag{Name: "id", Aliases: []string{"i"}, Usage: "id da pasta", Required: true},
					},
					Action: booki.RemoveById(wrapperRepository.bookmark),
				},
			},
		},
	}
}

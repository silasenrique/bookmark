package cli

import (
	booki "go-read-list/src/business/bookmark/application/interfaces/cli"
	colli "go-read-list/src/business/collection/application/interfaces/cli"

	"github.com/urfave/cli/v2"
)

func getCmds(wrapperRepository *wrapper) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "collection",
			Usage:   "obtem uma coleção",
			Aliases: []string{"c"},
			Subcommands: []*cli.Command{
				{
					Name:  "byId",
					Usage: "buscar utilizando o identificador unico",
					Flags: []cli.Flag{
						&cli.Int64Flag{Name: "id", Aliases: []string{"i"}, Usage: "id da coleção", Required: true},
						&cli.StringSliceFlag{Name: "omit", Aliases: []string{"o"}, Usage: "omitir uma coluna"},
					},
					Action: colli.GetById(wrapperRepository.collection),
				},
				{
					Name:  "byName",
					Usage: "buscar utilizando o nome",
					Flags: []cli.Flag{
						&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Usage: "nome da coleção. obs: é case sensitive, caso não exista nenhuma combinação, não irá retornar nada", Required: true},
					},
					Action: colli.GetByName(wrapperRepository.collection),
				},
				{
					Name:  "byParent",
					Usage: "buscar utilizando o nome",
					Flags: []cli.Flag{
						&cli.Int64Flag{Name: "parentId", Aliases: []string{"i"}, Usage: "nome da coleção. obs: é case sensitive, caso não exista nenhuma combinação, não irá retornar nada", Required: true},
					},
					Action: colli.GetByParentId(wrapperRepository.collection),
				},
			},
		},
		{
			Name:    "bookmark",
			Usage:   "obtem uma coleção",
			Aliases: []string{"b"},
			Subcommands: []*cli.Command{
				{
					Name:    "get",
					Aliases: []string{"g"},
					Usage:   "criando um novo favorito",
					Flags: []cli.Flag{
						&cli.StringFlag{Name: "id", Aliases: []string{"i"}, Usage: "id do favorito", Required: true},
					},
					Action: booki.GetById(wrapperRepository.bookmark),
				},
			},
		},
	}
}

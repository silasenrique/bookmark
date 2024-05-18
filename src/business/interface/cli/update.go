package cli

import (
	booki "go-read-list/src/business/bookmark/application/interfaces/cli"
	colli "go-read-list/src/business/collection/application/interfaces/cli"

	"github.com/urfave/cli/v2"
)

func updateCmds(wrapperRepository *wrapper) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "collection",
			Aliases: []string{"c"},
			Usage:   "Lista uma coleção",
			Subcommands: []*cli.Command{
				{
					Name:    "full",
					Aliases: []string{"f"},
					Usage:   "atualizar as informações de uma coleção",
					Flags: []cli.Flag{
						&cli.Int64Flag{Name: "id", Aliases: []string{"i"}, Usage: "identificador da collection", Required: true},
						&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Usage: "nome da collection"},
						&cli.BoolFlag{Name: "removeParent", Aliases: []string{"rp", "r"}, Usage: "indica se deve se desvincular da collection parent", Value: false},
						&cli.Int64Flag{Name: "parentId", Aliases: []string{"p", "pi"}, Usage: "identificador da nova collection pai"},
						&cli.StringFlag{Name: "icon", Aliases: []string{"c"}, Usage: "icone para a collection"},
					},
					Action: colli.Update(wrapperRepository.collection),
				},
				{
					Name:    "rename",
					Aliases: []string{"rn"},
					Usage:   "atualizar o titulo de uma coleção",
					Flags: []cli.Flag{
						&cli.Int64Flag{Name: "id", Aliases: []string{"i"}, Usage: "identificador da collection", Required: true},
						&cli.StringFlag{Name: "title", Aliases: []string{"t"}, Usage: "novo titulo da collection", Required: true},
					},
					Action: colli.Rename(wrapperRepository.collection),
				},
			},
		},
		{
			Name:    "bookmark",
			Aliases: []string{"b"},
			Usage:   "criando um novo favorito",
			Action:  booki.Add(wrapperRepository.bookmark),
		},
	}
}

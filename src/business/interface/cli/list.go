package cli

import (
	booki "go-read-list/src/business/bookmark/application/interfaces/cli"
	colli "go-read-list/src/business/collection/application/interfaces/cli"

	"github.com/urfave/cli/v2"
)

func listCmds(wrapperRepository *wrapper) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "collection",
			Aliases: []string{"c"},
			Usage:   "Lista uma coleção",
			Subcommands: []*cli.Command{
				{
					Name:    "top",
					Usage:   "listar todas as coleções solitárias",
					Aliases: []string{"p"},
					Action:  colli.GetTopLevel(wrapperRepository.collection),
				},
				{
					Name:    "tree", // TODO: ainda não ta pronta não
					Usage:   "listar as coleções de forma hierarquica de arvore",
					Action:  colli.GetTopLevel(wrapperRepository.collection),
					Aliases: []string{"t"},
				},
				{
					Name:    "all", // TODO: ainda não ta pronta não
					Usage:   "listar todas as coleções criadas",
					Action:  colli.List(wrapperRepository.collection),
					Aliases: []string{"a"},
				},
			},
		},
		{
			Name:    "bookmark",
			Aliases: []string{"b"},
			Usage:   "criando um novo favorito",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "title", Aliases: []string{"t"}, Usage: "nome do favorito", Required: true},
				&cli.StringFlag{Name: "uri", Aliases: []string{"u"}, Usage: "url do favorito", Required: true},
				&cli.StringFlag{Name: "cover", Aliases: []string{"c"}, Usage: "url da imagem"},
				&cli.StringFlag{Name: "pathId", Aliases: []string{"p"}, Usage: "id da pasta"},
			},
			Action: booki.Add(wrapperRepository.bookmark),
		},
	}
}

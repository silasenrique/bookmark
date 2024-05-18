package cli

import (
	booki "go-read-list/src/business/bookmark/application/interfaces/cli"
	colli "go-read-list/src/business/collection/application/interfaces/cli"

	"github.com/urfave/cli/v2"
)

func newCmds(wrapperRepository *wrapper) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "collection",
			Aliases: []string{"c"},
			Usage:   "criar nova coleção",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Usage: "nome da collection", Required: true},
				&cli.Int64Flag{Name: "parentId", Aliases: []string{"p"}, Usage: "identificador da collection pai"},
				&cli.StringFlag{Name: "icon", Aliases: []string{"c"}, Usage: "icone para a collection"},
			},
			Action: colli.Add(wrapperRepository.collection),
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

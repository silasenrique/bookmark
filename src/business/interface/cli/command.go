package cli

import (
	"github.com/urfave/cli/v2"
)

func LoadCommands(wrapperRepository *wrapper) *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "new",
				Aliases:     []string{"n"},
				Usage:       "Nova collection, bookmark",
				UsageText:   "Cadastra uma nova collection ou um novo bookmark",
				Subcommands: newCmds(wrapperRepository),
			},
			{
				Name:        "list",
				Aliases:     []string{"l"},
				Usage:       "Retorna uma lista de collection, bookmark",
				UsageText:   "Permite listar uma coleção ou um bookmark",
				Subcommands: listCmds(wrapperRepository),
			},
			{
				Name:        "update",
				Aliases:     []string{"u"},
				Usage:       "Atualiza collection, bookmark",
				UsageText:   "Atualiza as informações das collections ou bookmarks",
				Subcommands: updateCmds(wrapperRepository),
			},
			{
				Name:        "delete",
				Aliases:     []string{"d"},
				Usage:       "Deleta collection, bookmark",
				UsageText:   "Deleta as informações das collections ou bookmarks",
				Subcommands: deleteCmds(wrapperRepository),
			},
			{
				Name:        "get",
				Aliases:     []string{"g"},
				Usage:       "Busca collection, bookmark",
				UsageText:   "Busca as informações de uma collections ou um bookmarks",
				Subcommands: getCmds(wrapperRepository),
			},
		},
	}
}

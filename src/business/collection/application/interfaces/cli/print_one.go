package cli

import (
	"go-read-list/src/business/collection/application/command"
)

type Column struct {
	Name  string
	Value interface{}
	Omit  bool
}

func printOneOmint(colle *command.CollectionResponse, omitColumns []string) {
	columnToOmit := map[string]bool{}

	for _, column := range omitColumns {
		columnToOmit[column] = true
	}

	data := []*Column{
		{Name: "icon", Value: colle.Icon},
		{Name: "id", Value: colle.Id},
		{Name: "creatAt", Value: colle.CreateAt},
		{Name: "updateAt", Value: colle.UpdateAt},
		{Name: "name", Value: colle.Name},
	}

	for _, v := range data {
		if columnToOmit[v.Name] {
			v.Omit = true
		}
	}

	printData(data, true)
	printData(data, false)
}

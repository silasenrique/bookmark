package cli

import (
	"go-read-list/src/business/collection/application/mapper"
)

func printManyOmit(colle *mapper.CollectionsResponse, omitColumns []string) {
	columnToOmit := map[string]bool{}

	for _, column := range omitColumns {
		columnToOmit[column] = true
	}

	data := []*Column{
		{Name: "icon"},
		{Name: "id"},
		{Name: "creatAt"},
		{Name: "updateAt"},
		{Name: "name"},
	}

	for _, v := range data {
		if columnToOmit[v.Name] {
			v.Omit = true
		}
	}

	printData(data, true)

	for _, o := range *colle {
		data := []*Column{
			{Name: "icon", Value: o.Icon, Omit: columnToOmit["icon"]},
			{Name: "id", Value: o.Id, Omit: columnToOmit["id"]},
			{Name: "creatAt", Value: o.CreateAt, Omit: columnToOmit["creatAt"]},
			{Name: "updateAt", Value: o.UpdateAt, Omit: columnToOmit["updateAt"]},
			{Name: "name", Value: o.Name, Omit: columnToOmit["name"]},
		}

		printData(data, false)
	}

}

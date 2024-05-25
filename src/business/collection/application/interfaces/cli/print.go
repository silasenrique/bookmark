package cli

import (
	"fmt"
	"go-read-list/src/business/collection/application/command"
	"go-read-list/src/business/collection/application/mapper"
	"strconv"
)

func printOne(colle *command.CollectionResponse) {
	fmt.Printf("%-5v %-5v %-19v %-19v %-10v\n",
		"icon",
		"id",
		"creatAt",
		"updateAt",
		"name",
	)

	fmt.Printf("%-5v %-5v %-19v %-19v %-10v \n",
		colle.Icon,
		colle.Id,
		colle.CreateAt,
		colle.UpdateAt,
		colle.Name,
	)
}

func printMany(colle *mapper.CollectionsResponse) {
	fmt.Printf("%-5v %-5v %-5v %-19v %-19v %-10v\n",
		"icon",
		"id",
		"parent",
		"creatAt",
		"updateAt",
		"name")

	for _, c := range *colle {

		parentId := ""
		if c.Parent != nil {
			parentId = strconv.FormatInt(c.Parent.Id, 2)
		}

		fmt.Printf("%-5v %-5v %-6v %-19v %-19v %-10v\n",
			c.Icon,
			c.Id,
			parentId,
			c.CreateAt,
			c.UpdateAt,
			c.Name)
	}
}

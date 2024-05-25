package cli

import (
	"fmt"
	"go-read-list/src/business/collection/application/command"
	"go-read-list/src/business/collection/application/mapper"
)

func printOne(colle *command.CollectionResponse) {
	fmt.Printf("%-5v %-5v %-10v %-19v %-30v\n",
		"icon",
		"id",
		"name",
		"creatAt",
		"updateAt")

	fmt.Printf("%-5v %-5v %-10v %-19v %-30v\n",
		colle.Icon,
		colle.Id,
		colle.Name,
		colle.CreateAt,
		colle.UpdateAt,
	)
}

func printMany(colle *mapper.CollectionsResponse) {
	fmt.Printf("%-5v %-5v %-10v %-19v %-30v\n",
		"icon",
		"id",
		"name",
		"creatAt",
		"updateAt")

	for _, c := range *colle {
		fmt.Printf("%-5v %-5v %-10v %-19v %-30v\n",
			c.Icon,
			c.Id,
			c.Name,
			c.CreateAt,
			c.UpdateAt)
	}
}

package cli

import (
	"fmt"
	"go-read-list/src/business/collection/application/command"
	"go-read-list/src/business/collection/application/mapper"
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
	fmt.Printf("%-5v %-5v %-19v %-19v %-10v\n",
		"icon",
		"id",
		"creatAt",
		"updateAt",
		"name")

	for _, c := range *colle {
		fmt.Printf("%-5v %-5v %-19v %-19v %-10v\n",
			c.Icon,
			c.Id,
			c.CreateAt,
			c.UpdateAt,
			c.Name)
	}
}

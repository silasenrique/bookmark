package cli

import (
	"fmt"
	"go-read-list/src/business/collection/application/command"
)

var omitColumn = map[string]bool{
	"icon":     true,
	"id":       true,
	"creatAt":  true,
	"updateAt": true,
	"name":     true,
}

var columnName = map[string]any{
	"icon":     "icon",
	"id":       "id",
	"creatAt":  "creatAt",
	"updateAt": "updateAt",
	"name":     "name",
}

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

func omit(omitMap map[string]bool, value map[string]interface{}) {
	text := ""
	for field, show := range omitMap {
		if show {
			switch field {
			case "icon":
				text = fmt.Sprintf("%v%-5v ", text, value["icon"])
			case "id":
				text = fmt.Sprintf("%v%-5v ", text, value["id"])
			case "creatAt":
				text = fmt.Sprintf("%v%-19v ", text, value["creatAt"])
			case "updateAt":
				text = fmt.Sprintf("%v%-19v ", text, value["updateAt"])
			case "name":
				text = fmt.Sprintf("%v%v", text, value["name"])
			}
		}
	}

	fmt.Printf("%s\n", text)
}

func printOneOmint(colle *command.CollectionResponse, omitColumns []string) {

	for _, field := range omitColumns {
		omitColumn[field] = false
	}

	omitValue := map[string]interface{}{
		"icon":     colle.Icon,
		"id":       colle.Id,
		"creatAt":  colle.CreateAt,
		"updateAt": colle.UpdateAt,
		"name":     colle.Name,
	}

	omit(omitColumn, columnName)
	omit(omitColumn, omitValue)

}

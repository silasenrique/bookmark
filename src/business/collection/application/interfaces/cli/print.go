package cli

import "fmt"

func getValue(column *Column, title bool) any {
	if title {
		return column.Name
	}

	return column.Value
}

func printData(col []*Column, title bool) {
	text := ""
	for _, show := range col {
		if !show.Omit {
			switch show.Name {
			case "icon":
				text = fmt.Sprintf("%v%-5v ", text, getValue(show, title))
			case "id":
				text = fmt.Sprintf("%v%-5v ", text, getValue(show, title))
			case "creatAt":
				text = fmt.Sprintf("%v%-19v ", text, getValue(show, title))
			case "updateAt":
				text = fmt.Sprintf("%v%-19v ", text, getValue(show, title))
			case "name":
				text = fmt.Sprintf("%v%v", text, getValue(show, title))
			}
		}
	}

	fmt.Printf("%s\n", text)
}

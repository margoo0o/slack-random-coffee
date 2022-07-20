package group_generation

import (
	"fmt"
	"strings"
)

func formatGroups(subgroups [][]string) string {
	groupText := ""
	for i, subgroup := range subgroups {
		groupText += "\U00002022"
		for j, val := range subgroup {
			subgroup[j] = fmt.Sprintf(`<@%v>`, val)
		}
		groupText += fmt.Sprintln(strings.Join(subgroups[i][:], " \U00002615  "))
		fmt.Sprintln(strings.Join(subgroups[i][:], " \U00002615  "))
		groupText += "\n"
	}
	return groupText
}

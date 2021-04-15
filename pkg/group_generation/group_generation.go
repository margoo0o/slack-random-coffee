package group_generation

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// TODO Exclude the Coffee Bot users
// U01U2BQ2CSX
func GenerateGroups(users []string, groupSize int) string {
	// Seed based on the generated timestamp - without seeding, same pseudorandom sequence is generated
	rand.Seed(time.Now().UnixNano())
	// Shuffle order
	rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })

	// Split into groups of user defined size e.g 4
	var groups [][]string
	for {
		// No users should exit
		if len(users) == 0 {
			break
		}

		// necessary check to avoid slicing beyond slice capacity
		if len(users) < groupSize {
			groupSize = len(users)
		}

		groups = append(groups, users[0:groupSize])
		users = users[groupSize:]
	}

	// if group <= groupSize / 2, move people into other groups
	var lastGroup = groups[len(groups)-1]
	fmt.Println("last group", lastGroup)
	fmt.Println("got here", len(lastGroup), "     ", groupSize)
	if len(lastGroup) <= groupSize/2 {
		for i := range lastGroup {
			groups[i] = append(groups[i], lastGroup[i])
			fmt.Println("new group", groups[i])
		}
		//finally, remove small group
		groups = groups[:len(groups)-1]
	}
	fmt.Println(groups)

	return formatGroups(groups)
}

func formatGroups(subgroups [][]string) string {
	groupText := ""
	for i, subgroup := range subgroups {
		for j, val := range subgroup {
			subgroup[j] = fmt.Sprintf(`<@%v>`, val)
		}
		groupText += fmt.Sprintln(strings.Join(subgroups[i][:], " \U00002615  "))
	}
	return groupText
}

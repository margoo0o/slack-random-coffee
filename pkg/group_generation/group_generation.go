package group_generation

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateGroups(users []string, groupSize int) string {
	userDefinedGroupSize := groupSize
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
	lastGroup := groups[len(groups)-1]
	if len(lastGroup) <= userDefinedGroupSize / 2 {
		// Round robin approach - 1st person in small group goes to first full group, 2nd goes to 2nd group etc.
		for i := range lastGroup {
			appendUser := lastGroup[i]
			fmt.Println("append user", appendUser)
			groups[i] = copyAndAppend(groups[i], lastGroup[i])
			fmt.Println(groups[i])
		}
		// Finally, pop the last group once all members have been reallocated
		groups = groups[:len(groups)-1]
	}
	fmt.Println("groups")
	fmt.Println(groups)

	return formatGroups(groups)
}

// Appending to a slice... yikes!
// The confusion I had is due to the fact that append both changes the underlying array and returns a new slice (since the length changes).
// You'd imagine that it copies that backing array, but it doesn't, it just allocates a new slice object that points at it.
func copyAndAppend(i []string, vals ...string) []string {
	j := make([]string, len(i), len(i)+len(vals))
	copy(j, i)
	return append(j, vals...)
}


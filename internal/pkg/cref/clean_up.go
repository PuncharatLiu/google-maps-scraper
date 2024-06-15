package cref

import (
	"strings"
)

func CleanUp() {
	for i, name := range names {
		names[i] = strings.ReplaceAll(name, " (copy)", "")
	}

	// Debug
	// for i := 1; i <= 20; i++ {
	// 	fmt.Println(names[i])
	// }
}

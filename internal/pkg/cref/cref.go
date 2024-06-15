package cref

import (
	"fmt"
)

func Cref(gmapsTitles []string) {
	// Debug
	fmt.Println("gmapsData from cref pkg: ", gmapsTitles)

	// Mock data
	gmapsTitles_mock := []string{
		"Genuine Disability Services",
		"Gunnedah Preschool",
		"some name",
		"another name",
	}

	mdTitlesMap := make(map[string]bool)
	for _, title := range names {
		// names is the titles from monday board
		mdTitlesMap[title] = true
	}

	// for _, title := range gmapsTitles {
	// 	if _, exists := mdTitlesMap[title]; exists {
	// 		text := fmt.Sprintf("%s: In board", title)
	// 		fmt.Println(text)
	// 	} else {
	// 		text := fmt.Sprintf("%s: Not in board", title)
	// 		fmt.Println(text)
	// 	}
	// }

	// mock
	for _, title := range gmapsTitles_mock {
		if _, exists := mdTitlesMap[title]; exists {
			text := fmt.Sprintf("%s: In board", title)
			fmt.Println(text)
		} else {
			text := fmt.Sprintf("%s: Not in board", title)
			fmt.Println(text)
		}
	}

}

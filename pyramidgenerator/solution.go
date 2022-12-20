package pyramidgenerator

import (
	"fmt"
	"strings"
)

func Generator(n int) []string {
	if n == 1 {
		return []string{"#"}
	}
	pyramid := []string{}

	max := n*2 - 1
	for i := 1; i <= max; i = i + 2 {
		mid := strings.Repeat("#", i)
		spaces := max - i
		line := strings.Repeat(" ", spaces/2) + mid + strings.Repeat(" ", spaces/2)
		pyramid = append(pyramid, line)
		fmt.Println(line)
	}
	return pyramid
}

package pyramidgenerator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	cases := []struct {
		input    int
		expected []string
	}{
		{input: 1, expected: []string{"#"}},
		{input: 2, expected: []string{
			" # ",
			"###"}},
		{input: 3, expected: []string{
			"  #  ",
			" ### ",
			"#####"}},
		{input: 10, expected: []string{
			"         #         ",
			"        ###        ",
			"       #####       ",
			"      #######      ",
			"     #########     ",
			"    ###########    ",
			"   #############   ",
			"  ###############  ",
			" ################# ",
			"###################"}},
	}

	for _, c := range cases {
		out := Generator(c.input)
		assert.Equal(t, c.expected, out)
	}
}

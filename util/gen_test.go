package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenAliasFor(t *testing.T) {
	assert := assert.New(t)
	testcases := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    `TITLE ALL UPCASE`,
			Expected: `title-all-upcase`,
		},
		{
			Input:    `Title without symbols`,
			Expected: `title-without-symbols`,
		},
		{
			Input:    `!@#$%^&*()remove!@#$%symbols[]|{}":<>?,./;'test`,
			Expected: `remove-symbols-test`,
		},
		{
			Input:    `with digits 123 456`,
			Expected: `with-digits-123-456`,
		},
	}

	for _, e := range testcases {
		output := GenAliasFor(e.Input)
		assert.Equal(output, e.Expected)
		fmt.Printf("input: %s\noutput: %s\nexpected: %s\n\n", e.Input, output, e.Expected)
	}
}

package checkmsg

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckExclusionWord(t *testing.T) {
	err := CheckExclusionWord("testdata/testExclusion1.txt", "hi! it's is test text")
	assert.Equal(t, nil, err)

	err = CheckExclusionWord("testdata/testExclusion2.txt", "hi! it's is test text")
	assert.Equal(t, errors.New("Found Exclusion Word"), err)
}

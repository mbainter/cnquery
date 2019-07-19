package stringslice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mondoo.io/mondoo/motor/stringslice"
)

func TestContains(t *testing.T) {
	assert.True(t, stringslice.Contains([]string{"ab", "aa"}, "ab"))
	assert.False(t, stringslice.Contains([]string{"ab", "aa"}, "a"))
	assert.False(t, stringslice.Contains([]string{"ab", "aa"}, "bs"))
	assert.True(t, stringslice.Contains([]string{"hello", "world"}, "world"))
	assert.True(t, stringslice.Contains([]string{"hello", "world"}, "hello"))
	assert.False(t, stringslice.Contains([]string{"hello", "world"}, "john"))
}

func TestRemoveEmpty(t *testing.T) {
	assert.Equal(t, []string{"aa"}, stringslice.RemoveEmpty([]string{"", "aa"}))
	assert.Equal(t, []string{"aa"}, stringslice.RemoveEmpty([]string{"aa", ""}))
	assert.Equal(t, []string{"aa", "ab"}, stringslice.RemoveEmpty([]string{"", "aa", "", "ab", ""}))
}

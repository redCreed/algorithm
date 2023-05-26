package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPrefix(t *testing.T) {
	s := []string{"flower", "flow", "flight"}
	prefix := GetPrefix(s)
	as := assert.New(t)
	as.Equal("fl", prefix, "2个值必须相等")
}

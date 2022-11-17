package constant_test

import (
	"strconv"
	"testing"

	"github.com/unsurper/goGB2011/constant"

	"github.com/stretchr/testify/assert"
)

func TestAction_String(t *testing.T) {
	for k, v := range constant.ActionNames {
		assert.Equal(t, k.String(), "["+strconv.Itoa(int(k))+"]"+v)
	}
	assert.Equal(t, constant.Action(200).String(), "[200]")
}

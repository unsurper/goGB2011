package constant_test

import (
	"strconv"
	"testing"

	"github.com/unsurper/goGB2011/constant"

	"github.com/stretchr/testify/assert"
)

func TestEquipmentType_String(t *testing.T) {
	for k, v := range constant.EquipmentTypeNames {
		assert.Equal(t, k.String(), "["+strconv.Itoa(int(k))+"]"+v)
	}
	assert.Equal(t, constant.EquipmentType(200).String(), "[200]")
}

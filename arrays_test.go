package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidArray(test *testing.T) {

	Arr := Array{"test"}
	assert.False(test, Arr.isValid())

	Arr = Array{true}
	assert.False(test, Arr.isValid())

	Arr = Array{1.2}
	assert.False(test, Arr.isValid())
}

func TestMapArray(test *testing.T) {

	Arr := Array{map[string]string{"key1": "value1", "key2": "value2"}}

	assert.True(test, Arr.Exists("key1"))
	assert.Equal(test, "value1", Arr.Get("key1"))
	assert.Equal(test, "value2", Arr.Get("key2"))
	assert.Equal(test, map[interface{}]interface{}{"key1": "value1"}, Arr.Only("key1"))
	assert.True(test, Arr.AddMap("key3", "value3").Has("key3"))
}

func TestSliceArray(test *testing.T) {

	Arr := Array{[]string{"key1", "key2"}}

	assert.True(test, Arr.Has("key1"))
	assert.Equal(test, "key1", Arr.Index(0))
	assert.Equal(test, "key2", Arr.Index(1))
}
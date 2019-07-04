package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidCollection(test *testing.T) {

	Arr := New("test")
	assert.False(test, Arr.IsValid())

	Arr = New(true)
	assert.False(test, Arr.IsValid())

	Arr = New(1.2)
	assert.False(test, Arr.IsValid())
}

func TestMapCollection(test *testing.T) {

	Arr := New(map[string]string{"key1": "value1", "key2": "value2"})

	assert.True(test, Arr.Exists("key1"))
	assert.Equal(test, "value1", Arr.Get("key1"))
	assert.Equal(test, "value2", Arr.Get("key2"))

	assert.Equal(test, map[interface{}]interface{}{
		"key1": "value1",
	}, Arr.Only("key1"))

	assert.True(test, Arr.AddMap("key3", "value3").Has("key3"))
}

func TestSliceCollection(test *testing.T) {

	Arr := New([]string{"key1", "key2"})

	assert.True(test, Arr.Has("key1"))
	assert.Equal(test, "key1", Arr.Index(0))
	assert.Equal(test, "key2", Arr.Index(1))
}

func TestStructCollection(test *testing.T) {
	mockStrcut := struct {
		Field1 string
		Field2 int
		Field3 []string
		Field4 map[string]string
	}{
		"field1",
		1,
		[]string{"test1", "test2"},
		map[string]string{"field1": "result"},
	}

	s := New(mockStrcut)

	assert.True(test, s.Has("Field1"))
	assert.Equal(test, "field1", s.Get("Field1"))
	assert.Equal(test, 1, s.Get("Field2"))
	assert.Equal(test, []string{"test1", "test2"}, s.Get("Field3"))
	assert.Equal(test, map[string]string{"field1": "result"}, s.Get("Field4"))
}

func TestJsonCollection(test *testing.T) {
	c := New([]string{"index1", "index2"})
	assert.Equal(test, "[\"index1\",\"index2\"]", c.ToJson())
	assert.Equal(test, []byte(c.ToJson()), c.ToBytes())
}

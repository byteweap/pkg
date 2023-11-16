package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValueMapArray(t *testing.T) {
	// Test case 1: Empty byte array
	data := []byte{}
	expected := []ValueMap{}
	result, err := NewValueMapArray(data)

	if assert.NoError(t, err) {
		assert.Equal(t, result, expected)
	}

	// Test case 2: Byte array with valid JSON
	data = []byte(`[{"key": "value"}]`)
	expected = []ValueMap{{"key": "value"}}
	result, err = NewValueMapArray(data)

	if assert.NoError(t, err) {
		assert.Equal(t, result, expected)
	}

	// Test case 3: Byte array with invalid JSON
	data = []byte(`invalid json`)
	_, err = NewValueMapArray(data)
	assert.Error(t, err)

}

func TestGetIntArray(t *testing.T) {

	vm := ValueMap{}
	vm["arr"] = []int{1, 2, 3}
	t.Logf("%v", vm.GetIntArray("arr"))
}

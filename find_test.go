package finder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Dog struct {
	ID     string
	Name   string
	Colour string
}

func TestFind(t *testing.T) {

	t.Run("test_array", func(t *testing.T) {
		arr := []Dog{
			{ID: "12345"},
			{ID: "12347"},
		}
		fieldsFromArr := Find(arr, "ID")
		assert.Contains(t, fieldsFromArr, "12345")
		assert.Contains(t, fieldsFromArr, "12347")
	})

	t.Run("test_map", func(t *testing.T) {
		m := map[string]Dog{
			"a": {ID: "12345"},
			"b": {ID: "12347"},
		}
		fieldsFromMap := Find(m, "ID")
		assert.Contains(t, fieldsFromMap, "12345")
		assert.Contains(t, fieldsFromMap, "12347")

	})

	t.Run("test_array_with_pointers", func(t *testing.T) {
		arrPointers := []*Dog{
			{ID: "12345"},
			{ID: "12347"},
		}
		fieldsFromArrPointers := Find(arrPointers, "ID")
		assert.Contains(t, fieldsFromArrPointers, "12345")
		assert.Contains(t, fieldsFromArrPointers, "12347")
	})
}

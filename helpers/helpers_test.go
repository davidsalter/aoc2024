package helpers

import (
	"reflect"
	"testing"
)

func TestRemoveFirstEntry(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	result := Remove(slice, 0)

	if !reflect.DeepEqual(result, []string{"b", "c", "d"}) {
		t.Error("Expected ['b', 'c', 'd'] but got", result)
	}
}

func TestRemoveMiddleEntry(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	result := Remove(slice, 2)

	if !reflect.DeepEqual(result, []string{"a", "b", "d"}) {
		t.Error("Expected ['a', 'b', 'd'] but got", result)
	}
}

func TestRemoveLastEntry(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	result := Remove(slice, 3)

	if !reflect.DeepEqual(result, []string{"a", "b", "c"}) {
		t.Error("Expected ['a', 'b', 'c'] but got", result)
	}
}

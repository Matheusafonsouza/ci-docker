package main

import (
	"reflect"
	"testing"
)

// Checks if values are equal
func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func TestSum(t *testing.T) {
	result := sum(2, 2)
	expected := 4
	assertEqual(t, result, expected)
}

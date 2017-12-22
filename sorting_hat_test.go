package potterapi

import (
	"reflect"
	"testing"
)

func TestSortingHat(t *testing.T) {
	c := ClientInitialize("1234")

	house, err := c.SortingHat()

	if err != nil {
		t.Fatal(err)
	}

	if reflect.Type.String(reflect.TypeOf(house)) == "string" {
		t.Log("You got a house!")
	} else {
		t.Fatal("No house for you...")
	}
}

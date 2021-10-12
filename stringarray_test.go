/*
 * (c) Oleg Puchinin 2021
 * puchininolegigorevich@gmail.com
 */

package opgolib

import (
	"fmt"
	"testing"
)

func dump_array(arr *StringArray) {
	for idx, one := range arr.Array_get_array() {
		fmt.Printf("%d %s\n", idx, one)
	}
}

func TestStringArray1(t *testing.T) {
	var (
		arr *StringArray
	//	idx int
	)
	//arr2 := []string{"Six", "Seven", "Eight"}
	arr = NewStringArray("One", 1)
	arr.Array_set_at(1, "Two")
	arr.Array_set_at(2, "Three")
	arr.Array_set_at(3, "Four")
	arr.Array_set_at(4, "Five")
	arr.Array_set_at(-1, "Zero")
	//arr.Array_truncate(5)
	//arr.Array_concat(arr2[:])
	//fmt.Printf("zero index %d\n", idx)
	//dump_array(arr)
	arr.Array_all_valid()
	//arr.HalfSize()
	//	os.Stdout.Write(arr.Array_join())
	//dump_array(arr)
}

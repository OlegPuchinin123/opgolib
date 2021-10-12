/*
 * (c) Oleg Puchinin 2021
 * puchininolegigorevich@gmail.com
 */

package opgolib

import (
	"bytes"
	"container/list"
	"errors"
)

type StringArray struct {
	m_arr         []string
	m_stack_index int
}

func (arr *StringArray) DoubleSize() error {
	var (
		new_array []string
	)
	new_array = make([]string, len(arr.m_arr)<<1)
	copy(new_array[:], arr.m_arr[:])
	arr.m_arr = new_array
	return nil
}

func (arr *StringArray) HalfSize() error {
	var (
		new_array []string
		new_len   int
	)
	new_len = len(arr.m_arr) / 2
	new_array = make([]string, new_len)
	copy(new_array[:], arr.m_arr[:new_len])
	arr.m_arr = new_array
	return nil
}

func (arr *StringArray) Array_set_at(idx int, s string) int {
	var (
		index     int
		new_array []string
	)
	index = len(arr.m_arr)
	if idx < 0 {
		new_array = make([]string, index*2)
		copy(new_array[index:], arr.m_arr)
		arr.m_arr = new_array
		arr.m_arr[index-1] = s
		return index - 1
	}
	if idx >= len(arr.m_arr) {
		arr.DoubleSize()
	}
	arr.m_arr[idx] = s
	return idx
}

func (arr *StringArray) Array_get_at(idx int) string {
	return arr.m_arr[idx]
}

func (arr *StringArray) Array_get_array() []string {
	return arr.m_arr
}

func (arr *StringArray) Array_concat(another []string) []string {
	var (
		new_array []string
	)
	new_array = make([]string, len(arr.m_arr)+len(another))
	copy(new_array, arr.m_arr)
	copy(new_array[len(arr.m_arr):], another)
	arr.m_arr = new_array
	return arr.m_arr
}

func (arr *StringArray) Array_truncate(size int) []string {
	var (
		new_array []string
	)
	new_array = make([]string, size)
	copy(new_array, arr.m_arr[:size])
	arr.m_arr = new_array
	return new_array
}

func (arr *StringArray) Array_all_valid() []string {
	var (
		new_array *StringArray
		idx       int
	)
	idx = 0
	new_array = NewStringArray("", 128)
	for _, one := range arr.m_arr {
		if one != "" {
			new_array.Array_set_at(idx, one)
			idx++
		}
	}
	new_array.Array_truncate(idx)
	arr.m_arr = new_array.m_arr
	return arr.m_arr
}

func (arr *StringArray) Array_to_list() *list.List {
	var (
		lst *list.List
		one string
	)
	lst = list.New()
	for _, one = range arr.m_arr {
		if one != "" {
			lst.PushBack(one)
		}
	}
	return lst
}

func (arr *StringArray) Array_from_list(lst *list.List) []string {
	var (
		e   *list.Element
		idx int
		a   *StringArray
	)
	if lst == nil {
		return nil
	}
	idx = 0
	a = NewStringArray("", 128)
	for e = lst.Front(); e != nil; e = e.Next() {
		a.Array_set_at(idx, e.Value.(string))
		idx++
	}
	a.Array_all_valid()
	return a.m_arr
}

func (arr *StringArray) Array_join() []byte {
	var (
		buf *bytes.Buffer
		one string
	)
	buf = bytes.NewBuffer(nil)
	for _, one = range arr.m_arr {
		buf.WriteString(one + "\n")
	}
	return buf.Bytes()
}

func (arr *StringArray) Stack_push(s string) error {
	arr.Array_set_at(arr.m_stack_index, s)
	arr.m_stack_index++
	return nil
}

func (arr *StringArray) Stack_pop() (string, error) {
	var (
		s string
	)
	if arr.m_stack_index < 0 {
		return "", errors.New("Out of bound.")
	}
	s = arr.Array_get_at(arr.m_stack_index)
	arr.m_stack_index--
	return s, nil
}

func NewStringArray(s string, count int) *StringArray {
	var (
		arr *StringArray
	)
	arr = new(StringArray)
	arr.m_arr = make([]string, count)
	arr.m_stack_index = 0
	arr.m_arr[0] = s
	return arr
}

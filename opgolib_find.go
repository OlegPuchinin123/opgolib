/*
 * (c) Oleg Puchinin 2021
 * puchininolegigorevich@gmail.com
 */

package opgolib

import (
	"os"
	"sort"
)

type DirTree struct {
	tree       *StringArray
	last_index int
}

func (tr *DirTree) StartReadTree() error {
	var (
		idx int
		s   string
	)
	for idx = 0; idx <= tr.last_index; idx++ {
		s = tr.tree.Array_get_at(idx)
		if last_slash(s) {
			tr.AddAnotherDir(s)
		}
	}
	return nil
}

func (tr *DirTree) AddAnotherDir(dirname string) {
	var (
		dir_entryes []os.DirEntry
		str         string
	)
	dir_entryes, _ = os.ReadDir(dirname)
	for _, file := range dir_entryes {
		//fmt.Printf("%s\n", file.Name())
		str = file.Name()
		str = dirname + str
		if file.IsDir() {
			str += "/"
		}
		tr.last_index++
		tr.tree.Array_set_at(tr.last_index, str)
	}
}

func last_slash(s string) bool {
	var (
		s2 string
	)
	s2 = string(s[len(s)-1:])
	if s2 == "/" {
		return true
	} else {
		return false
	}
}

func Find(dir_name string) []string {
	var (
		s   string
		ret []string
	)
	s = dir_name
	if !last_slash(s) {
		s += "/"
	}
	tr := new(DirTree)
	tr.tree = NewStringArray(s, 16)
	tr.last_index = 0
	tr.StartReadTree()
	arr := tr.tree.Array_get_array()
	sort.Strings(arr[:tr.last_index+1])
	ret = make([]string, tr.last_index+1)
	copy(ret, arr[:tr.last_index+1])
	arr = nil
	return ret
}

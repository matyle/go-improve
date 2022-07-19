package zuoer

import "errors"

type Undo []func() //undo函数类型

// (u *Undo) Add(function func()){
// 	*u = append(*u,function)
// }

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("no undo function")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil
	}

	functions = functions[:index]

	return nil
}

type IntSet2 struct {
	data map[int]struct{}
	undo Undo
}

func NewIntSet2() *IntSet2 {
	return &IntSet2{
		data: make(map[int]struct{}),
	}
}

func (set *IntSet2) Undo() error {
	return set.undo.Undo()
}

func (set *IntSet2) Contains(i int) bool {
	if _, ok := set.data[i]; ok {
		return true
	}
	return false
}

func (set *IntSet2) Add(i int) {
	if !set.Contains(i) {
		set.data[i] = struct{}{}
		set.undo.Add(func() { set.Delete(i) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet2) Delete(i int) {
	if set.Contains(i) {
		delete(set.data, i)
		set.undo.Add(func() { set.Add(i) })
	} else {
		set.undo.Add(nil)
	}
}

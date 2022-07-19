package zuoer

import "errors"

type IntSet struct {
	data map[int]struct{}
}

func NewIntSet() *IntSet {
	return &IntSet{
		data: make(map[int]struct{}),
	}
}

func (s *IntSet) Add(i int) {
	s.data[i] = struct{}{}
}

func (s *IntSet) Delete(i int) {
	delete(s.data, i)
}

func (s *IntSet) Contains(i int) bool {
	_, ok := s.data[i]
	return ok
}

type UndoInset struct {
	IntSet // 匿名结构体 实现多态效果
	funcs  []func()
}

func NewUndoInset() *UndoInset {
	return &UndoInset{
		IntSet: *NewIntSet(),
		funcs:  nil,
	}
}

func (u *UndoInset) Add(i int) {
	if !u.Contains(i) {
		u.IntSet.Add(i)
		u.funcs = append(u.funcs, func() {
			u.IntSet.Delete(i)
		})
	} else {
		u.funcs = append(u.funcs, nil)
	}
}

func (u *UndoInset) Delete(i int) {
	if u.Contains(i) {
		u.IntSet.Delete(i)
		u.funcs = append(u.funcs, func() {
			u.IntSet.Add(i)
		})
	} else {
		u.funcs = append(u.funcs, nil)
	}
}

func (u *UndoInset) Undo() error {
	if len(u.funcs) <= 0 {
		return errors.New("no undo")
	}
	i := len(u.funcs) - 1

	if function := u.funcs[i]; function != nil {
		function()
		u.funcs[i] = nil

	}
	u.funcs = u.funcs[:i]
	return nil
}

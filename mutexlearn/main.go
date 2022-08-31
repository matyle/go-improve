package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	Name string
	Age  int
	sync.RWMutex
}

type Students struct {
	ps map[string]*Person
	sync.RWMutex
}

var mapStudents = make(map[string]*Students, 10)

func main() {
	person := &Person{
		Name: "san",
		Age:  1,
	}
	s := &Students{
		ps: make(map[string]*Person, 10),
	}

	PutPerson(person, s)

	p := GetPerson("san", s)
	fmt.Println("main: ", p.Name, p.Age)
}

func GetPerson(name string, s *Students) *Person {
	s.RLock()
	defer s.RUnlock()
	for i := 0; i < 100000; i++ {
		go DoSomething(s.ps[name])
	}
	time.Sleep(time.Second * 10)
	return s.ps[name]
}

func PutPerson(p *Person, s *Students) {
	s.Lock()
	defer s.Unlock()
	s.ps[p.Name] = p
}

func Get(id string) *Students {
	s := mapStudents[id]
	s.RLock()
	defer s.RUnlock()
	return s
}

func Put(id string, s *Students) {
	fmt.Println("put start")
	s.Lock()
	mapStudents[id] = s
	s.Unlock()
	fmt.Println("put end")
}

func DoSomething(p *Person) {
	p.Lock()
	p.Age++
	p.Unlock()
	fmt.Println("do something end", p.Name, p.Age)
}

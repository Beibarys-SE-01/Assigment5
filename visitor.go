package main

import "fmt"

type students interface {
	getName() string
	accept(teacher)
}

type Miras struct {
	name string
}

func (s *Miras) accept(t teacher) {
	t.visitForMiras(s)
}

func (s *Miras) getName() string {
	return s.name
}

type Beibarys struct {
	name string
}

func (s *Beibarys) accept(t teacher) {
	t.visitForBeibarys(s)
}

func (s *Beibarys) getName() string {
	return s.name
}

type Sanat struct {
	name string
}

func (s *Sanat) accept(t teacher) {
	t.visitForSanat(s)
}

func (s *Sanat) getName() string {
	return s.name
}

type teacher interface {
	visitForBeibarys(*Beibarys)
	visitForSanat(*Sanat)
	visitForMiras(*Miras)
}

type Abi struct {
	list []string
}

func (t *Abi) visitForBeibarys(s *Beibarys)  {
	t.list = append(t.list, s.getName() + ", visited")
}

func (t *Abi) visitForSanat(s *Sanat)  {
	t.list = append(t.list, s.getName() + ", visited")
}

func (t *Abi) visitForMiras(s *Miras)  {
	t.list = append(t.list, s.getName() + ", visited")
}

func (t *Abi) showList() {
	for i, building := range t.list{
		fmt.Println(i, building)
	}
}

func ivsdsf() {
	miras := &Miras{"Miras"}
	beibarys := &Beibarys{"Beibarys"}
	sanat := &Sanat{"Sanat"}
	abi := Abi{[]string{}}
	abi.visitForBeibarys(beibarys)
	abi.visitForSanat(sanat)
	abi.visitForMiras(miras)
	abi.showList()
}
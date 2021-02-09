package main

import "fmt"

type Person2 struct {
	Name string
	FatIndex float64
}

func (c *Person2) String() string {
	return fmt.Sprintf("%s with fat index of body %v%%", c.Name, c.FatIndex)
}

func NewPerson (name string, fatIndex float64) *Person2 {
	return &Person2{name, fatIndex}
}

// My Interface Handler
type Exercises interface {
	AddExercise(e Exercises)
	DoExercise()
}

// My Base Handler
type PersonModifier struct {
	next Exercises
	person *Person2
}

func (c *PersonModifier) AddExercise(m Exercises)  {
	if c.next != nil{
		c.next.AddExercise(m)
	} else {
		c.next = m
	}
}

func (c *PersonModifier) DoExercise()  {
	if c.next != nil {
		c.next.DoExercise()
	}
}

func NewPersonModifier(c *Person2) *PersonModifier {
	return &PersonModifier{person: c}
}

// My First Concrete Handler
type PushUpExercises struct {
	PersonModifier
}

func (d *PushUpExercises) DoExercise() {
	fmt.Println("Grade job, man!10 push up was done.Continue!",d.person.Name,",just do it!")
	d.person.FatIndex -= 0.5
	d.PersonModifier.DoExercise()
}

// My Second Concrete Handler
func NewPushUpExercises(c *Person2) *PushUpExercises {
	return &PushUpExercises{PersonModifier{person: c}}
}

type PullUpExercises struct {
	PersonModifier
}

func (d *PullUpExercises) DoExercise() {
	fmt.Println("Grade job, man!10 pull up was done.Continue!",d.person.Name,",just do it!")
	d.person.FatIndex -= 0.7
	d.PersonModifier.DoExercise()
}

func NewPullUpExercises(c *Person2) *PullUpExercises {
	return &PullUpExercises{PersonModifier{person: c}}
}

func jds() {
	beibarys := NewPerson("Beibarys", 20.2)
	fmt.Println(beibarys)
	root := NewPersonModifier(beibarys)
	root.AddExercise(NewPushUpExercises(beibarys))
	root.AddExercise(NewPullUpExercises(beibarys))
	root.DoExercise()
	fmt.Println(beibarys)
}
package main

import "fmt"

type Creature struct {
	Name string
	Attack, Defense int
}
func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)",c.Name,c.Attack,c.Defense)
}
func NewCreature(name string,attack,defense int) *Creature {
	return &Creature{name,attack,defense}
}
// 1 Interface Handler
type Modifier interface {
	Add(m Modifier)
	Apply()
}
// 2 Base Handler
type CreatureModifier struct {
	next Modifier
	creature *Creature
}
func (c *CreatureModifier) Add(m Modifier)  {
	if c.next != nil{
		c.next.Add(m)
	} else {
		c.next = m
	}
}
func (c *CreatureModifier) Apply()  {
	if c.next != nil {
		c.next.Apply()
	}
}
func NewCreatureModifier(c *Creature) *CreatureModifier {
	return &CreatureModifier{creature: c}
}
// 3 Concrete Handler
type DoubleDamageModifier struct {
	CreatureModifier
}
func (d *DoubleDamageModifier) Apply() {
	fmt.Println("double damage rune applied for",d.creature.Name)
	d.creature.Attack *=2
	//!!!
	d.CreatureModifier.Apply()
}

func NewDoubleDamageModifier(c *Creature) *DoubleDamageModifier {
	return &DoubleDamageModifier{CreatureModifier{creature: c}}
}

type FrostShieldModifier struct {
	CreatureModifier
}
func (d *FrostShieldModifier) Apply() {
	if d.creature.Attack <=150 {
		fmt.Println("Frost shield was applied")
		d.creature.Defense +=20
	}
	d.CreatureModifier.Apply()
}
func NewFrostShieldModifier(c *Creature) *FrostShieldModifier {
	return &FrostShieldModifier{CreatureModifier{creature: c}}
}
type NoBonusesModifier struct {
}
func chain()  {
	Juggernaut := NewCreature("Jugger",61,8)
	fmt.Println(Juggernaut)
	root := NewCreatureModifier(Juggernaut)
	root.Add(NewDoubleDamageModifier(Juggernaut))
	root.Add(NewFrostShieldModifier(Juggernaut))
	root.Add(NewDoubleDamageModifier(Juggernaut))
	root.Add(NewFrostShieldModifier(Juggernaut))
	root.Apply()
	fmt.Println(Juggernaut)
}
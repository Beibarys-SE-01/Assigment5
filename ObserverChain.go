package main
/*
import (
	"fmt"
	"sync"
)
type Argument int
const (
	Attack Argument = iota
	Defense
)
type Query struct {
	CreatureName string
	WhatToQuery Argument
	Value int
}
// Handle interface
type Observer2 interface {
	Handle(*Query)
}
type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}
type Game struct {
	observers sync.Map
}
func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
	// ↑↑↑ empty anon struct
}
func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}
func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}
type Creature struct {
	game *Game
	Name string
	attack, defense int // ← private!
}
func (c *Creature) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	return q.Value
}
func (c *Creature) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)
	return q.Value
}
func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack(), c.Defense())
}
func NewCreature(game *Game, name string, attack int, defense int) *Creature {
	return &Creature{game: game, Name: name, attack: attack, defense: defense}
}
// Base Handler
type CreatureModifier struct {
	game *Game
	creature *Creature
}
func (c *CreatureModifier) Handle(*Query) {
	// nothing here!
}
// Concrete
type DoubleAttackModifier struct {
	CreatureModifier
}
func (d *DoubleAttackModifier) Handle(q *Query) {
	if q.CreatureName == d.creature.Name &&
		q.WhatToQuery == Attack {
		q.Value *= 2
	}
}
func (d *DoubleAttackModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}
func NewDoubleAttackModifier(g *Game, c *Creature) *DoubleAttackModifier {
	d := &DoubleAttackModifier{CreatureModifier{g, c}}
	g.Subscribe(d)
	return d
}

type WeakDefenseModifier struct {
	CreatureModifier
}
func (w *WeakDefenseModifier) Handle(q *Query) {
	if q.CreatureName == w.creature.Name &&
		q.WhatToQuery == Defense {
		q.Value /= 2
	}
}
func (w *WeakDefenseModifier) Close() error {
	w.game.Unsubscribe(w)
	return nil
}
func NewWeakDefenseModifier(g *Game, c *Creature) *WeakDefenseModifier {
	w := &WeakDefenseModifier{CreatureModifier{g, c}}
	g.Subscribe(w)
	return w
}
func main() {
	game := &Game{sync.Map{}}
	goblin := NewCreature(game, "Strong Goblin", 2, 2)
	fmt.Println(goblin)
	m := NewDoubleAttackModifier(game,goblin)
	w := NewWeakDefenseModifier(game, goblin)
	fmt.Println(goblin)
	m.Close()
	w.Close()
	fmt.Println(goblin)
}*/
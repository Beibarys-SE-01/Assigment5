package main

import "fmt"

type Subject interface {
	register(Observer)
	unregister(Observer)
	notifyAll()
}

type Cinema struct {
	subscribers []Observer
	movies []string
}

func (c *Cinema) register(observer Observer)  {
	c.subscribers = append(c.subscribers, observer)
}
func (c *Cinema) unregister(observer Observer)  {
	for i, sub := range c.subscribers{
		if sub == observer {
			c.subscribers = removeFromSlice(c.subscribers, i)
		}
	}
}

func removeFromSlice(s []Observer, i int) []Observer {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func (c *Cinema) notifyAll() {
	for _,observer := range c.subscribers {
		observer.update(c.movies)
	}
}

func (c *Cinema) addMovie(name string) {
	c.movies = append(c.movies, name)
	c.notifyAll()
}
func (c *Cinema) removeMovie(name string) {
	for i, nameOfmovie := range c.movies{
		if nameOfmovie == name {
			c.movies = remove(c.movies, i)
		}
	}
	c.notifyAll()
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

type Observer interface {
	update([]string)
	getID() int
}
type Person struct {
	id int
	Name string
}

func (p *Person) update(movies []string)  {
	fmt.Println("Hello,",p.Name)
	fmt.Println("We have new movies in our cinema")
	for i, movie := range movies{
		fmt.Println(i, movie)
	}
}
func (p *Person)getID() int {
	return p.id
}

func main() {
	chaplinCinema := Cinema{[]Observer{}, []string{}}
	chaplinCinema.addMovie("Sonic")

	beibarys := &Person{1, "Beibarys"}
	tomiris := &Person{2, "Tomiris"}
	bagdaulet := &Person{3, "Bagdaulet"}

	chaplinCinema.register(beibarys)
	chaplinCinema.register(tomiris)

	chaplinCinema.addMovie("Forest Gump")

	chaplinCinema.unregister(tomiris)

	chaplinCinema.addMovie("Naruto")

	chaplinCinema.removeMovie("Sonic")

	chaplinCinema.register(bagdaulet)
	chaplinCinema.unregister(beibarys)

	chaplinCinema.addMovie("Harry Potter")
}

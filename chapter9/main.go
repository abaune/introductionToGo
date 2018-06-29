package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

type Shape interface {
	area() float64
	perimeter() float64
}

type MultiShape struct {
	shapes []Shape
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	return r.x1 + r.x2 + r.y1 + r.y2
}

func (p *Person) Talk() {
	fmt.Println("Hi my name is", p.Name)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	// var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	// fmt.Println(circleArea(cx, cy, cr))

	c := Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)
	c.x = 10
	c.y = 5
	fmt.Println(circleArea(&c))
	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	a := new(Android)
	a.Person.Name = "Test"
	a.Person.Talk()
	a.Talk()

	fmt.Println(totalArea(&c, &r))
	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())
}

/*Problems
1. A method is a function within the scope of an object.
2. When we want to use and 'is a' relationship. For example. An Android is a person, not an android has a person.
3. see perimeter() method on shape interface
*/

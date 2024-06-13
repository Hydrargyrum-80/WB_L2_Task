package pattern

import "fmt"

/*
Фабричный метод – это способ создавать объекты каких-то неизвестных нам структур, про которые знает только сам фабричный метод.
Может применяться в ситуациях, когда необходимо единообразно работать с некоторым набором "объектов"
*/

type Shape interface {
	draw()
}

type Circle struct{}

func (c *Circle) draw() {
	fmt.Println("Circle")
}

type Rectangle struct{}

func (p *Rectangle) draw() {
	fmt.Println("Rectangle")
}

type FactoryMethod interface {
	createShapes(shapeName string) Shape
}

type ShapesFabric struct{}

func (c *ShapesFabric) createShapes(shapeName string) Shape {
	switch shapeName {
	case "Rectangle":
		return &Rectangle{}
	case "Circle":
		return &Circle{}
	default:
		return nil
	}
}

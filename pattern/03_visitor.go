package pattern

import "fmt"

/*
Visitor позволяет добавлять новые методы для иерархии структур, не изменяя при этом "абстрактного родителя".
Недостатком будет повышенная сложность программы, а так же возможность возможность распределения методов структур по
разным частям программы.
*/

type Visitor interface {
	handlePoint(p *VPoint)
	handleGroup(g *VGroup)
}

type VShape interface {
	Apply(handler Visitor)
}

type VPoint struct {
	X, Y int
}

func (p *VPoint) Apply(handler Visitor) {
	handler.handlePoint(p)
}

type VGroup struct {
	shapes []VShape
}

func (g *VGroup) Apply(handler Visitor) {
	handler.handleGroup(g)
}

type Drawer struct{}

func (d *Drawer) handlePoint(p *VPoint) {
	fmt.Println("***VPoint ", p.X, ":", p.Y, "***")
}

func (d *Drawer) handleGroup(p *VGroup) {
	fmt.Println("***VGroup: ")
	for _, i := range p.shapes {
		i.Apply(d)
	}
	fmt.Println("***")
}

type Mover struct{}

func (m *Mover) handlePoint(p *VPoint) {
	p.X++
	p.Y++
}

func (m *Mover) handleGroup(p *VGroup) {
	for _, i := range p.shapes {
		i.Apply(m)
	}
}

/*
func main() {
	shapes := make([]VShape, 0, 4)
	shapes = append(shapes, &VPoint{1, 2})
	shapes = append(shapes, &VPoint{3, 4})
	shapes = append(shapes, &VPoint{5, 6})
	shapes = append(shapes, &VGroup{shapes})
	for _, i := range shapes {
		i.Apply(&Mover{})
	}
	for _, i := range shapes {
		i.Apply(&Drawer{})
	}
}
*/

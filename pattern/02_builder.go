package pattern

/*
Builder позволяет не зависеть от того из каких частей состоит создаваемый объект, так же позволяя получить
различные его вариации в процессе создания
*/

type Builder interface {
	BuildPartA() int
	BuildPartB() float32
	BuildPartC() string
}

type ConcreteBuilder struct{}

func (b *ConcreteBuilder) BuildPartA() int {
	return 1
}

func (b *ConcreteBuilder) BuildPartB() float32 {
	return 1.1
}

func (b *ConcreteBuilder) BuildPartC() string {
	return "111"
}

type Creator struct {
	A int
	B float32
	C string
}

func (c *Creator) Create(builder Builder) {
	c.A = builder.BuildPartA()
	c.B = builder.BuildPartB()
	c.C = builder.BuildPartC()
}

/*
func main() {
	creator := &Creator{}
	builder := ConcreteBuilder{}
	creator.Create(&builder)
	fmt.Println(creator)
}
*/

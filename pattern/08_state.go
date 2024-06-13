package pattern

import "fmt"

/*
State позволяет объекту изменять свое поведение в зависимости от внутреннего состояния.
Например, действие при нажатии на кнопку телефона в зависимости от заряда батареи может приводит к экрану разблокировки/зарядки и т.д
*/

type State interface {
	Handle(context Context)
}

type Context struct {
	state State
}

func (c *Context) Request() {
	c.state.Handle(*c)
}

type StateA struct {
}

func (S *StateA) Handle(context Context) {
	context.state = &StateB{}
	fmt.Println("StateA -> StateB")
}

type StateB struct {
}

func (S *StateB) Handle(context Context) {
	context.state = &StateA{}
	fmt.Println("StateB -> StateA")
}

/*
func main() {
	context := &Context{state: &StateA{}}
	context.Request()
	context.Request()
}
*/

package pattern

import "fmt"

/*
Facade позволяет скрыть сложность системы с помощью предоставления упрощенного интерфейса для взаимодействия с ней.
Используется, например, в IDE: приложение позволяет писать код, скрывая при этом сложность компиляции и запуска программы.
*/

type SubsystemA struct{}

type SubsystemB struct{}

type SubsystemC struct{}

func (s *SubsystemA) SubA() {
	fmt.Println("A")
}

func (s *SubsystemB) SubB() {
	fmt.Println("B")
}

func (s *SubsystemC) SubC() {
	fmt.Println("C")
}

type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
	subsystemC *SubsystemC
}

func NewFacade(sA *SubsystemA, sB *SubsystemB, sC *SubsystemC) *Facade {
	return &Facade{sA, sB, sC}
}

/*
func main() {
	facade := NewFacade(&SubsystemA{}, &SubsystemB{}, &SubsystemC{})
	facade.subsystemB.SubB()
	facade.subsystemA.SubA()
	facade.subsystemC.SubC()
}
*/

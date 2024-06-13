package pattern

/*
Chain of Responsibility позволяет разорвать жёсткую связь между отправителем запроса и его исполнителем.
Цепочка исполнителей может формироваться "на лету" и перестраиваться в процессе работы программы.
Может использоваться, например, когда между вызовом метода бизнес-логики должно выполняться ещё какое-то действие
*/

type Handler interface {
	handle(request int) bool
}

type DivisionChecker struct {
	next Handler
	val  int
}

func NewDivisionChecker(next Handler, val int) *DivisionChecker {
	return &DivisionChecker{next: next, val: val}
}

func (dv *DivisionChecker) handle(request int) bool {
	if dv.val == request {
		return true
	}
	if dv.next == nil {
		return false
	} else {
		return (dv.next).handle(request)
	}
}

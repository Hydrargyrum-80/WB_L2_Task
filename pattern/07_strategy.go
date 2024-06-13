package pattern

/*
Основная идея этого паттерна заключается в том, чтобы инкапсулироать, «обернуть», «вынести» разные возможные алгоритмы какого-то сложного действия –
в отдельные объекты, чтобы потом иметь возможность изменять алгоритм поведения путём просто замены объекта.
Cтратегия перемещения в пространстве игры — игрок ходит, либо бегает, но, возможно, в будущем он также сможет плавать, летать, телепортироваться, рыть под землей и др.
*/

type MultStrategy interface {
	multiply(arr *[]int, cnt int) int
}

type ForwardMultStrategy struct{}

func (f *ForwardMultStrategy) multiply(arr *[]int, cnt int) int {
	var result int = 1
	for i := 0; i < cnt; i++ {
		result *= (*arr)[i]
	}
	return result
}

type BackwardMultStrategy struct{}

func (b *BackwardMultStrategy) multiply(arr *[]int, cnt int) int {
	var result int = 1
	for i := len(*arr) - 1; i > cnt; i-- {
		result *= (*arr)[i]
	}
	return result
}

/*
func main() {
	arr := make([]int, 5)
	for i := 0; i < cap(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	var st MultStrategy = &ForwardMultStrategy{}
	fmt.Println(st.multiply(&arr, 3))
	st = &BackwardMultStrategy{}
	fmt.Println(st.multiply(&arr, 2))
}
*/

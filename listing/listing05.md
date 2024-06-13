Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

```
Почти аналогично listing03: интерфейс представляет из себя структуру, имеющую поля tab, data.
В test возвращается указатель на customError с значением nil. Аналогично tab = nil и data = nil не выполнено
(т.к data=nil но tab!=nil) => и ошибка !=nil, которой вроде как и нет.
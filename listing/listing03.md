Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
nil
false

```
Сам по себе интерфейс представляет из себя структуру, имеющую поля tab, data. 
В методе Foo инициализация ошибки определенного типа PathError = nil, однако теперь в этом поле будет храниться 
tab равный указателю на PathError !=nil, => при сравнении с nil условие tab = nil и data = nil не выполнено.
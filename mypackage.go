/*
来点注释
*/
package mypackage

import (
	"fmt"
)

//人物
type Person struct {
	Name string
}

//say hello
func (p *Person) SayHello(toPerson *Person) {
	fmt.Println(p.Name, ": Hello, ", toPerson.Name)
}

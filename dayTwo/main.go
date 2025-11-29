// EchoTwo вывщодит все аргументы командной строки
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "__"
	}
	echo3()
	IndexwithValue()
	fmt.Println(s)
}

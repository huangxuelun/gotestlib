package gotestlib

import "fmt"

type Message struct {
	Title   string
	Content string
	Footer  string
}

func (this Message) Hello() {
	fmt.Println(this)
}

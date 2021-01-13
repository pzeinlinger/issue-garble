package main

import (
	"fmt"

	rice "github.com/GeertJohan/go.rice"
)

const (
	token = "asdfasdfasdfasdfasdfasdfasfdasdfasfdasdf"
)

func main() {
	fmt.Println(Somestruct{}.SayHello())
	fmt.Println(token)

	rice.MustFindBox("box")
	rice.MustFindBox("pdfjs")
	rice.MustFindBox("assets")
}

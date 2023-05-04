package main

import "fmt"

// Server to init server object
type Server struct {
	ID                    string
	Name                  string
	SomeLongNameField     string
	SomeVertLongNameField string
}

func main() {
	fmt.Println("add")
	fmt.Println("some fix in the server")
}

package main

import "fmt"

func main() {
	a := App{}

	a.Initialize()
	fmt.Println("All Connected........")
	a.StartServer(":8000")
}

package main

import (
	"fmt"
	"gowebstarter/adapter"
)

func main() {
	r := adapter.InitGin()
	fmt.Println("Listening on port 8080...")
	r.Run(":8080")

}

package main

import "fmt"

func escapeFunction() func() func() {
	return func() func() {
		return func() {
			fmt.Println("hello closure function")
		}
	}
}

var (
	_ = escapeFunction
)

func main() {

}

package main

func closure() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}

func main() {
	closure()()
}

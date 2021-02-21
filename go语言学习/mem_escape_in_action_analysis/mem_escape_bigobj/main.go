package main

func bigObj() []int {
	return make([]int, 10000+1)
}

func main() {
	_ = bigObj()
}

# go 语言内存逃逸场景分析

重点：
```text
// go 语言内存逃逸分析
// 1. 分配在栈上的空间不会逃逸, 函数结束后回收
// 2. 堆上分配的内存使用完毕之后交给 GC 处理
// 3. 逃逸分析在编译阶段完成， 目的是决定对象分配到**堆**还是**栈**
```

## 闭包造成的逃逸
原本在函数运行栈空间上分配的内存，由于闭包的关系，变量在函数的作用域之外使用
```go
package main 

func closure() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}
```

## 返回指向栈变量的指针
返回的变量是栈对象的指针，编译器认为该对象在函数结束之后还需要使用
```go
package main

type Empty struct {}

func Demo() *Empty  {
	return &Empty{}
}

func main() {
	
}

```

## 申请大对象造成的逃逸
申请的空间过大， 也会直接在堆上分配空间
```go
package main

func bigObj() []int {
	return make([]int, 10000+1)
}

func main() {
	_ = bigObj()
}

```

## 申请可变长空间造成的逃逸
编译器无法知道需要分配多大的空间， 由于 ln 是可变的变量
```go
package main

func allocateVar(ln int) []int {

	return make([]int, ln)
}

func main() {

}

```

## 返回局部引用的 slice 造成的逃逸
编译器认为返回的对象可能会在函数调用完成之后， 还会再次使用
所以， 编译器会在堆上分配内存空间
```go
package main

func returnSlice() [] int {
	return make([]int,10)
}


func main() {
	_ = make(map[string]struct{})
}

```

## 返回局部引用的 map 造成的逃逸
编译器认为返回的对象可能会在函数调用完成之后， 还会再次使用
所以， 编译器会在堆上分配内存空间
```go
package main

func returnMap() map[string] struct {} {
	return  make(map[string]struct{})
}

func main() {

}

```
## 返回函数造成的逃逸

```go
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
```

## 总结
1. 闭包造成的逃逸
2. 返回指向栈变量的指针
3. 申请大对象造成的逃逸
4. 申请可变长空间造成的逃逸
5. 返回局部引用的 slice 造成的逃逸
6. 返回局部引用的 map 造成的逃逸
7. 返回函数造成的逃逸

----
1. 逃逸是在编译期间完成的， 主要是决定是在栈中或者堆中分配内存
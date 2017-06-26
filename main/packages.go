package main

import (
	"fmt"
)

// var c, python, java bool

// type Vertex struct {
// 	X int
// 	Y int
// }

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)

	// fmt.Println("My favorite number is", rand.Intn(10))
	// fmt.Println(add(1, 2))
	// a, b := swap("hello", "world")
	// fmt.Println(split(100))
	// fmt.Println(a, b)
	// var i int
	// fmt.Println(i, c, python, java)

	// var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	// var m uint
	// fmt.Println(x, y, z, m)

	// v := 1 // 修改这里！
	// fmt.Printf("v is of type %T\n", v)

	// fmt.Println("start")
	// for i := 1; i <= 100; i++ {
	// 	defer fmt.Println(i)
	// }
	// fmt.Println("done")

	// i, j := 42, 2701

	// p := &i         // point to i
	// fmt.Println(*p) // read i through the pointer
	// *p = 21         // set i through the pointer
	// fmt.Println(i)  // see the new value of i

	// p = &j // point to j
	// fmt.Println(*p)
	// *p = *p / 37   // divide j through the pointer
	// fmt.Println(j) // see the new value of j

	// v := Vertex{1, 2}
	// p := &v
	// p.X = 1e9
	// v.X = 1e9
	// fmt.Println(v)

	// defer func() {
	// 	fmt.Println("c")
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("d")
	// }()
	// panic("ex")

}

// func f() {
// 	fmt.Println("a")
// 	panic("ex")
// 	fmt.Println("b")
// 	fmt.Println("f")
// }

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {

	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 6 / 2
	y = sum - 1
	return
}

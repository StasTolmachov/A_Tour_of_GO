// package main
//
// import (
//
//	"fmt"
//	"math/cmplx"
//
// )
//
// var (
//
//	ToBe   bool       = false
//	MaxInt uint64     = 1<<64 - 1
//	z      complex128 = cmplx.Sqrt(-5 + 12i)
//
// )
//
//	func main() {
//		fmt.Printf("Тип: %T Значення: %v\n", ToBe, ToBe)
//		fmt.Printf("Тип: %T Значення: %v\n", MaxInt, MaxInt)
//		fmt.Printf("Тип: %T Значення: %v\n", z, z)
//	}
//
// package main
//
// import (
//
//	"fmt"
//	"strconv"
//
// )
//
//	func main() {
//		a := 3
//		fmt.Printf("%v %T\n", a, a)
//
//		b := strconv.Itoa(a)
//
//		fmt.Printf("%s %T\n", b, b)
//		c, _ := strconv.Atoi(b)
//
//		fmt.Printf("%v %T\n", c, c)
//
// }
//
// package main
//
// import "fmt"
//
//	func main() {
//		var number int
//		number = 1
//		a := number << 1
//		b := number << 2
//		c := number << 3
//		d := number << 4
//
//		fmt.Printf("Десятичная: %d %d %d %d\n", a, b, c, d)
//		fmt.Printf("Двоичная: %b %b %b %b\n", a, b, c, d)
//	}
//
// package main
//
// import (
//
//	"fmt"
//	"math"
//
// )
//
//	func Sqrt(x float64) float64 {
//		z := 1.0           // Початкове наближення
//		threshold := 1e-10 // Поріг для зупинки
//
//		for {
//			newZ := (z + x/z) / 2
//			if math.Abs(newZ-z) < threshold {
//				break
//			}
//			z = newZ
//			fmt.Printf("z = %v\n", z)
//		}
//		return z
//	}
//
//	func main() {
//		fmt.Printf("Sqrt(1): %v (math.Sqrt: %v)\n", Sqrt(1), math.Sqrt(1))
//		fmt.Printf("Sqrt(2): %v (math.Sqrt: %v)\n", Sqrt(2), math.Sqrt(2))
//		fmt.Printf("Sqrt(3): %v (math.Sqrt: %v)\n", Sqrt(3), math.Sqrt(3))
//		fmt.Printf("Sqrt(4): %v (math.Sqrt: %v)\n", Sqrt(4), math.Sqrt(9))
//	}
//
// package main
//
// import (
//
//	"fmt"
//	"runtime"
//
// )
//
//	func main() {
//		fmt.Print("Go запущено на ")
//		switch os := runtime.GOOS; os {
//		case "darwin":
//			fmt.Println("OS X.")
//		case "linux":
//			fmt.Println("Linux.")
//		default:
//			// freebsd, openbsd,
//			// plan9, windows...
//			fmt.Printf("%s.\n", os)
//		}
//	}
//
// package main
//
// import (
//
//	"fmt"
//	"time"
//
// )
//
//	func main() {
//		fmt.Println("Коли Субота?")
//		today := time.Now().Weekday()
//		switch time.Saturday {
//		case today + 0:
//			fmt.Println("Сьогодні.")
//		case today + 1:
//			fmt.Println("Завтра.")
//		case today + 2:
//			fmt.Println("Через два дні.")
//		default:
//			fmt.Println("Колись буде..")
//		}
//	}
//
// package main
//
// import (
//
//	"fmt"
//	"time"
//
// )
//
//	func main() {
//		t := time.Now()
//		switch {
//		case t.Hour() < 12:
//			fmt.Println("Доброго ранку!")
//		case t.Hour() < 17:
//			fmt.Println("Добрий день.")
//		default:
//			fmt.Println("Добрий вечір.")
//		}
//	}
//
// package main
//
// import "fmt"
//
//	func main() {
//		fmt.Println("1")
//		defer fmt.Println("2")
//		fmt.Println("3")
//		defer fmt.Println("4")
//		fmt.Println("5")
//
// }
//
// package main
//
// import "fmt"
//
//	func main() {
//		var a, b = 2, 3
//		fmt.Println("a = ", a, "\nb = ", b)
//
//		p := &a
//		fmt.Println("p := &a:", p, "\n*p:", *p)
//		fmt.Println("____________")
//
//		fmt.Println("p = ", p)   // p =  0xc000094018
//		fmt.Println("&p = ", &p) // &p =  0xc0000a2020
//		fmt.Println("*p = ", *p) // *p =  2
//
// }
//
// package main
//
// import "fmt"
//
//	type Vertex struct {
//		X int
//		Y int
//	}
//
//	func main() {
//		var v Vertex
//
//		v.X = 4
//		fmt.Println(v.X)
//
//		p := &v
//
//		fmt.Println(*p)
//		p.Y = 5
//		fmt.Println(*p)
//
// }
//
// package main
//
// import "fmt"
//
//	type Vertex struct {
//		X, Y int
//	}
//
// var (
//
//	v1 = Vertex{1, 2}  // має тип Vertex
//	v2 = Vertex{X: 1}  // Y:0 є неявним
//	v3 = Vertex{}      // X:0 та Y:0
//	p  = &Vertex{1, 2} // має тип *Vertex
//
// )
//
//	func main() {
//		fmt.Println(v1, p, v2, v3)
//	}
//
// package main
//
// import "fmt"
//
//	func main() {
//		var a [2]string
//		a[0] = "Привіт"
//		a[1] = "Світ"
//		fmt.Println(a[0], a[1])
//		fmt.Println(a)
//
//		primes := [6]int{2, 3, 5, 7, 11, 13}
//		fmt.Println(primes)
//	}
//
// package main
//
// import "fmt"
//
//	func main() {
//		a := [5]int{1, 2, 3, 4, 5}
//		fmt.Println("arr: ", a)
//		s := a[2:4]
//		fmt.Println(s)
//	}
//
// package main
//
// import "fmt"
//
//	func main() {
//		q := []int{2, 3, 5, 7, 11, 13}
//		fmt.Println(q)
//
//		r := []bool{true, false, true, true, false, true}
//		fmt.Println(r)
//
//		s := []struct {
//			i int
//			b bool
//		}{
//			{2, true},
//			{3, false},
//			{5, true},
//			{7, true},
//			{11, false},
//			{13, true},
//		}
//		fmt.Println(s)
//	}
//
// package main
//
// import (
//
//	"fmt"
//	"reflect"
//	"unsafe"
//
// )
//
//	func main() {
//		// Создаем срез с помощью make
//		s2 := make([]int, 5)
//		fmt.Println(s2)
//		fmt.Println("Длина:", len(s2), "Емкость:", cap(s2))
//		fmt.Println("__________________________________________________")
//
//		// Печать указателя на срез
//		fmt.Println("Указатель на срез:", &s2)
//
//		// Используем unsafe.Pointer и reflect.SliceHeader для получения базового массива
//		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&s2))
//		baseArrayPtr := unsafe.Pointer(sliceHeader.Data)
//		fmt.Println("Указатель на базовый массив:", baseArrayPtr)
//
//		// Печать адресов элементов среза для сравнения
//		for i := 0; i < len(s2); i++ {
//			fmt.Printf("Адрес элемента %d: %p\n", i, &s2[i])
//		}
//		/*
//		   [0 0 0 0 0]
//		   	Длина: 5 Емкость: 5
//		   	__________________________________________________
//		   	Указатель на срез: 0xc000004078
//		   	Указатель на базовый массив: 0xc00000a0a0
//		   	Адрес элемента 0: 0xc00000a0a0
//		   	Адрес элемента 1: 0xc00000a0a8
//		   	Адрес элемента 2: 0xc00000a0b0
//		   	Адрес элемента 3: 0xc00000a0b8
//		   	Адрес элемента 4: 0xc00000a0c0
//		*/
//	}
//
// package main
//
// import "fmt"
//
//	func main() {
//		s := [][]int{
//			[]int{1, 2, 3},
//			[]int{4, 5, 6},
//			[]int{7, 8, 9},
//		}
//		fmt.Print(s[0], "\n", s[1], "\n", s[2], "\n")
//		fmt.Println("__________________________________________________")
//
//		var s2 []int
//		fmt.Println(s2)
//		if s2 == nil {
//			fmt.Println("s2 == nil")
//		}
//		fmt.Println("__________________________________________________")
//		s2 = append(s2, 0)
//		fmt.Println(s2)
//		if s2 == nil {
//			fmt.Println("s2 == nil")
//		}
//	}
package main

import (
	"golang.org/x/tour/pic"
)

// Pic функция генерирует изображение
func Pic(dx, dy int) [][]uint8 {
	// Создаем двумерный срез
	image := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		// Создаем срез для каждой строки
		image[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// Используем одну из функций для генерации значения
			// image[y][x] = uint8((x + y) / 2)
			// image[y][x] = uint8(x * y)
			image[y][x] = uint8(x ^ y)
		}
	}
	return image
}

func main() {
	// Используем функцию Pic из пакета tour для отображения изображения
	pic.Show(Pic)
}

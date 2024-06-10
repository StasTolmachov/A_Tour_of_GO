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
//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func main() {
//	fmt.Print("Go запущено на ")
//	switch os := runtime.GOOS; os {
//	case "darwin":
//		fmt.Println("OS X.")
//	case "linux":
//		fmt.Println("Linux.")
//	default:
//		// freebsd, openbsd,
//		// plan9, windows...
//		fmt.Printf("%s.\n", os)
//	}
//}
//
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	fmt.Println("Коли Субота?")
//	today := time.Now().Weekday()
//	switch time.Saturday {
//	case today + 0:
//		fmt.Println("Сьогодні.")
//	case today + 1:
//		fmt.Println("Завтра.")
//	case today + 2:
//		fmt.Println("Через два дні.")
//	default:
//		fmt.Println("Колись буде..")
//	}
//}

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Доброго ранку!")
	case t.Hour() < 17:
		fmt.Println("Добрий день.")
	default:
		fmt.Println("Добрий вечір.")
	}
}

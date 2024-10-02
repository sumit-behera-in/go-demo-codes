//go build -o www.exe "c:\Users\Sumit\Desktop\Go-cli-demo\output_input_of_other"
// change .exe name

package main

import (
	"fmt"
	"sync"
	"time"
)

type balance interface {
	Date() int
	Month() int
	Year() int
}

type Balance struct {
	date Dob
}

//"strings"

// type People struct {
// 	firstName   string
// 	lastname    string
// 	dateofBirth Dob
// }

type Dob struct {
	date  int
	month int
	year  int
}

func (d Dob) get() {
	fmt.Println(d.date, "/", d.month, "/", d.year)
}

var wg = sync.WaitGroup{}

func main() {

	// a := 1

	var a = make(chan int, 100)

	date := Dob{
		date:  15,
		month: 8,
		year:  2001,
	}

	date.get()

	go func() {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go remain(a, i)
			// remain2(&a, 33)
		}
		wg.Wait()
		close(a)
	}()

	for n := range a {
		fmt.Println(n)
	}
	// var name int

	// // var arr = []int{10, 20, 30}

	// // arr = append(arr, 77)

	// // for index, data := range arr {
	// // 	fmt.Println(index, data)
	// // }

	// var mp = make(map[int]Dob)

	// for {
	// 	fmt.Print("Enter the Num : ")
	// 	fmt.Scan(&name)
	// 	if name == -2 {
	// 		break
	// 	}

	// 	// mp[name] = Op(name)

	// 	// switch name {
	// 	// case 1:
	// 	// 	fmt.Println("One")
	// 	// case 0:
	// 	// 	fmt.Println("Zero")
	// 	// default:
	// 	// 	fmt.Println(name, Op(name))
	// 	// }

	// 	// const ab = 55
	// 	// fmt.Println("Hello World")
	// 	// fmt.Println(name, ab)

	// 	mp[name] = Dob{
	// 		date:  name,
	// 		month: Op(name),
	// 		year:  Op(name * 6),
	// 	}

	// }

	// for a, b := range mp {
	// 	fmt.Println(a, "  ", b)
	// }

	// fmt.Println(arr[0])

}

// func Op(name int) int {
// 	return name * 100
// }

func remain(a chan int, i int) {
	a <- i % 2
	println("thread ", i, " value", a)
	time.Sleep(5 * time.Second)
}

func (d Dob) Date() int {
	return d.date
}

func (d Dob) Month() int {
	return d.month
}

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var a, b, s string
	var x, y int
	var result1 int
	var _ error
	var i1 string = "I"
	var i2 string = "II"
	var i3 string = "III"
	var i4 string = "IV"
	var i5 string = "V"
	var i6 string = "VI"
	var i7 string = "VII"
	var i8 string = "VIII"
	var i9 string = "IX"
	var i10 string = "X"

	fmt.Println("Введите выражение ")
	fmt.Scan(&a, &s, &b)

	if a != i1 || a != i2 || a != i3 || a != i4 || a != i5 || a != i6 || a != i7 || a != i8 || a != i9 || a != i10 {
		x, _ := strconv.Atoi(a)
		y, _ := strconv.Atoi(b)
		if s == "+" {
			result1 = x + y
			fmt.Println(result1)
		} else if s == "-" {
			result1 = x - y
			fmt.Println(result1)
		} else if s == "*" {
			result1 = x * y
			fmt.Println(result1)
		} else if s == "/" {
			result1 = x / y
			fmt.Println(result1)
		} else {
			fmt.Println("Выражение не является математической операцией")
		}
	}

	if reflect.TypeOf(a).String() == "string" && reflect.TypeOf(b).String() == "string" {
		if a == i1 || a == i2 || a == i3 || a == i4 || a == i5 || a == i6 || a == i7 || a == i8 || a == i9 || a == i10 {
			if b == i1 || b == i2 || b == i3 || b == i4 || b == i5 || b == i6 || b == i7 || b == i8 || b == i9 || b == i10 {
				switch a {
				case "I":
					x = 1
				case "II":
					x = 2
				case "III":
					x = 3
				case "IV":
					x = 4
				case "V":
					x = 5
				case "VI":
					x = 6
				case "VII":
					x = 7
				case "VIII":
					x = 8
				case "IX":
					x = 9
				case "X":
					x = 10
				default:
					fmt.Println("Takoe chislo ne podderzhivaetsya")
				}

				switch b {
				case "I":
					y = 1
				case "II":
					y = 2
				case "III":
					y = 3
				case "IV":
					y = 4
				case "V":
					y = 5
				case "VI":
					y = 6
				case "VII":
					y = 7
				case "VIII":
					y = 8
				case "IX":
					y = 9
				case "X":
					y = 10
				default:
					fmt.Println("Такое число не поддерживается")
				}
				if s == "+" {
					result1 = x + y
				} else if s == "-" {
					result1 = x - y
					if result1 == 0 {
						fmt.Println("В римскоим исчеслении нет числа 0")
					}
				} else if s == "*" {
					result1 = x * y
				} else if s == "/" {
					result1 = x / y
				} else {
					fmt.Println("Выражение не является математической операцией")
				}
				switch result1 {
				case 1:
					fmt.Println("I")
				case 2:
					fmt.Println("II")
				case 3:
					fmt.Println("III")
				case 4:
					fmt.Println("IV")
				case 5:
					fmt.Println("V")
				case 6:
					fmt.Println("VI")
				case 7:
					fmt.Println("VII")
				case 8:
					fmt.Println("VIII")
				case 9:
					fmt.Println("IX")
				case 10:
					fmt.Println("X")
				default:
					fmt.Println("Такое число не поддерживается")

				}
			}
		}

	}
}

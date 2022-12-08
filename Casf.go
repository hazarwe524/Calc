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

	if len(a) > 4 && len(s) > 1 && len(b) > 4 {
		fmt.Println("Несоотвествующее выражение")
	}

	if (a != i1 || a != i2 || a != i3 || a != i4 || a != i5 || a != i6 || a != i7 || a != i8 || a != i9 || a != i10) &&
		(b != i1 || b != i2 || b != i3 || b != i4 || b != i5 || b != i6 || b != i7 || b != i8 || b != i9 || b != i10) {
		x, _ := strconv.Atoi(a)
		y, _ := strconv.Atoi(b)
		if x >= 1 && x <= 10 && y >= 1 && y <= 10 {
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
	} else {
		fmt.Println("Некорректное выражение")
	}

	if reflect.TypeOf(a).String() == "string" && reflect.TypeOf(b).String() == "string" {
		if (a == i1 || a == i2 || a == i3 || a == i4 || a == i5 || a == i6 || a == i7 || a == i8 || a == i9 || a == i10) && (reflect.TypeOf(a).String() != "int") {
			if (b == i1 || b == i2 || b == i3 || b == i4 || b == i5 || b == i6 || b == i7 || b == i8 || b == i9 || b == i10) && (reflect.TypeOf(a).String() != "int ") {
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
				case 11:
					fmt.Println("XI")
				case 12:
					fmt.Println("XII")
				case 13:
					fmt.Println("XIII")
				case 14:
					fmt.Println("XIV")
				case 15:
					fmt.Println("XV")
				case 16:
					fmt.Println("XVI")
				case 17:
					fmt.Println("XVII")
				case 18:
					fmt.Println("XVIII")
				case 19:
					fmt.Println("XIX")
				case 20:
					fmt.Println("XX")
				case 21:
					fmt.Println("XXI")
				case 22:
					fmt.Println("XXII")
				case 23:
					fmt.Println("XXIII")
				case 24:
					fmt.Println("XXIV")
				case 25:
					fmt.Println("XXV")
				case 26:
					fmt.Println("XXVI")
				case 27:
					fmt.Println("XXVII")
				case 28:
					fmt.Println("XXVIII")
				case 29:
					fmt.Println("XXIX")
				case 30:
					fmt.Println("XXX")
				case 31:
					fmt.Println("XXXI")
				case 32:
					fmt.Println("XXXII")
				case 33:
					fmt.Println("XXXIII")
				case 34:
					fmt.Println("XXXIV")
				case 35:
					fmt.Println("XXXV")
				case 36:
					fmt.Println("XXXVI")
				case 37:
					fmt.Println("XXXVII")
				case 38:
					fmt.Println("XXXVIII")
				case 39:
					fmt.Println("XXXIX")
				case 40:
					fmt.Println("XL")
				case 41:
					fmt.Println("XLI")
				case 42:
					fmt.Println("XLII")
				case 43:
					fmt.Println("XLIII")
				case 44:
					fmt.Println("XLIV")
				case 45:
					fmt.Println("XLV")
				case 46:
					fmt.Println("XLVI")
				case 47:
					fmt.Println("XLVII")
				case 48:
					fmt.Println("XLVIII")
				case 49:
					fmt.Println("XLIX")
				case 50:
					fmt.Println("LI")
				case 51:
					fmt.Println("LI")
				case 52:
					fmt.Println("LII")
				case 53:
					fmt.Println("LIII")
				case 54:
					fmt.Println("LIV")
				case 55:
					fmt.Println("LV")
				case 56:
					fmt.Println("LVI")
				case 57:
					fmt.Println("LVII")
				case 58:
					fmt.Println("LVIII")
				case 59:
					fmt.Println("LIX")
				case 60:
					fmt.Println("LX")
				case 61:
					fmt.Println("LXI")
				case 62:
					fmt.Println("LXII")
				case 63:
					fmt.Println("LXIII")
				case 64:
					fmt.Println("LXIV")
				case 65:
					fmt.Println("LXV")
				case 66:
					fmt.Println("LXVI")
				case 67:
					fmt.Println("LXVII")
				case 68:
					fmt.Println("LXVIII")
				case 69:
					fmt.Println("LXIX")
				case 70:
					fmt.Println("LXX")
				case 71:
					fmt.Println("LXXI")
				case 72:
					fmt.Println("LXXII")
				case 73:
					fmt.Println("LXXIII")
				case 74:
					fmt.Println("LXXIV")
				case 75:
					fmt.Println("LXXV")
				case 76:
					fmt.Println("LXXVI")
				case 77:
					fmt.Println("LXXVII")
				case 78:
					fmt.Println("LXXVIII")
				case 79:
					fmt.Println("LXXIX")
				case 80:
					fmt.Println("LXXX")
				case 81:
					fmt.Println("LXXXI")
				case 82:
					fmt.Println("LXXXII")
				case 83:
					fmt.Println("LXXXIII")
				case 84:
					fmt.Println("LXXXIV")
				case 85:
					fmt.Println("LXXXV")
				case 86:
					fmt.Println("LXXXVI")
				case 87:
					fmt.Println("LXXXVII")
				case 88:
					fmt.Println("LXXXVIII")
				case 89:
					fmt.Println("LXXXIX")
				case 90:
					fmt.Println("XC")
				case 91:
					fmt.Println("XCI")
				case 92:
					fmt.Println("XCII")
				case 93:
					fmt.Println("XCIII")
				case 94:
					fmt.Println("XCIV")
				case 95:
					fmt.Println("XCV")
				case 96:
					fmt.Println("XCVI")
				case 97:
					fmt.Println("XCVII")
				case 98:
					fmt.Println("XCVIII")
				case 99:
					fmt.Println("XCIX")
				case 100:
					fmt.Println("C")
				default:
					fmt.Println("Такое число не поддерживается")

				}
			}
		}

	}
}

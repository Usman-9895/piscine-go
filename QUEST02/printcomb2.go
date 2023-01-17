package main

import "github.com/01-edu/z01"

func PrintComb2() {
    for i := 0; i <= 98; i++ {
        for j := i + 1; j <= 99; j++ {
            /* if i < 10 {
                z01.PrintRune('0')
            } */
            z01.PrintRune(rune(i/10) + '0')
            z01.PrintRune(rune(i%10) + '0')
            z01.PrintRune(' ')
            /* if j < 10 {
                z01.PrintRune('0')
            } */
            z01.PrintRune(rune(j/10) + '0')
            z01.PrintRune(rune(j%10) + '0')
            if i != 98 || j != 99 {
                z01.PrintRune(',')
                z01.PrintRune(' ')
            }
        }
    }
    z01.PrintRune('\n')
}

func main() {
	PrintComb2()
}

//00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99$
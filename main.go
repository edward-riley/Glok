package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	zero := [5]string{
		" 0000 ",
		"00  00",
		"00  00",
		"00  00",
		" 0000 ",
	}
	one := [5]string{
		"1111  ",
		"  11  ",
		"  11  ",
		"  11  ",
		"111111",
	}

	two := [5]string{
		" 2222 ",
		"22  22",
		"   22 ",
		"  22  ",
		"222222",
	}
	three := [5]string{
		" 3333 ",
		"33  33",
		"   333",
		"33  33",
		" 3333 ",
	}
	four := [5]string{
		"44  44",
		"44  44",
		"444444",
		"    44",
		"    44",
	}
	five := [5]string{
		"555555",
		"55    ",
		"55555 ",
		"    55",
		"55555 ",
	}
	six := [5]string{
		" 6666 ",
		"66    ",
		"66666 ",
		"66  66",
		" 6666 ",
	}
	seven := [5]string{
		"777777",
		"   77 ",
		"  77  ",
		" 77   ",
		"77    ",
	}
	eight := [5]string{
		" 8888 ",
		"88  88",
		" 8888 ",
		"88  88",
		" 8888 ",
	}
	nine := [5]string{
		" 9999 ",
		"99  99",
		" 99999",
		"    99",
		" 9999 ",
	}
	twodots := [5]string{
		"     ",
		"  ▀  ",
		"     ",
		"  ▀  ",
		"     ",
	}

	digits := [...][5]string{zero, one, two, three, four, five, six, seven, eight, nine}
	screen.Clear()
	for {
		screen.Clear()
		screen.MoveTopLeft()
		hour := time.Now().Hour()
		minute := time.Now().Minute()
		second := time.Now().Second()

		clock := [...][5]string{
			digits[hour/10], digits[hour%10],
			twodots,
			digits[minute/10], digits[minute%10],
			twodots,
			digits[second/10], digits[second%10],
		}

		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(1 * time.Second)

	}
}

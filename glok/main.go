package main
import "fmt"
import "time"
import "github.com/inancgumus/screen"

func main() {
	zero := [5]string{
		"█ █ █",
		"█   █",
		"█   █",
		"█   █",
		"█ █ █",
	}
	one := [5]string{
		"███  ",
		"  █  ",
		"  █  ",
		"  █  ",
		"█ █ █",
	}

	two := [5]string{
		"█ █ █",
		"    █",
		"█ █ █",
		"█    ",
		"█ █ █",
	}
	three := [5]string{
		"█ █ █",
		"    █",
		"█ █ █",
		"    █",
		"█ █ █",
	}
	four := [5]string{
		"█   █",
		"█   █",
		"█ █ █",
		"    █",
		"    █",
	}
	five := [5]string{
		"█ █ █",
		"█    ",
		"█ █ █",
		"    █",
		"█ █ █",
	}
	six := [5]string{
		"█ █ █",
		"█    ",
		"█ █ █",
		"█   █",
		"█ █ █",
	}
	seven := [5]string{
		"█ █ █",
		"    █",
		"    █",
		"    █",
		"    █",
	}
	eight := [5]string{
		"█ █ █",
		"█   █",
		"█ █ █",
		"█   █",
		"█ █ █",
	}
	nine := [5]string{
		"█ █ █",
		"█   █",
		"█ █ █",
		"    █",
		"█ █ █",
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
		screen.Clear()
		screen.MoveTopLeft()
	}
}
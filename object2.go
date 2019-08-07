package main

import "fmt"

type Color byte
type Box struct {
	width, height, depth float64
	color Color
}

const(
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)


type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) setColor(c Color)  {
	b.color = c
}

func (bl BoxList) BiggestBoxColor() Color {
	value := 0.00
	color := Color(WHITE)
	for _, b := range bl {
		if boxVolume := b.Volume(); boxVolume > value {
			value = boxVolume
			color = b.color
		}
	}
	return color
}

func(bl BoxList) PrintAllBlack()  {
	for i := range bl {
		bl[i].setColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main()  {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(),"cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestBoxColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PrintAllBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestBoxColor().String())
}






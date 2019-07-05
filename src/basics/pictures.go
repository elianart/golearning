package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	b := make([][]uint8, dy) // len(b)=0, cap(b)=5
	//b:= [dx][dy]

	for i := 0; i < dy; i++ {
		b[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			//b[i][j] = uint8((i+j)/2)
			//b[i][j] =  uint8(i*j)
			b[i][j] = uint8(j ^ i)
		}
	}

	return b
}

func main() {
	pic.Show(Pic)
}

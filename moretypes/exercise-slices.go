package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		res[y] = make([]uint8, dx)
	}
	return res
}

func main() {
	pic.Show(Pic)
}

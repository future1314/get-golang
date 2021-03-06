package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		image[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			image[i][j] = (uint8)(i*j) % 255 // 这里的函数用要求给的三个函数可以得到图像的不同的效果
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}

package main

import (
	"github.com/disintegration/imaging"
)

func main() {
	src, _ := imaging.Open("asset/nike.jpg")
	dstImg := imaging.Resize(
		src,
		200,
		0,
		imaging.Lanczos,
	)

	imaging.Save(dstImg, "dist/resized_nike.jpg")
}

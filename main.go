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

	src, _ = imaging.Open("asset/white_sign.png")
	dstImg = imaging.Resize(
		src,
		200,
		0,
		imaging.Lanczos,
	)

	imaging.Save(dstImg, "dist/resized_white_sign.png")
}

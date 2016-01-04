/**
 * https://tour.golang.org/methods/16
 *
 * Exercise: Images
 *
 * Define your own Image type, implement the necessary methods,
 * and call pic.ShowImage. 
 */
package main

import (
    "image"
    "image/color"
    "golang.org/x/tour/pic"
)

type Image struct{}

const width = 75
const height = 75

func (img *Image) ColorModel() color.Model {

    return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {

    return image.Rect(0, 0, width, height)
}

func (img *Image) At(x, y int) color.Color {

    return color.RGBA{

        uint8(x * y),
        uint8((x + y) / 2),
        uint8((x + y) * 2),
        255,
    }
}

func main() {

    pic.ShowImage(&Image{})
}

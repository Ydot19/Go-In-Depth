package main

/**
Lissajous generates GIF animations of random Lissajous figures.
Section highlights animation capabilities of Golang

Exercise 1.5: Change the Lissajous program's color palette to Green on black
 */

import (
	"image/color"
)

var exercise15 = []color.Color{color.RGBA{R: 1, G: 152, B: 117, A: 1}, color.Black}

//const (
//	whiteIndex = 0 // first color in palette
//	blackIndex = 1 // next color in palette
//)

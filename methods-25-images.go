/*
https://tour.golang.org/methods/25
*/
package main

import "golang.org/x/tour/pic"

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
package main

import (
	"math"
)

func getAreaQuadrilateral(base, height float64) float64 {
	return base * height
}

func getAreaTriangle(base, height float64) float64 {
	return (base * height) / 2
}

func getAreaCircle(radius float64) float64 {
	return math.Pi * radius * radius
}

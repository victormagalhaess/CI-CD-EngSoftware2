package main

import (
	"testing"
)

func TestGetAreaQuadrilateral(t *testing.T) {
	base := 10.0
	height := 20.0
	expected := 200.0
	actual := getAreaQuadrilateral(base, height)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}

func TestGetAreaTriangle(t *testing.T) {
	base := 10.0
	height := 20.0
	expected := 100.0
	actual := getAreaTriangle(base, height)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}

func TestGetAreaCircle(t *testing.T) {
	radius := 10.0
	expected := 314.1592653589793
	actual := getAreaCircle(radius)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}

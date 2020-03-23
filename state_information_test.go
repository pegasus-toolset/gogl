package gogl

import (
	"testing"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.0/glfw"
)

func InitGL() {
	if !glfw.Init() {
		panic("Failed to initialize GLFW")
	}

	window, err := glfw.CreateWindow(800, 600, "gogl", nil, nil)
	if err != nil {
		panic("Failed to create window")
	}
	window.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		panic("Failed to initialize OpenGL")
	}
}

func TestGetAlphaBits(t *testing.T) {
	InitGL()

	expected := int32(8)
	result := GetAlphaBits()
	if result != expected {
		t.Errorf("Expected GetAlphaBits() to return %d, got %d instead", expected, result)
	}
}

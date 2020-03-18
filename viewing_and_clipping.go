package gogl

import "github.com/go-gl/gl/v2.1/gl"

// Scissor sets a scissor box, which limits the drawing to a specified
// rectangle.
func Scissor(x, y, width, height int32) {
	gl.Scissor(x, y, width, height)
}

// Viewport sets the viewport, which specifies the affine transformation of x
// and y from normalized device coordinates to window coordinates.
func Viewport(x, y, width, height int32) {
	gl.Viewport(x, y, width, height)
}

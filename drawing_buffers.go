package gogl

import "github.com/go-gl/gl/v2.1/gl"

// Clear clears buffers to preset values.
//
// The preset values can be set by ClearColor, ClearDepth or ClearStencil.
//
// The scissor box, dithering, and buffer writemasks can affect the Clear
// function.
func Clear(mask GLEnum) {
	gl.Clear(uint32(mask))
}

// DrawArrays renders primitives from array data.
func DrawArrays(mode GLEnum, first, count int32) {
	gl.DrawArrays(uint32(mode), first, count)
}

// TODO: DrawElements

// Finish blocks execution until all previously called commands are finished.
func Finish() {
	gl.Finish()
}

// Flush empties different buffer commands, causing all commands to be executed
// as quickly as possible.
func Flush() {
	gl.Flush()
}

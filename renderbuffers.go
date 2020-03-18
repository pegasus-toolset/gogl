package gogl

import "github.com/go-gl/gl/v2.1/gl"

// BindRenderbuffer binds a given Renderbuffer to a target, which must be
// GLRenderbuffer.
func BindRenderbuffer(target GLEnum, renderbuffer Renderbuffer) {
	gl.BindRenderbuffer(uint32(target), uint32(renderbuffer))
}

// CreateRenderbuffer creates and initializes a Renderbuffer object.
func CreateRenderbuffer() Renderbuffer {
	var renderbuffer uint32
	gl.GenRenderbuffers(1, &renderbuffer)
	return Renderbuffer(renderbuffer)
}

// Delete deletes the Renderbuffer object. This function has no effect if the
// render buffer has already been deleted.
func (renderbuffer Renderbuffer) Delete() {
	// TODO: Is it somehow possible to get &uint32(renderbuffer) without assigning it to renderbuffers?
	renderbuffers := uint32(renderbuffer)
	gl.DeleteRenderbuffers(1, &renderbuffers)
}

// TODO: GetRenderbufferParameter

// IsRenderbuffer returns true if the Renderbuffer is valid and false otherwise.
func (renderbuffer Renderbuffer) IsRenderbuffer() bool {
	return gl.IsRenderbuffer(uint32(renderbuffer))
}

// RenderbufferStorage creates and initializes a renderbuffer object's data
// store.
func RenderbufferStorage(target, internalFormat GLEnum, width, height int32) {
	gl.RenderbufferStorage(uint32(target), uint32(internalFormat), width, height)
}

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

// GetRenderbufferWidth returns an int32 indicating the width of the image of
// the currently bound renderbuffer.
func GetRenderbufferWidth(target GLEnum) int32 {
	var params int32
	gl.GetRenderbufferParameteriv(uint32(target), gl.RENDERBUFFER_WIDTH, &params)
	return params
}

// GetRenderbufferHeight returns an int32 indicating the height of the image of
// the currently bound renderbuffer.
func GetRenderbufferHeight(target GLEnum) int32 {
	var params int32
	gl.GetRenderbufferParameteriv(uint32(target), gl.RENDERBUFFER_HEIGHT, &params)
	return params
}

// GetRenderbufferInternalFormat returns a GLEnum indicating the internal format
// of the currently bound renderbuffer. The default is GLRGBA4.
func GetRenderbufferInternalFormat(target GLEnum) GLEnum {
	var params int32
	gl.GetRenderbufferParameteriv(uint32(target), gl.RENDERBUFFER_INTERNAL_FORMAT, &params)
	return GLEnum(params)
}

// GetRenderbufferGreenSize returns an int32 that is the resolution size (in
// bits) for the green color.
func GetRenderbufferGreenSize(target GLEnum) int32 {
	var params int32
	gl.GetRenderbufferParameteriv(uint32(target), gl.RENDERBUFFER_GREEN_SIZE, &params)
	return params
}

// GetRenderbufferBlueSize returns an int32 that is the resolution size (in
// bits) for the blue color.
func GetRenderbufferBlueSize(target GLEnum) int32 {
	var params int32
	gl.GetRenderbufferParameteriv(uint32(target), gl.RENDERBUFFER_BLUE_SIZE, &params)
	return params
}

// GetRenderbufferRedSize returns an int32 that is the resolution size (in bits)
// for the red color.
func GetRenderbufferRedSize(target GLEnum) int32 {
	var params int32
	gl.GetRenderbufferParameteriv(uint32(target), gl.RENDERBUFFER_RED_SIZE, &params)
	return params
}

// GetRenderbufferAlphaSize returns an int32 that is the resolution size (in
// bits) for the alpha component.
func GetRenderbufferAlphaSize(target GLEnum) int32 {
	var params int32
	gl.GetRenderbufferParameteriv(uint32(target), gl.RENDERBUFFER_ALPHA_SIZE, &params)
	return params
}

// GetRenderbufferDepthSize returns an int32 that is the resolution size (in
// bits) for the depth component.
func GetRenderbufferDepthSize(target GLEnum) int32 {
	var params int32
	gl.GetRenderbufferParameteriv(uint32(target), gl.RENDERBUFFER_DEPTH_SIZE, &params)
	return params
}

// GetRenderbufferStencilSize returns an int32 that is the resolution size (in
// bits) for the stencil component.
func GetRenderbufferStencilSize(target GLEnum) int32 {
	var params int32
	gl.GetRenderbufferParameteriv(uint32(target), gl.RENDERBUFFER_STENCIL_SIZE, &params)
	return params
}

// IsRenderbuffer returns true if the Renderbuffer is valid and false otherwise.
func (renderbuffer Renderbuffer) IsRenderbuffer() bool {
	return gl.IsRenderbuffer(uint32(renderbuffer))
}

// RenderbufferStorage creates and initializes a renderbuffer object's data
// store.
func RenderbufferStorage(target, internalFormat GLEnum, width, height int32) {
	gl.RenderbufferStorage(uint32(target), uint32(internalFormat), width, height)
}

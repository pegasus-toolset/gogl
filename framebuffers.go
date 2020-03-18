package gogl

import "github.com/go-gl/gl/v2.1/gl"

// BindFramebuffer binds a given Framebuffer to a target.
func BindFramebuffer(target GLEnum, framebuffer Framebuffer) {
	gl.BindFramebuffer(uint32(target), uint32(framebuffer))
}

// CheckFramebufferStatus returns the completeness status of the Framebuffer
// object.
func CheckFramebufferStatus(target GLEnum) GLEnum {
	return GLEnum(gl.CheckFramebufferStatus(uint32(target)))
}

// CreateFramebuffer creates and initializes a Framebuffer object.
func CreateFramebuffer() Framebuffer {
	var framebuffer uint32
	gl.GenFramebuffers(1, &framebuffer)
	return Framebuffer(framebuffer)
}

// Delete deletes the Framebuffer object. This function has no effect if the
// frame buffer has already been deleted.
func (framebuffer Framebuffer) Delete() {
	// TODO: Is it somehow possible to get &uint32(framebuffer) without assigning it to framebuffers?
	framebuffers := uint32(framebuffer)
	gl.DeleteFramebuffers(1, &framebuffers)
}

// FramebufferRenderbuffer attaches a Renderbuffer object to a Framebuffer
// object.
func FramebufferRenderbuffer(target, attachment, renderbuffertarget GLEnum, renderbuffer Renderbuffer) {
	gl.FramebufferRenderbuffer(uint32(target), uint32(attachment), uint32(renderbuffertarget), uint32(renderbuffer))
}

// FramebufferTexture2D attaches a texture to a Framebuffer.
func FramebufferTexture2D(target, attachment, textarget GLEnum, texture Texture, level int32) {
	gl.FramebufferTexture2D(uint32(target), uint32(attachment), uint32(textarget), uint32(texture), level)
}

// TODO: GetFramebufferAttachmentParameter

// IsFramebuffer returns true if the Framebuffer is valid and false otherwise.
func (framebuffer Framebuffer) IsFramebuffer() bool {
	return gl.IsFramebuffer(uint32(framebuffer))
}

// TODO: ReadPixels

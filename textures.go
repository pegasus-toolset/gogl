package gogl

import (
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

// BindTexture binds a given Texture to a target (binding point).
func BindTexture(target GLEnum, texture Texture) {
	gl.BindTexture(uint32(target), uint32(texture))
}

// CompressedTexImage2D and CompressedTexImage3D specify a two- or
// three-dimensional texture image in a compressed format.
//
// Compressed image formats must be enabled by OpenGL extensions before using
// these functions.
func CompressedTexImage2D(target GLEnum, level int32, internalformat GLEnum, width, height, border, imageSize int32, pixels []float32) {
	gl.CompressedTexImage2D(uint32(target), level, uint32(internalformat), width, height, border, imageSize, unsafe.Pointer(&pixels[0]))
}

// CompressedTexSubImage2D specifies a two-dimensional sub-rectangle for a
// texture image in a compressed format.
//
// Compressed image formats must be enabled by OpenGL extensions before using
// this function.
func CompressedTexSubImage2D(target GLEnum, level, xoffset, yoffset, width, height int32, format GLEnum, imageSize int32, pixels []float32) {
	gl.CompressedTexSubImage2D(uint32(target), level, xoffset, yoffset, width, height, uint32(format), imageSize, unsafe.Pointer(&pixels[0]))
}

// CopyTexImage2D copies pixels from the current Framebuffer into a 2D texture
// image.
func CopyTexImage2D(target GLEnum, level int32, internalformat GLEnum, x, y, width, height, border int32) {
	gl.CopyTexImage2D(uint32(target), level, uint32(internalformat), x, y, width, height, border)
}

// CopyTexSubImage2D copies pixels from the current Framebuffer into an existing
// 2D texture sub-image.
func CopyTexSubImage2D(target GLEnum, level, xoffset, yoffset, x, y, width, height int32) {
	gl.CopyTexSubImage2D(uint32(target), level, xoffset, yoffset, x, y, width, height)
}

// CreateTexture creates and initializes a Texture object.
func CreateTexture() Texture {
	var texture uint32
	gl.GenTextures(1, &texture)
	return Texture(texture)
}

// Delete deletes the Texture object. This method has no effect if the texture
// has already been deleted.
func (texture Texture) Delete() {
	// TODO: Is it somehow possible to get &uint32(texture) without assigning it to textures?
	textures := uint32(texture)
	gl.DeleteTextures(1, &textures)
}

// GenerateMipmap generates a set of mipmaps for a Texture object.
func GenerateMipmap(target GLEnum) {
	gl.GenerateMipmap(uint32(target))
}

// TODO: GetTexParameter

// IsTexture returns true if the Texture is valid and false otherwise.
func (texture Texture) IsTexture() bool {
	return gl.IsTexture(uint32(texture))
}

// TexImage2D specifies a two-dimensional texture image.
func TexImage2D(target GLEnum, level int32, internalformat GLEnum, width, height, border int32, format, xtype GLEnum, pixels []float32) {
	gl.TexImage2D(uint32(target), level, int32(internalformat), width, height, border, uint32(format), uint32(xtype), unsafe.Pointer(&pixels[0]))
}

// TexSubImage2D specifies a sub-rectangle of the current texture.
func TexSubImage2D(target GLEnum, level, xoffset, yoffset, width, height int32, format, xtype GLEnum, pixels []float32) {
	gl.TexSubImage2D(uint32(target), level, xoffset, yoffset, width, height, uint32(format), uint32(xtype), unsafe.Pointer(&pixels[0]))
}

// TexParameterf and TexParameteri set texture parameters.
func TexParameterf(target, pname GLEnum, param float32) {
	gl.TexParameterf(uint32(target), uint32(pname), param)
}

// TexParameteri and TexParameterf set texture parameters.
func TexParameteri(target, pname GLEnum, param int32) {
	gl.TexParameteri(uint32(target), uint32(pname), param)
}

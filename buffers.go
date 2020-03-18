package gogl

import (
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

// BindBuffer binds a given Buffer to a target.
func BindBuffer(target GLEnum, buffer Buffer) {
	gl.BindBuffer(uint32(target), uint32(buffer))
}

// BufferData initializes and creates the buffer object's data store.
func BufferData(target GLEnum, srcData []float32, usage GLEnum) {
	gl.BufferData(uint32(target), len(srcData)*4, unsafe.Pointer(&srcData[0]), uint32(usage))
}

// BufferSubData updates a subset of a buffer object's data store.
func BufferSubData(target GLEnum, offset int, srcData []float32) {
	gl.BufferSubData(uint32(target), offset*4, len(srcData)*4, unsafe.Pointer(&srcData[0]))
}

// CreateBuffer creates and initializes a Buffer storing data such as vertices
// or colors.
func CreateBuffer() Buffer {
	var buffer uint32
	gl.GenBuffers(1, &buffer)
	return Buffer(buffer)
}

// Delete deletes the Buffer. This method has no effect if the buffer has
// already been deleted.
func (buffer Buffer) Delete() {
	// TODO: Is it somehow possible to get &uint32(buffer) without assigning it to buffers?
	buffers := uint32(buffer)
	gl.DeleteBuffers(1, &buffers)
}

// TODO: GetBufferParameter

// IsBuffer returns true if the Buffer is valid and false otherwise.
func (buffer Buffer) IsBuffer() bool {
	return gl.IsBuffer(uint32(buffer))
}

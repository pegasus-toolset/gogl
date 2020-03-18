package gogl

import "github.com/go-gl/gl/v2.1/gl"

// DisableVertexAttribArray turns the generic vertex attribute array off at a
// given index position.
func DisableVertexAttribArray(index uint32) {
	gl.DisableVertexAttribArray(index)
}

// EnableVertexAttribArray turns on the generic attribute array at the specified
// index into the list of attribute arrays.
//
// In OpenGL, values that apply to a specific vertex are stored in attributes.
// These are only available to the Go code and the vertex shader. Attributes are
// referenced by an index number into the list of attributes maintained by the
// GPU. Some vertex attribute indices may have predefined purposes, depending on
// the platform and/or the GPU. Others are assigned by the OpenGL layer when you
// create the attributes.
//
// Either way, since attributes cannot be used unless enabled, and are disabled
// by default, you need to call EnableVertexAttribArray to enable individual
// attributes so that they can be used. Once that's been done, other functions
// can be used to access the attribute, including VertexAttribPointer,
// VertexAttrib, and GetVertexAttrib.
func EnableVertexAttribArray(index uint32) {
	gl.EnableVertexAttribArray(index)
}

// TODO: GetActiveAttrib

// TODO: GetActiveUniform

// GetAttribLocation returns the location of an attribute variable in the
// Program.
func (program Program) GetAttribLocation(name string) int32 {
	return gl.GetAttribLocation(uint32(program), gl.Str(name+"\x00"))
}

// TODO: GetUniform

// GetUniformLocation returns the location of a specific uniform variable which
// is part of the Program. The uniform variable is returned as a UniformLocation
// object, which is an opaque identifier used to specify where in the GPU's
// memory that uniform variable is located.
//
// Once you have the uniform's location, you can access the uniform itself using
// one of the other uniform access functions, passing in the uniform location as
// one of the inputs.
//
// The uniform itself is declared in the shader program using GLSL.
func (program Program) GetUniformLocation(name string) UniformLocation {
	return UniformLocation(gl.GetUniformLocation(uint32(program), gl.Str(name+"\x00")))
}

// TODO: GetVertexAttrib

// TODO: GetVertexAttribOffset

// Uniform1Float specifies values of uniform variables.
func Uniform1Float(location UniformLocation, v0 float32) {
	gl.Uniform1f(int32(location), v0)
}

// Uniform1FloatArray specifies values of uniform variables.
func Uniform1FloatArray(location UniformLocation, value []float32) {
	gl.Uniform1fv(int32(location), 1, &value[0])
}

// Uniform1Int specifies values of uniform variables.
func Uniform1Int(location UniformLocation, v0 int32) {
	gl.Uniform1i(int32(location), v0)
}

// Uniform1IntArray specifies values of uniform variables.
func Uniform1IntArray(location UniformLocation, value []int32) {
	gl.Uniform1iv(int32(location), 1, &value[0])
}

// Uniform2Float specifies values of uniform variables.
func Uniform2Float(location UniformLocation, v0, v1 float32) {
	gl.Uniform2f(int32(location), v0, v1)
}

// Uniform2FloatArray specifies values of uniform variables.
func Uniform2FloatArray(location UniformLocation, value []float32) {
	gl.Uniform2fv(int32(location), 1, &value[0])
}

// Uniform2Int specifies values of uniform variables.
func Uniform2Int(location UniformLocation, v0, v1 int32) {
	gl.Uniform2i(int32(location), v0, v1)
}

// Uniform2IntArray specifies values of uniform variables.
func Uniform2IntArray(location UniformLocation, value []int32) {
	gl.Uniform2iv(int32(location), 1, &value[0])
}

// Uniform3Float specifies values of uniform variables.
func Uniform3Float(location UniformLocation, v0, v1, v2 float32) {
	gl.Uniform3f(int32(location), v0, v1, v2)
}

// Uniform3FloatArray specifies values of uniform variables.
func Uniform3FloatArray(location UniformLocation, value []float32) {
	gl.Uniform3fv(int32(location), 1, &value[0])
}

// Uniform3Int specifies values of uniform variables.
func Uniform3Int(location UniformLocation, v0, v1, v2 int32) {
	gl.Uniform3i(int32(location), v0, v1, v2)
}

// Uniform3IntArray specifies values of uniform variables.
func Uniform3IntArray(location UniformLocation, value []int32) {
	gl.Uniform3iv(int32(location), 1, &value[0])
}

// Uniform4Float specifies values of uniform variables.
func Uniform4Float(location UniformLocation, v0, v1, v2, v3 float32) {
	gl.Uniform4f(int32(location), v0, v1, v2, v3)
}

// Uniform4FloatArray specifies values of uniform variables.
func Uniform4FloatArray(location UniformLocation, value []float32) {
	gl.Uniform4fv(int32(location), 1, &value[0])
}

// Uniform4Int specifies values of uniform variables.
func Uniform4Int(location UniformLocation, v0, v1, v2, v3 int32) {
	gl.Uniform4i(int32(location), v0, v1, v2, v3)
}

// Uniform4IntArray specifies values of uniform variables.
func Uniform4IntArray(location UniformLocation, value []int32) {
	gl.Uniform4iv(int32(location), 1, &value[0])
}

// UniformMatrix2fv specifies matrix values for uniform variables.
//
// The three versions of this method (UniformMatrix2fv, UniformMatrix3fv, and
// UniformMatrix4fv) take as the input value 2-component, 3-component, and
// 4-component square matrices, respectively. They are expected to have 4, 9 or
// 16 floats.
func UniformMatrix2fv(location UniformLocation, transpose bool, value []float32) {
	gl.UniformMatrix2fv(int32(location), 1, transpose, &value[0])
}

// UniformMatrix3fv specifies matrix values for uniform variables.
//
// The three versions of this method (UniformMatrix2fv, UniformMatrix3fv, and
// UniformMatrix4fv) take as the input value 2-component, 3-component, and
// 4-component square matrices, respectively. They are expected to have 4, 9 or
// 16 floats.
func UniformMatrix3fv(location UniformLocation, transpose bool, value []float32) {
	gl.UniformMatrix3fv(int32(location), 1, transpose, &value[0])
}

// UniformMatrix4fv specifies matrix values for uniform variables.
//
// The three versions of this method (UniformMatrix2fv, UniformMatrix3fv, and
// UniformMatrix4fv) take as the input value 2-component, 3-component, and
// 4-component square matrices, respectively. They are expected to have 4, 9 or
// 16 floats.
func UniformMatrix4fv(location UniformLocation, transpose bool, value []float32) {
	gl.UniformMatrix4fv(int32(location), 1, transpose, &value[0])
}

// VertexAttrib1f specifies constant values for generic vertex attributes.
func VertexAttrib1f(index uint32, v0 float32) {
	gl.VertexAttrib1f(index, v0)
}

// VertexAttrib2f specifies constant values for generic vertex attributes.
func VertexAttrib2f(index uint32, v0, v1 float32) {
	gl.VertexAttrib2f(index, v0, v1)
}

// VertexAttrib3f specifies constant values for generic vertex attributes.
func VertexAttrib3f(index uint32, v0, v1, v2 float32) {
	gl.VertexAttrib3f(index, v0, v1, v2)
}

// VertexAttrib4f specifies constant values for generic vertex attributes.
func VertexAttrib4f(index uint32, v0, v1, v2, v3 float32) {
	gl.VertexAttrib4f(index, v0, v1, v2, v3)
}

// VertexAttrib1fv specifies constant values for generic vertex attributes.
func VertexAttrib1fv(index uint32, value []float32) {
	gl.VertexAttrib1fv(index, &value[0])
}

// VertexAttrib2fv specifies constant values for generic vertex attributes.
func VertexAttrib2fv(index uint32, value []float32) {
	gl.VertexAttrib2fv(index, &value[0])
}

// VertexAttrib3fv specifies constant values for generic vertex attributes.
func VertexAttrib3fv(index uint32, value []float32) {
	gl.VertexAttrib3fv(index, &value[0])
}

// VertexAttrib4fv specifies constant values for generic vertex attributes.
func VertexAttrib4fv(index uint32, value []float32) {
	gl.VertexAttrib4fv(index, &value[0])
}

// TODO: VertexAttribPointer

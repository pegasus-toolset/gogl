package gogl

import (
	"strings"

	"github.com/go-gl/gl/v2.1/gl"
)

// AttachShader attaches either a fragment or vertex Shader to the Program.
func (program Program) AttachShader(shader Shader) {
	gl.AttachShader(uint32(program), uint32(shader))
}

// BindAttribLocation binds a generic vertex index to an attribute variable.
func BindAttribLocation(program Program, index uint32, name string) {
	gl.BindAttribLocation(uint32(program), index, gl.Str(name+"\x00"))
}

// Compile compiles the GLSL shader into binary data so that it can be used by a
// Program.
func (shader Shader) Compile() {
	gl.CompileShader(uint32(shader))
}

// CreateProgram creates and initializes a Program object.
func CreateProgram() Program {
	return Program(gl.CreateProgram())
}

// CreateShader creates a Shader that can then be configured further using
// ShaderSource and CompileShader.
func CreateShader(xtype GLEnum) Shader {
	return Shader(gl.CreateShader(uint32(xtype)))
}

// Delete deletes the Program object. This method has no effect if the program
// has already been deleted.
func (program Program) Delete() {
	gl.DeleteProgram(uint32(program))
}

// Delete marks the Shader object for deletion. It will then be deleted whenever
// the shader is no longer in use. This function has no effect if the shader has
// already been deleted, and the Shader is automatically marked for deletion
// when it is destroyed by the garbage collector.
func (shader Shader) Delete() {
	gl.DeleteShader(uint32(shader))
}

// DetachShader detaches a previously attached Shader from the Program.
func (program Program) DetachShader(shader Shader) {
	gl.DetachShader(uint32(program), uint32(shader))
}

// TODO: GetAttachedShaders

// GetDeleteStatus returns a bool indicating whether or not the program is
// flagged for deletion.
func (program Program) GetDeleteStatus() bool {
	var params int32
	gl.GetProgramiv(uint32(program), gl.DELETE_STATUS, &params)
	return params == gl.TRUE
}

// GetLinkStatus returns a bool indicating whether or not the last link
// operation was successful.
func (program Program) GetLinkStatus() bool {
	var params int32
	gl.GetProgramiv(uint32(program), gl.LINK_STATUS, &params)
	return params == gl.TRUE
}

// GetValidateStatus returns a bool indicating whether or not the last
// validation operation was successful.
func (program Program) GetValidateStatus() bool {
	var params int32
	gl.GetProgramiv(uint32(program), gl.VALIDATE_STATUS, &params)
	return params == gl.TRUE
}

// GetAttachedShaders returns an int32 indicating the number of attached shaders
// to a program.
func (program Program) GetAttachedShaders() int32 {
	var params int32
	gl.GetProgramiv(uint32(program), gl.ATTACHED_SHADERS, &params)
	return params
}

// GetActiveAttributes returns an int32 indicating the number of active
// attribute variables to a program.
func (program Program) GetActiveAttributes() int32 {
	var params int32
	gl.GetProgramiv(uint32(program), gl.ACTIVE_ATTRIBUTES, &params)
	return params
}

// GetActiveUniforms returns an int32 indicating the number of active uniform
// variables to a program.
func (program Program) GetActiveUniforms() int32 {
	var params int32
	gl.GetProgramiv(uint32(program), gl.ACTIVE_UNIFORMS, &params)
	return params
}

// GetInfoLog returns the information log for the Program object. It contains
// errors that occurred during failed linking or validation of Program objects.
func (program Program) GetInfoLog() string {
	var bufSize int32
	gl.GetProgramiv(uint32(program), gl.INFO_LOG_LENGTH, &bufSize)
	infoLog := strings.Repeat("\x00", int(bufSize+1))
	gl.GetProgramInfoLog(uint32(program), bufSize, nil, gl.Str(infoLog))
	return infoLog
}

// GetDeleteStatus returns a bool indicating whether or not the shader is
// flagged for deletion.
func (shader Shader) GetDeleteStatus() bool {
	var params int32
	gl.GetShaderiv(uint32(shader), gl.DELETE_STATUS, &params)
	return params == gl.TRUE
}

// GetCompileStatus returns a bool indicating whether or not the last shader
// compilation was successful.
func (shader Shader) GetCompileStatus() bool {
	var params int32
	gl.GetShaderiv(uint32(shader), gl.COMPILE_STATUS, &params)
	return params == gl.TRUE
}

// GetShaderType returns a GLEnum indicating whether the shader is a vertex
// shader (GLVertexShader) or fragment shader (GLFragmentShader) object.
func (shader Shader) GetShaderType() GLEnum {
	var params int32
	gl.GetShaderiv(uint32(shader), gl.SHADER_TYPE, &params)
	return GLEnum(params)
}

// TODO: GetShaderPrecisionFormat

// GetInfoLog returns the information log for the Shader object. It contains
// warnings, debugging and compile information.
func (shader Shader) GetInfoLog() string {
	var bufSize int32
	gl.GetShaderiv(uint32(shader), gl.INFO_LOG_LENGTH, &bufSize)
	infoLog := strings.Repeat("\x00", int(bufSize+1))
	gl.GetShaderInfoLog(uint32(shader), bufSize, nil, gl.Str(infoLog))
	return infoLog
}

// TODO: GetShaderSource

// IsProgram returns true if the Program is valid, false otherwise.
func (program Program) IsProgram() bool {
	return gl.IsProgram(uint32(program))
}

// IsShader returns true if the Shader is valid, false otherwise.
func (shader Shader) IsShader() bool {
	return gl.IsShader(uint32(shader))
}

// Link links the Program, completing the process of preparing the GPU code for
// the program's fragment and vertex shaders.
func (program Program) Link() {
	gl.LinkProgram(uint32(program))
}

// Source sets the source code of the Shader.
func (shader Shader) Source(source string) {
	cstrs, free := gl.Strs(source + "\x00")
	gl.ShaderSource(uint32(shader), 1, cstrs, nil)
	free()
}

// Use sets the Program as part of the current rendering state.
func (program Program) Use() {
	gl.UseProgram(uint32(program))
}

// Validate validates the Program. It checks if it is successfully linked and if
// it can be used in the current OpenGL state.
func (program Program) Validate() {
	gl.ValidateProgram(uint32(program))
}

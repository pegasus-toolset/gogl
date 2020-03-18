// Package gogl is a Go wrapper for OpenGL originally based on WebGL. Its
// purpose is to simplify and adapt OpenGL functions to the features of the Go
// programming language.
package gogl

import "github.com/go-gl/gl/v2.1/gl"

// Init intializes the OpenGL bindings by loading the function pointers (for
// each OpenGL function) from the active OpenGL context.
//
// It must be called under the presence of an active OpenGL context, e.g.,
// always after calling window.MakeContextCurrent and always before calling any
// OpenGL functions exported by this package.
//
// On Windows, Init loads pointers that are context-specific (and hence you must
// re-init if switching between OpenGL contexts, although not calling Init again
// after switching between OpenGL contexts may work if the contexts belong to
// the same graphics driver/device).
//
// On macOS and the other POSIX systems, the behavoir is different, but code
// written compatible with the Windows behavior is compatible with macOS and the
// other POSIX systems. That is, always Init under an active OpenGL context, and
// always re-init after switching graphics contexts.
//
// For information about caveats of Init, you should read the "Platform Specific
// Function Retrieval" section of
// https://www.opengl.org/wiki/Load_OpenGL_Functions.
func Init() error {
	return gl.Init()
}

// GetString returns a string describing the current GL connection.
func GetString(name GLEnum) string {
	return gl.GoStr(gl.GetString(uint32(name)))
}

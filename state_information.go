package gogl

import "github.com/go-gl/gl/v2.1/gl"

// ActiveTexture specifies which texture unit to make active.
func ActiveTexture(texture GLEnum) {
	gl.ActiveTexture(uint32(texture))
}

// BlendColor is used to set the source and destination blending factors.
func BlendColor(red, green, blue, alpha float32) {
	gl.BlendColor(red, green, blue, alpha)
}

// BlendEquation is used to set both the RGB blend equation and alpha blend
// equation to a single equation.
//
// The blend equation determines how a new pixel is combined with a pixel
// already in the Framebuffer.
func BlendEquation(mode GLEnum) {
	gl.BlendEquation(uint32(mode))
}

// BlendEquationSeparate is used to set the RGB blend equation and alpha blend
// equation separately.
//
// The blend equation determines how a new pixel is combined with a pixel
// already in the Framebuffer.
func BlendEquationSeparate(modeRGB, modeAlpha GLEnum) {
	gl.BlendEquationSeparate(uint32(modeRGB), uint32(modeAlpha))
}

// BlendFunc defines which function is used for blending pixel arithmetic.
func BlendFunc(sfactor, dfactor GLEnum) {
	gl.BlendFunc(uint32(sfactor), uint32(dfactor))
}

// BlendFuncSeparate defines which function is used for blending pixel
// arithmetic for RGB and alpha components separately.
func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha GLEnum) {
	gl.BlendFuncSeparate(uint32(srcRGB), uint32(dstRGB), uint32(srcAlpha), uint32(dstAlpha))
}

// ClearColor specifies the color values used when clearing color buffers.
//
// This specifies what color values to use when calling the Clear method. The
// values are clamped between 0 and 1.
func ClearColor(red, green, blue, alpha float32) {
	gl.ClearColor(red, green, blue, alpha)
}

// ClearDepth specifies the clear value for the depth buffer.
//
// This specifies what depth value to use when calling the Clear method. The
// value is clamped between 0 and 1.
func ClearDepth(depth float32) {
	gl.ClearDepth(float64(depth))
}

// ClearStencil specifies the clear value for the stencil buffer.
//
// This specifies what stencil value to use when calling the Clear method.
func ClearStencil(s int32) {
	gl.ClearStencil(s)
}

// ColorMask sets which color components to enable or to disable when drawing or
// rendering to a Framebuffer.
func ColorMask(red, green, blue, alpha bool) {
	gl.ColorMask(red, green, blue, alpha)
}

// CullFace specifies whether or not front- and/or back-facing polygons can be
// culled.
func CullFace(mode GLEnum) {
	gl.CullFace(uint32(mode))
}

// DepthFunc specifies a function that compares incoming pixel depth to the
// current depth buffer value.
func DepthFunc(xfunc GLEnum) {
	gl.DepthFunc(uint32(xfunc))
}

// DepthMask sets whether writing into the depth buffer is enabled or disabled.
func DepthMask(flag bool) {
	gl.DepthMask(flag)
}

// DepthRange specifies the depth range mapping from normalized device
// coordinates to window or viewport coordinates.
func DepthRange(zNear, zFar float32) {
	gl.DepthRange(float64(zNear), float64(zFar))
}

// Disable disables specific OpenGL capabilities.
func Disable(cap GLEnum) {
	gl.Disable(uint32(cap))
}

// Enable enables specific OpenGL capabilities.
func Enable(cap GLEnum) {
	gl.Enable(uint32(cap))
}

// FrontFace specifies whether polygons are front- or back-facing by setting a
// winding orientation.
func FrontFace(mode GLEnum) {
	gl.FrontFace(uint32(mode))
}

// TODO: GetParameter

// GetError returns error information.
func GetError() GLEnum {
	return GLEnum(gl.GetError())
}

// Hint specifies hints for certain behaviors. The interpretation of these hints
// depend on the implementation.
func Hint(target, mode GLEnum) {
	gl.Hint(uint32(target), uint32(mode))
}

// IsEnabled tests whether a specific OpenGL capability is enabled or not.
//
// By default, all capabilities except GLDither are disabled.
func IsEnabled(cap GLEnum) bool {
	return gl.IsEnabled(uint32(cap))
}

// LineWidth sets the line width of rasterized lines.
func LineWidth(width float32) {
	gl.LineWidth(width)
}

// PixelStorei specifies the pixel storage modes.
func PixelStorei(pname GLEnum, param int32) {
	gl.PixelStorei(uint32(pname), param)
}

// PolygonOffset specifies the scale factors and units to calculate depth
// values.
//
// The offset is added before the depth test is performed and before the value
// is written into the depth buffer.
func PolygonOffset(factor, units float32) {
	gl.PolygonOffset(factor, units)
}

// SampleCoverage specifies multi-sample coverage parameters for anti-aliasing
// effects.
func SampleCoverage(value float32, invert bool) {
	gl.SampleCoverage(value, invert)
}

// StencilFunc sets the front and back function and reference value for stencil
// testing.
//
// Stenciling enables and disables drawing on a per-pixel basis. It is typically
// used in multipass rendering to achieve special effects.
func StencilFunc(xfunc GLEnum, ref int32, mask uint32) {
	gl.StencilFunc(uint32(xfunc), ref, mask)
}

// StencilFuncSeparate sets the front and/or back function and reference value
// for stencil testing.
//
// Stenciling enables and disables drawing on a per-pixel basis. It is typically
// used in multipass rendering to achieve special effects.
func StencilFuncSeparate(face, xfunc GLEnum, ref int32, mask uint32) {
	gl.StencilFuncSeparate(uint32(face), uint32(xfunc), ref, mask)
}

// StencilMask controls enabling and disabling of both the front and back
// writing of individual bits in the stencil planes.
//
// The StencilMaskSeparate function can set front and back stencil writemasks to
// different values.
func StencilMask(mask GLEnum) {
	gl.StencilMask(uint32(mask))
}

// StencilMaskSeparate controls enabling and disabling of front and/or back
// writing of individual bits in the stencil planes.
//
// The StencilMask function can set both, the front and back stencil writemasks
// to one value at the same time.
func StencilMaskSeparate(face GLEnum, mask uint32) {
	gl.StencilMaskSeparate(uint32(face), mask)
}

// StencilOp sets both the front and back-facing stencil test actions.
func StencilOp(fail, zfail, zpass GLEnum) {
	gl.StencilOp(uint32(fail), uint32(zfail), uint32(zpass))
}

// StencilOpSeparate sets the front and/or back-facing stencil test actions.
func StencilOpSeparate(face, fail, zfail, zpass GLEnum) {
	gl.StencilOpSeparate(uint32(face), uint32(fail), uint32(zfail), uint32(zpass))
}

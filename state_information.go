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

// GetActiveTexture returns a value for the passed parameter name.
func GetActiveTexture() GLEnum {
	var data int32
	gl.GetIntegerv(gl.ACTIVE_TEXTURE, &data)
	return GLEnum(data)
}

// GetAliasedLineWidthRange returns a value for the passed parameter name.
func GetAliasedLineWidthRange() [2]float32 {
	var data [2]float32
	gl.GetFloatv(gl.ALIASED_LINE_WIDTH_RANGE, &data[0])
	return data
}

// GetAliasedPointSizeRange returns a value for the passed parameter name.
func GetAliasedPointSizeRange() [2]float32 {
	var data [2]float32
	gl.GetFloatv(gl.ALIASED_POINT_SIZE_RANGE, &data[0])
	return data
}

// GetAlphaBits returns a value for the passed parameter name.
func GetAlphaBits() int32 {
	var data int32
	gl.GetIntegerv(gl.ALPHA_BITS, &data)
	return data
}

// GetArrayBufferBinding returns a value for the passed parameter name.
func GetArrayBufferBinding() Buffer {
	var data int32
	gl.GetIntegerv(gl.ARRAY_BUFFER_BINDING, &data)
	return Buffer(data)
}

// GetBlend returns a value for the passed parameter name.
func GetBlend() bool {
	var data bool
	gl.GetBooleanv(gl.BLEND, &data)
	return data
}

// GetBlendColor returns a value for the passed parameter name.
func GetBlendColor() [4]float32 {
	var data [4]float32
	gl.GetFloatv(gl.BLEND_COLOR, &data[0])
	return data
}

// GetBlendDstAlpha returns a value for the passed parameter name.
func GetBlendDstAlpha() GLEnum {
	var data int32
	gl.GetIntegerv(gl.BLEND_DST_ALPHA, &data)
	return GLEnum(data)
}

// GetBlendDstRGB returns a value for the passed parameter name.
func GetBlendDstRGB() GLEnum {
	var data int32
	gl.GetIntegerv(gl.BLEND_DST_RGB, &data)
	return GLEnum(data)
}

// GetBlendEquation returns a value for the passed parameter name.
func GetBlendEquation() GLEnum {
	var data int32
	gl.GetIntegerv(gl.BLEND_EQUATION, &data)
	return GLEnum(data)
}

// GetBlendEquationAlpha returns a value for the passed parameter name.
func GetBlendEquationAlpha() GLEnum {
	var data int32
	gl.GetIntegerv(gl.BLEND_EQUATION_ALPHA, &data)
	return GLEnum(data)
}

// GetBlendEquationRGB returns a value for the passed parameter name.
func GetBlendEquationRGB() GLEnum {
	var data int32
	gl.GetIntegerv(gl.BLEND_EQUATION_RGB, &data)
	return GLEnum(data)
}

// GetBlendSrcAlpha returns a value for the passed parameter name.
func GetBlendSrcAlpha() GLEnum {
	var data int32
	gl.GetIntegerv(gl.BLEND_SRC_ALPHA, &data)
	return GLEnum(data)
}

// GetBlendSrcRGB returns a value for the passed parameter name.
func GetBlendSrcRGB() GLEnum {
	var data int32
	gl.GetIntegerv(gl.BLEND_SRC_RGB, &data)
	return GLEnum(data)
}

// GetBlueBits returns a value for the passed parameter name.
func GetBlueBits() int32 {
	var data int32
	gl.GetIntegerv(gl.BLUE_BITS, &data)
	return data
}

// GetColorClearValue returns a value for the passed parameter name.
func GetColorClearValue() [4]float32 {
	var data [4]float32
	gl.GetFloatv(gl.COLOR_CLEAR_VALUE, &data[0])
	return data
}

// GetColorWritemask returns a value for the passed parameter name.
func GetColorWritemask() [4]bool {
	var data [4]bool
	gl.GetBooleanv(gl.COLOR_WRITEMASK, &data[0])
	return data
}

// TODO: GetCompressedTextureFormats

// GetCullFace returns a value for the passed parameter name.
func GetCullFace() bool {
	var data bool
	gl.GetBooleanv(gl.CULL_FACE, &data)
	return data
}

// GetCullFaceMode returns a value for the passed parameter name.
func GetCullFaceMode() GLEnum {
	var data int32
	gl.GetIntegerv(gl.CULL_FACE_MODE, &data)
	return GLEnum(data)
}

// GetCurrentProgram returns a value for the passed parameter name.
func GetCurrentProgram() Program {
	var data int32
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &data)
	return Program(data)
}

// GetDepthBits returns a value for the passed parameter name.
func GetDepthBits() int32 {
	var data int32
	gl.GetIntegerv(gl.DEPTH_BITS, &data)
	return data
}

// GetDepthClearValue returns a value for the passed parameter name.
func GetDepthClearValue() float32 {
	var data float32
	gl.GetFloatv(gl.DEPTH_CLEAR_VALUE, &data)
	return data
}

// GetDepthFunc returns a value for the passed parameter name.
func GetDepthFunc() GLEnum {
	var data int32
	gl.GetIntegerv(gl.DEPTH_FUNC, &data)
	return GLEnum(data)
}

// GetDepthRange returns a value for the passed parameter name.
func GetDepthRange() [2]float32 {
	var data [2]float32
	gl.GetFloatv(gl.DEPTH_RANGE, &data[0])
	return data
}

// GetDepthTest returns a value for the passed parameter name.
func GetDepthTest() bool {
	var data bool
	gl.GetBooleanv(gl.DEPTH_TEST, &data)
	return data
}

// GetDepthWritemask returns a value for the passed parameter name.
func GetDepthWritemask() bool {
	var data bool
	gl.GetBooleanv(gl.DEPTH_WRITEMASK, &data)
	return data
}

// GetDither returns a value for the passed parameter name.
func GetDither() bool {
	var data bool
	gl.GetBooleanv(gl.DITHER, &data)
	return data
}

// GetElementArrayBufferBinding returns a value for the passed parameter name.
func GetElementArrayBufferBinding() Buffer {
	var data int32
	gl.GetIntegerv(gl.ELEMENT_ARRAY_BUFFER_BINDING, &data)
	return Buffer(data)
}

// GetFramebufferBinding returns a value for the passed parameter name.
func GetFramebufferBinding() Framebuffer {
	var data int32
	gl.GetIntegerv(gl.FRAMEBUFFER_BINDING, &data)
	return Framebuffer(data)
}

// GetFrontFace returns a value for the passed parameter name.
func GetFrontFace() GLEnum {
	var data int32
	gl.GetIntegerv(gl.FRONT_FACE, &data)
	return GLEnum(data)
}

// GetGenerateMipmapHint returns a value for the passed parameter name.
func GetGenerateMipmapHint() GLEnum {
	var data int32
	gl.GetIntegerv(gl.GENERATE_MIPMAP_HINT, &data)
	return GLEnum(data)
}

// GetGreenBits returns a value for the passed parameter name.
func GetGreenBits() int32 {
	var data int32
	gl.GetIntegerv(gl.GREEN_BITS, &data)
	return data
}

// GetImplementationColorReadFormat returns a value for the passed parameter
// name.
func GetImplementationColorReadFormat() GLEnum {
	var data int32
	gl.GetIntegerv(gl.IMPLEMENTATION_COLOR_READ_FORMAT, &data)
	return GLEnum(data)
}

// GetImplementationColorReadType returns a value for the passed parameter name.
func GetImplementationColorReadType() GLEnum {
	var data int32
	gl.GetIntegerv(gl.IMPLEMENTATION_COLOR_READ_TYPE, &data)
	return GLEnum(data)
}

// GetLineWidth returns a value for the passed parameter name.
func GetLineWidth() float32 {
	var data float32
	gl.GetFloatv(gl.LINE_WIDTH, &data)
	return data
}

// GetMaxCombinedTextureImageUnits returns a value for the passed parameter
// name.
func GetMaxCombinedTextureImageUnits() int32 {
	var data int32
	gl.GetIntegerv(gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS, &data)
	return data
}

// GetMaxCubeMapTextureSize returns a value for the passed parameter name.
func GetMaxCubeMapTextureSize() int32 {
	var data int32
	gl.GetIntegerv(gl.MAX_CUBE_MAP_TEXTURE_SIZE, &data)
	return data
}

// GetMaxFragmentUniformVectors returns a value for the passed parameter name.
func GetMaxFragmentUniformVectors() int32 {
	var data int32
	gl.GetIntegerv(gl.MAX_FRAGMENT_UNIFORM_VECTORS, &data)
	return data
}

// GetMaxRenderbufferSize returns a value for the passed parameter name.
func GetMaxRenderbufferSize() int32 {
	var data int32
	gl.GetIntegerv(gl.MAX_RENDERBUFFER_SIZE, &data)
	return data
}

// GetMaxTextureImageUnits returns a value for the passed parameter name.
func GetMaxTextureImageUnits() int32 {
	var data int32
	gl.GetIntegerv(gl.MAX_TEXTURE_IMAGE_UNITS, &data)
	return data
}

// GetMaxTextureSize returns a value for the passed parameter name.
func GetMaxTextureSize() int32 {
	var data int32
	gl.GetIntegerv(gl.MAX_TEXTURE_SIZE, &data)
	return data
}

// GetMaxVaryingVectors returns a value for the passed parameter name.
func GetMaxVaryingVectors() int32 {
	var data int32
	gl.GetIntegerv(gl.MAX_VARYING_VECTORS, &data)
	return data
}

// GetMaxVertexAttribs returns a value for the passed parameter name.
func GetMaxVertexAttribs() int32 {
	var data int32
	gl.GetIntegerv(gl.MAX_VERTEX_ATTRIBS, &data)
	return data
}

// GetMaxVertexTextureImageUnits returns a value for the passed parameter name.
func GetMaxVertexTextureImageUnits() int32 {
	var data int32
	gl.GetIntegerv(gl.MAX_VERTEX_TEXTURE_IMAGE_UNITS, &data)
	return data
}

// GetMaxVertexUniformVectors returns a value for the passed parameter name.
func GetMaxVertexUniformVectors() int32 {
	var data int32
	gl.GetIntegerv(gl.MAX_VERTEX_UNIFORM_VECTORS, &data)
	return data
}

// GetMaxViewportDims returns a value for the passed parameter name.
func GetMaxViewportDims() [2]int32 {
	var data [2]int32
	gl.GetIntegerv(gl.MAX_VIEWPORT_DIMS, &data[0])
	return data
}

// GetPackAlignment returns a value for the passed parameter name.
func GetPackAlignment() int32 {
	var data int32
	gl.GetIntegerv(gl.PACK_ALIGNMENT, &data)
	return data
}

// GetPolygonOffsetFactor returns a value for the passed parameter name.
func GetPolygonOffsetFactor() float32 {
	var data float32
	gl.GetFloatv(gl.POLYGON_OFFSET_FACTOR, &data)
	return data
}

// GetPolygonOffsetFill returns a value for the passed parameter name.
func GetPolygonOffsetFill() bool {
	var data bool
	gl.GetBooleanv(gl.POLYGON_OFFSET_FILL, &data)
	return data
}

// GetPolygonOffsetUnits returns a value for the passed parameter name.
func GetPolygonOffsetUnits() float32 {
	var data float32
	gl.GetFloatv(gl.POLYGON_OFFSET_UNITS, &data)
	return data
}

// GetRedBits returns a value for the passed parameter name.
func GetRedBits() int32 {
	var data int32
	gl.GetIntegerv(gl.RED_BITS, &data)
	return data
}

// GetRenderbufferBinding returns a value for the passed parameter name.
func GetRenderbufferBinding() Renderbuffer {
	var data int32
	gl.GetIntegerv(gl.RENDERBUFFER_BINDING, &data)
	return Renderbuffer(data)
}

// TODO: GetRenderer

// GetSampleBuffers returns a value for the passed parameter name.
func GetSampleBuffers() int32 {
	var data int32
	gl.GetIntegerv(gl.SAMPLE_BUFFERS, &data)
	return data
}

// GetSampleCoverageInvert returns a value for the passed parameter name.
func GetSampleCoverageInvert() bool {
	var data bool
	gl.GetBooleanv(gl.SAMPLE_COVERAGE_INVERT, &data)
	return data
}

// GetSampleCoverageValue returns a value for the passed parameter name.
func GetSampleCoverageValue() float32 {
	var data float32
	gl.GetFloatv(gl.SAMPLE_COVERAGE_VALUE, &data)
	return data
}

// GetSamples returns a value for the passed parameter name.
func GetSamples() int32 {
	var data int32
	gl.GetIntegerv(gl.SAMPLES, &data)
	return data
}

// GetScissorBox returns a value for the passed parameter name.
func GetScissorBox() [4]int32 {
	var data [4]int32
	gl.GetIntegerv(gl.SCISSOR_BOX, &data[0])
	return data
}

// GetScissorTest returns a value for the passed parameter name.
func GetScissorTest() bool {
	var data bool
	gl.GetBooleanv(gl.SCISSOR_TEST, &data)
	return data
}

// TODO: GetShadingLanguageVersion

// GetStencilBackFail returns a value for the passed parameter name.
func GetStencilBackFail() GLEnum {
	var data int32
	gl.GetIntegerv(gl.STENCIL_BACK_FAIL, &data)
	return GLEnum(data)
}

// GetStencilBackFunc returns a value for the passed parameter name.
func GetStencilBackFunc() GLEnum {
	var data int32
	gl.GetIntegerv(gl.STENCIL_BACK_FUNC, &data)
	return GLEnum(data)
}

// GetStencilBackPassDepthFail returns a value for the passed parameter name.
func GetStencilBackPassDepthFail() GLEnum {
	var data int32
	gl.GetIntegerv(gl.STENCIL_BACK_PASS_DEPTH_FAIL, &data)
	return GLEnum(data)
}

// GetStencilBackPassDepthPass returns a value for the passed parameter name.
func GetStencilBackPassDepthPass() GLEnum {
	var data int32
	gl.GetIntegerv(gl.STENCIL_BACK_PASS_DEPTH_PASS, &data)
	return GLEnum(data)
}

// GetStencilBackRef returns a value for the passed parameter name.
func GetStencilBackRef() int32 {
	var data int32
	gl.GetIntegerv(gl.STENCIL_BACK_REF, &data)
	return data
}

// GetStencilBackValueMask returns a value for the passed parameter name.
func GetStencilBackValueMask() uint32 {
	var data int32
	gl.GetIntegerv(gl.STENCIL_BACK_VALUE_MASK, &data)
	return uint32(data)
}

// GetStencilBackWritemask returns a value for the passed parameter name.
func GetStencilBackWritemask() uint32 {
	var data int32
	gl.GetIntegerv(gl.STENCIL_BACK_WRITEMASK, &data)
	return uint32(data)
}

// GetStencilBits returns a value for the passed parameter name.
func GetStencilBits() int32 {
	var data int32
	gl.GetIntegerv(gl.STENCIL_BITS, &data)
	return data
}

// GetStencilClearValue returns a value for the passed parameter name.
func GetStencilClearValue() int32 {
	var data int32
	gl.GetIntegerv(gl.STENCIL_CLEAR_VALUE, &data)
	return data
}

// GetStencilFail returns a value for the passed parameter name.
func GetStencilFail() GLEnum {
	var data int32
	gl.GetIntegerv(gl.STENCIL_FAIL, &data)
	return GLEnum(data)
}

// GetStencilFunc returns a value for the passed parameter name.
func GetStencilFunc() GLEnum {
	var data int32
	gl.GetIntegerv(gl.STENCIL_FUNC, &data)
	return GLEnum(data)
}

// GetStencilPassDepthFail returns a value for the passed parameter name.
func GetStencilPassDepthFail() GLEnum {
	var data int32
	gl.GetIntegerv(gl.STENCIL_PASS_DEPTH_FAIL, &data)
	return GLEnum(data)
}

// GetStencilPassDepthPass returns a value for the passed parameter name.
func GetStencilPassDepthPass() GLEnum {
	var data int32
	gl.GetIntegerv(gl.STENCIL_PASS_DEPTH_PASS, &data)
	return GLEnum(data)
}

// GetStencilRef returns a value for the passed parameter name.
func GetStencilRef() int32 {
	var data int32
	gl.GetIntegerv(gl.STENCIL_REF, &data)
	return data
}

// GetStencilTest returns a value for the passed parameter name.
func GetStencilTest() bool {
	var data bool
	gl.GetBooleanv(gl.STENCIL_TEST, &data)
	return data
}

// GetStencilValueMask returns a value for the passed parameter name.
func GetStencilValueMask() uint32 {
	var data int32
	gl.GetIntegerv(gl.STENCIL_VALUE_MASK, &data)
	return uint32(data)
}

// GetStencilWritemask returns a value for the passed parameter name.
func GetStencilWritemask() uint32 {
	var data int32
	gl.GetIntegerv(gl.STENCIL_WRITEMASK, &data)
	return uint32(data)
}

// GetSubpixelBits returns a value for the passed parameter name.
func GetSubpixelBits() int32 {
	var data int32
	gl.GetIntegerv(gl.SUBPIXEL_BITS, &data)
	return data
}

// GetTextureBinding2D returns a value for the passed parameter name.
func GetTextureBinding2D() Texture {
	var data int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &data)
	return Texture(data)
}

// GetTextureBindingCubeMap returns a value for the passed parameter name.
func GetTextureBindingCubeMap() Texture {
	var data int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_CUBE_MAP, &data)
	return Texture(data)
}

// GetUnpackAlignment returns a value for the passed parameter name.
func GetUnpackAlignment() int32 {
	var data int32
	gl.GetIntegerv(gl.UNPACK_ALIGNMENT, &data)
	return data
}

// TODO: GetVendor

// TODO: GetVersion

// GetViewport returns a value for the passed parameter name.
func GetViewport() [4]int32 {
	var data [4]int32
	gl.GetIntegerv(gl.VIEWPORT, &data[0])
	return data
}

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

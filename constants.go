package gogl

import "github.com/go-gl/gl/v2.1/gl"

// GLEnum is used for enums.
type GLEnum uint32

// Clearing buffers
//
// Constants passed to Clear to clear buffer masks.
const (
	// GLDepthBufferBit is passed to Clear to clear the current depth buffer.
	GLDepthBufferBit GLEnum = gl.DEPTH_BUFFER_BIT
	// GLStencilBufferBit is passed to Clear to clear the current stencil
	// buffer.
	GLStencilBufferBit GLEnum = gl.STENCIL_BUFFER_BIT
	// GLColorBufferBit is passed to Clear to clear the current color buffer.
	GLColorBufferBit GLEnum = gl.COLOR_BUFFER_BIT
)

// Rendering primitives
//
// Constants passed to DrawElements or DrawArrays to specify what kind of
// primitive to render.
const (
	// GLPoints is passed to DrawElements or DrawArrays to draw single points.
	GLPoints GLEnum = gl.POINTS
	// GLLines is passed to DrawElements or DrawArrays to draw lines. Each
	// vertex connects to the one after it.
	GLLines GLEnum = gl.LINES
	// GLLineLoop is passed to DrawElements or DrawArrays to draw lines. Each
	// set of two vertices is treated as a separate line segment.
	GLLineLoop GLEnum = gl.LINE_LOOP
	// GLLineStrip is passed to DrawElements or DrawArrays to draw a connected
	// group of line segments from the first vertex to the last.
	GLLineStrip GLEnum = gl.LINE_STRIP
	// GLTriangles is passed to DrawElements or DrawArrays to draw triangles.
	// Each set of three vertices creates a separate triangle.
	GLTriangles GLEnum = gl.TRIANGLES
	// GLTriangleStrip is passed to DrawElements or DrawArrays to draw a
	// connected group of triangles.
	GLTriangleStrip GLEnum = gl.TRIANGLE_STRIP
	// GLTriangleFan is passed to DrawElements or DrawArrays to draw a connected
	// group of triangles. Each vertex connects to the previous and the first
	// vertex in the fan.
	GLTriangleFan GLEnum = gl.TRIANGLE_FAN
)

// Blending modes
//
// Constants passed to BlendFunc or BlendFuncSeparate to specify the blending
// mode (for both, RGB and alpha, or separately).
const (
	// GLZero is passed to BlendFunc or BlendFuncSeparate to turn off a
	// component.
	GLZero GLEnum = gl.ZERO
	// GLOne is passed to BlendFunc or BlendFuncSeparate to turn on a component.
	GLOne GLEnum = gl.ONE
	// GLSrcColor is passed to BlendFunc or BlendFuncSeparate to multiply a
	// component by the source elements color.
	GLSrcColor GLEnum = gl.SRC_COLOR
	// GLOneMinusSrcColor is passed to BlendFunc or BlendFuncSeparate to
	// multiply a component by one minus the source elements color.
	GLOneMinusSrcColor GLEnum = gl.ONE_MINUS_SRC_COLOR
	// GLSrcAlpha is passed to BlendFunc or BlendFuncSeparate to multiply a
	// component by the source's alpha.
	GLSrcAlpha GLEnum = gl.SRC_ALPHA
	// GLOneMinusSrcAlpha is passed to BlendFunc or BlendFuncSeparate to
	// multiply a component by one minus the source's alpha.
	GLOneMinusSrcAlpha GLEnum = gl.ONE_MINUS_SRC_ALPHA
	// GLDstAlpha is passed to BlendFunc or BlendFuncSeparate to multiply a
	// componentby the destination's alpha.
	GLDstAlpha GLEnum = gl.DST_ALPHA
	// GLOneMinusDstAlpha is passed to BlendFunc or BlendFuncSeparate to
	// multiplya component by one minus the destination's alpha.
	GLOneMinusDstAlpha GLEnum = gl.ONE_MINUS_DST_ALPHA
	// GLDstColor is passed to BlendFunc or BlendFuncSeparate to multiply a
	// component by the destination's color.
	GLDstColor GLEnum = gl.DST_COLOR
	// GLOneMinusDstColor is passed to BlendFunc or BlendFuncSeparate to
	// multiply a component by one minus the destination's color.
	GLOneMinusDstColor GLEnum = gl.ONE_MINUS_DST_COLOR
	// GLSrcAlphaSaturate is passed to BlendFunc or BlendFuncSeparate to
	// multiply a component by the minimum of source's alpha or one minus the
	// destination's alpha.
	GLSrcAlphaSaturate GLEnum = gl.SRC_ALPHA_SATURATE
	// GLConstantColor is passed to BlendFunc or BlendFuncSeparate to specify a
	// constant color blend function.
	GLConstantColor GLEnum = gl.CONSTANT_COLOR
	// GLOneMinusConstantColor is passed to BlendFunc or BlendFuncSeparate to
	// specify one minus a constant color blend function.
	GLOneMinusConstantColor GLEnum = gl.ONE_MINUS_CONSTANT_COLOR
	// GLConstantAlpha is passed to BlendFunc or BlendFuncSeparate to specify a
	// constant alpha blend function.
	GLConstantAlpha GLEnum = gl.CONSTANT_ALPHA
	// GLOneMinusConstantAlpha is passed to BlendFunc or BlendFuncSeparate to
	// specify one minus a constant alpha blend function.
	GLOneMinusConstantAlpha GLEnum = gl.ONE_MINUS_CONSTANT_ALPHA
)

// Blending equations
//
// Constants passed to BlendEquation or BlendEquationSeparate to control how the
// blending is calculated (for both, RGB and alpha, or separately).
const (
	// GLFuncAdd is passed to BlendEquation or BlendEquationSeparate to set an
	// addition blend function.
	GLFuncAdd GLEnum = gl.FUNC_ADD
	// GLFuncSubtract is passed to BlendEquation or BlendEquationSeparate to
	// specify a subtraction blend function (source - destination).
	GLFuncSubtract GLEnum = gl.FUNC_SUBTRACT
	// GLFuncReverseSubtract is passed to BlendEquation or BlendEquationSeparate
	// to specify a reverse subtraction blend function (destination - source).
	GLFuncReverseSubtract GLEnum = gl.FUNC_REVERSE_SUBTRACT
)

// Getting GL parameter information
//
// Constants passed to GetParameter to specify what information to return.
const (
	// GLBlendEquation is passed to GetParameter to get the current RGB blend
	// function.
	GLBlendEquation GLEnum = gl.BLEND_EQUATION
	// GLBlendEquationRGB is passed to GetParameter to get the current RGB blend
	// function. Same as GLBlendEquation.
	GLBlendEquationRGB GLEnum = gl.BLEND_EQUATION_RGB
	// GLBlendEquationAlpha is passed to GetParameter to get the current alpha
	// blend function. Same as GLBlendEquation.
	GLBlendEquationAlpha GLEnum = gl.BLEND_EQUATION_ALPHA
	// GLBlendDstRGB is passed to GetParameter to get the current destination
	// RGB blend function.
	GLBlendDstRGB GLEnum = gl.BLEND_DST_RGB
	// GLBlendSrcRGB is passed to GetParameter to get the current source RGB
	// blend function.
	GLBlendSrcRGB GLEnum = gl.BLEND_SRC_RGB
	// GLBlendDstAlpha is passed to GetParameter to get the current destination
	// alpha blend function.
	GLBlendDstAlpha GLEnum = gl.BLEND_DST_ALPHA
	// GLBlendSrcAlpha is passed to GetParameter to get the current source alpha
	// blend function.
	GLBlendSrcAlpha GLEnum = gl.BLEND_SRC_ALPHA
	// GLBlendColor is passed to GetParameter to return the current blend color.
	GLBlendColor GLEnum = gl.BLEND_COLOR
	// GLArrayBufferBinding is passed to GetParameter to get the array buffer
	// binding.
	GLArrayBufferBinding GLEnum = gl.ARRAY_BUFFER_BINDING
	// GLElementArrayBufferBinding is passed to GetParameter to get the current
	// element array buffer.
	GLElementArrayBufferBinding GLEnum = gl.ELEMENT_ARRAY_BUFFER_BINDING
	// GLLineWidth is passed to GetParameter to get the current LineWidth (set
	// by the LineWidth method).
	GLLineWidth GLEnum = gl.LINE_WIDTH
	// GLAliasedPointSizeRange is passed to GetParameter to get the current size
	// of a point drawn with GLPoints.
	GLAliasedPointSizeRange GLEnum = gl.ALIASED_POINT_SIZE_RANGE
	// GLAliasedLineWidthRange is passed to GetParameter to get the range of
	// available widths for a line. Returns a length-2 array with the low value
	// at 0, and high at 1.
	GLAliasedLineWidthRange GLEnum = gl.ALIASED_LINE_WIDTH_RANGE
	// GLCullFaceMode is passed to GetParameter to get the current value of
	// CullFace. Should return GLFront, GLBack, or GLFrontAndBack.
	GLCullFaceMode GLEnum = gl.CULL_FACE_MODE
	// GLFrontFace is passed to GetParameter to determine the current value of
	// FrontFace. Should return GLCW or GLCCW.
	GLFrontFace GLEnum = gl.FRONT_FACE
	// GLDepthRange is passed to GetParameter to return a length-2 array of
	// floats givingthe current depth range.
	GLDepthRange GLEnum = gl.DEPTH_RANGE
	// GLDepthWritemask is passed to GetParameter to determine if the depth
	// write mask is enabled.
	GLDepthWritemask GLEnum = gl.DEPTH_WRITEMASK
	// GLDepthClearValue is passed to GetParameter to determine the current
	// depth clear value.
	GLDepthClearValue GLEnum = gl.DEPTH_CLEAR_VALUE
	// GLDepthFunc is passed to GetParameter to get the current depth function.
	// Returns GLNever, GLAlways, GLLEss, GLEqual, GLLEqual, GLGreater,
	// GLGEqual, or GLNotEqual.
	GLDepthFunc GLEnum = gl.DEPTH_FUNC
	// GLStencilClearValue is passed to GetParameter to get the value the
	// stencil will be cleared to.
	GLStencilClearValue GLEnum = gl.STENCIL_CLEAR_VALUE
	// GLStencilFunc is passed to GetParameter to get the current stencil
	// function. Returns GLNever, GLAlways, GLLess, GLEqual, GLLEqual,
	// GLGreater, GLGEqual, or GLNotEqual.
	GLStencilFunc GLEnum = gl.STENCIL_FUNC
	// GLStencilFail is passed to GetParameter to get the current stencil fail
	// function. Should return GLKeep, GLReplace, GLIncr, GLDecr, GLInvert,
	// GLIncrWrap, or GLDecrWrap.
	GLStencilFail GLEnum = gl.STENCIL_FAIL
	// GLStencilPassDepthFail is passed to GetParameter to get the current
	// stencil fail function should the depth buffer test fail. Should return
	// GLKeep, GLReplace, GLIncr, GLDecr, GLInvert, GLIncrWrap, or GLDecrWrap.
	GLStencilPassDepthFail GLEnum = gl.STENCIL_PASS_DEPTH_FAIL
	// GLStencilPassDepthPass is passed to GetParameter to get the current
	// stencil fail function should the depth buffer test pass. Should return
	// GLKeep, GLReplace, GLIncr, GLDecr, GLInvert, GLIncrWrap, or GLDecrWrap.
	GLStencilPassDepthPass GLEnum = gl.STENCIL_PASS_DEPTH_PASS
	// GLStencilRef is passed to GetParameter to get the reference value used
	// for stencil tests.
	GLStencilRef               GLEnum = gl.STENCIL_REF
	GLStencilValueMask         GLEnum = gl.STENCIL_VALUE_MASK
	GLStencilWritemask         GLEnum = gl.STENCIL_WRITEMASK
	GLStencilBackFunc          GLEnum = gl.STENCIL_BACK_FUNC
	GLStencilBackFail          GLEnum = gl.STENCIL_BACK_FAIL
	GLStencilBackPassDepthFail GLEnum = gl.STENCIL_BACK_PASS_DEPTH_FAIL
	GLStencilBackPassDepthPass GLEnum = gl.STENCIL_BACK_PASS_DEPTH_PASS
	GLStencilBackRef           GLEnum = gl.STENCIL_BACK_REF
	GLStencilBackValueMask     GLEnum = gl.STENCIL_BACK_VALUE_MASK
	GLStencilBackWritemask     GLEnum = gl.STENCIL_BACK_WRITEMASK
	// GLViewport returns an int32 array with four elements for the current
	// viewport dimensions.
	GLViewport GLEnum = gl.VIEWPORT
	// GLScissorBox returns an int32 array with four elements for the current
	// scissor box dimensions.
	GLScissorBox                    GLEnum = gl.SCISSOR_BOX
	GLColorClearValue               GLEnum = gl.COLOR_CLEAR_VALUE
	GLColorWritemask                GLEnum = gl.COLOR_WRITEMASK
	GLUnpackAlignment               GLEnum = gl.UNPACK_ALIGNMENT
	GLPackAlignment                 GLEnum = gl.PACK_ALIGNMENT
	GLMaxTextureSize                GLEnum = gl.MAX_TEXTURE_SIZE
	GLMaxViewportDims               GLEnum = gl.MAX_VIEWPORT_DIMS
	GLSubpixelBits                  GLEnum = gl.SUBPIXEL_BITS
	GLRedBits                       GLEnum = gl.RED_BITS
	GLGreenBits                     GLEnum = gl.GREEN_BITS
	GLBlueBits                      GLEnum = gl.BLUE_BITS
	GLAlphaBits                     GLEnum = gl.ALPHA_BITS
	GLDepthBits                     GLEnum = gl.DEPTH_BITS
	GLStencilBits                   GLEnum = gl.STENCIL_BITS
	GLPolygonOffsetUnits            GLEnum = gl.POLYGON_OFFSET_UNITS
	GLPolygonOffsetFactor           GLEnum = gl.POLYGON_OFFSET_FACTOR
	GLTextureBinding2D              GLEnum = gl.TEXTURE_BINDING_2D
	GLSampleBuffers                 GLEnum = gl.SAMPLE_BUFFERS
	GLSamples                       GLEnum = gl.SAMPLES
	GLSampleCoverageValue           GLEnum = gl.SAMPLE_COVERAGE_VALUE
	GLSampleCoverageInvert          GLEnum = gl.SAMPLE_COVERAGE_INVERT
	GLCompressedTextureFormats      GLEnum = gl.COMPRESSED_TEXTURE_FORMATS
	GLVendor                        GLEnum = gl.VENDOR
	GLRenderer                      GLEnum = gl.RENDERER
	GLVersion                       GLEnum = gl.VERSION
	GLImplementationColorReadType   GLEnum = gl.IMPLEMENTATION_COLOR_READ_TYPE
	GLImplementationColorReadFormat GLEnum = gl.IMPLEMENTATION_COLOR_READ_FORMAT
)

// Buffers
//
// Constants passed to BufferData, BufferSubData, BindBuffer, or
// GetBufferParameter.
const (
	// GLStaticDraw is passed to BufferData as a hint about whether the contents
	// of the buffer are likely to be used often and not change often.
	GLStaticDraw GLEnum = gl.STATIC_DRAW
	// GLStreamDraw is passed to BufferData as a hint about whether the contents
	// of the buffer are likely to not be used often.
	GLStreamDraw GLEnum = gl.STREAM_DRAW
	// GLDynamicDraw is passed to BufferData as a hint about whether the
	// contents of the buffer are likely to be used often and change often.
	GLDynamicDraw GLEnum = gl.DYNAMIC_DRAW
	// GLArrayBuffer is passed to BindBuffer or BufferData to specify the type
	// of buffer being used.
	GLArrayBuffer GLEnum = gl.ARRAY_BUFFER
	// GLElementArrayBuffer is passed to BindBuffer or BufferData to specify the
	// type of buffer being used.
	GLElementArrayBuffer GLEnum = gl.ELEMENT_ARRAY_BUFFER
	// GLBufferSize is passed to GetBufferParameter to get a buffer's size.
	GLBufferSize GLEnum = gl.BUFFER_SIZE
	// GLBufferUsage is passed to GetBufferParameter to get the hint for the
	// buffer passed in when it was created.
	GLBufferUsage GLEnum = gl.BUFFER_USAGE
)

// Vertex attributes
//
// Constants passed to GetVertexAttrib.
const (
	// GLCurrentVertexAttrib is passed to GetVertexAttrib to read back the
	// current vertex attribute.
	GLCurrentVertexAttrib            GLEnum = gl.CURRENT_VERTEX_ATTRIB
	GLVertexAttribArrayEnabled       GLEnum = gl.VERTEX_ATTRIB_ARRAY_ENABLED
	GLVertexAttribArraySize          GLEnum = gl.VERTEX_ATTRIB_ARRAY_SIZE
	GLVertexAttribArrayStride        GLEnum = gl.VERTEX_ATTRIB_ARRAY_STRIDE
	GLVertexAttribArrayType          GLEnum = gl.VERTEX_ATTRIB_ARRAY_TYPE
	GLVertexAttribArrayNormalized    GLEnum = gl.VERTEX_ATTRIB_ARRAY_NORMALIZED
	GLVertexAttribArrayPointer       GLEnum = gl.VERTEX_ATTRIB_ARRAY_POINTER
	GLVertexAttribArrayBufferBinding GLEnum = gl.VERTEX_ATTRIB_ARRAY_BUFFER_BINDING
)

// Culling
//
// Constants passed to CullFace.
const (
	// GLCullFace is passed to Enable/Disable to turn on/off culling. Can also
	// be used with GetParameter to find the current culling method.
	GLCullFace GLEnum = gl.CULL_FACE
	// GLFront is passed to CullFace to specify that only front faces should be
	// culled.
	GLFront GLEnum = gl.FRONT
	// GLBack is passed to CullFace to specify that only back faces should be
	// culled.
	GLBack GLEnum = gl.BACK
	// GLFrontAndBack is passed to CullFace to specify that front and back faces
	// should be culled.
	GLFrontAndBack GLEnum = gl.FRONT_AND_BACK
)

// Enabling and disabling
//
// Constants passed to Enable or Disable.
const (
	// GLBlend is passed to Enable/Disable to turn on/off blending. Can also be
	// used with GetParameter to find the current blending method.
	GLBlend GLEnum = gl.BLEND
	// GLDepthTest is passed to Enable/Disable to turn on/off the depth test.
	// Can also be used with GetParameter to query the depth test.
	GLDepthTest GLEnum = gl.DEPTH_TEST
	// GLDither is passed to Enable/Disable to turn on/off dithering. Can also
	// be used with GetParameter to find the current dithering method.
	GLDither GLEnum = gl.DITHER
	// GLPolygonOffsetFill is passed to Enable/Disable to turn on/off the
	// polygon offset. Useful for rendering hidden-line images, decals, and/or
	// solids with highlighted edges. Can also be used with GetParameter to
	// query the polygon offset fill.
	GLPolygonOffsetFill GLEnum = gl.POLYGON_OFFSET_FILL
	// GLSampleAlphaToCoverage is passed to Enable/Disable to turn on/off the
	// alpha to coverage. Used in multi-sampling alpha channels.
	GLSampleAlphaToCoverage GLEnum = gl.SAMPLE_ALPHA_TO_COVERAGE
	// GLSampleCoverage is passed to Enable/Disable to turn on/off the sample
	// coverage. Used in multi-sampling.
	GLSampleCoverage GLEnum = gl.SAMPLE_COVERAGE
	// GLScissorTest is passed to Enable/Disable to turn on/off the scissor
	// test. Can also be used with GetParameter to query the scissor test.
	GLScissorTest GLEnum = gl.SCISSOR_TEST
	// GLStencilTest is passed to Enable/Disable to turn on/off the stencil
	// test. Can also be used with GetParameter to query the stencil test.
	GLStencilTest GLEnum = gl.STENCIL_TEST
)

// Errors
//
// Constants returned from GetError.
const (
	// GLNoError is returned from GetError.
	GLNoError GLEnum = gl.NO_ERROR
	// GLInvalidEnum is returned from GetError.
	GLInvalidEnum GLEnum = gl.INVALID_ENUM
	// GLInvalidValue is returned from GetError.
	GLInvalidValue GLEnum = gl.INVALID_VALUE
	// GLInvalidOperation is returned from GetError.
	GLInvalidOperation GLEnum = gl.INVALID_OPERATION
	// GLOutOfMemory is returned from GetError.
	GLOutOfMemory GLEnum = gl.OUT_OF_MEMORY
	// GLContextLost is returned from GetError.
	GLContextLost GLEnum = gl.CONTEXT_LOST
)

// Front face directions
//
// Constants passed to FrontFace.
const (
	// GLCW is passed to FrontFace to specify the front face of a polygon is
	// drawn in the clockwise direction.
	GLCW GLEnum = gl.CW
	// GLCCW is passed to FrontFace to specify the front face of a polygon is
	// drawn in the counter clockwise direction.
	GLCCW GLEnum = gl.CCW
)

// Hints
//
// Constants passed to Hint.
const (
	// GLDontCare means that there is no preference for this behavior.
	GLDontCare GLEnum = gl.DONT_CARE
	// GLFastest means that the most efficient behavior should be used.
	GLFastest GLEnum = gl.FASTEST
	// GLNicest means that the most correct or the highest quality option should
	// be used.
	GLNicest GLEnum = gl.NICEST
	// GLGenerateMipmapHint is a hint for the quality of filtering when
	// generating mimap images with GenerateMipmap.
	GLGenerateMipmapHint = gl.GENERATE_MIPMAP_HINT
)

// Data types
const (
	GLInt8    GLEnum = gl.BYTE
	GLUInt8   GLEnum = gl.UNSIGNED_BYTE
	GLInt16   GLEnum = gl.SHORT
	GLUInt16  GLEnum = gl.UNSIGNED_SHORT
	GLInt32   GLEnum = gl.INT
	GLUInt32  GLEnum = gl.UNSIGNED_INT
	GLFloat32 GLEnum = gl.FLOAT
)

// Pixel formats
const (
	GLDepthComponent GLEnum = gl.DEPTH_COMPONENT
	GLAlpha          GLEnum = gl.ALPHA
	GLRGB            GLEnum = gl.RGB
	GLRGBA           GLEnum = gl.RGBA
	GLLuminance      GLEnum = gl.LUMINANCE
	GLLuminanceAlpha GLEnum = gl.LUMINANCE_ALPHA
)

// Pixel types
const (
	GLUInt164444 GLEnum = gl.UNSIGNED_SHORT_4_4_4_4
	GLUInt165551 GLEnum = gl.UNSIGNED_SHORT_5_5_5_1
	GLUInt16565  GLEnum = gl.UNSIGNED_SHORT_5_6_5
)

// Shaders
//
// Constants passed to CreateShader or GetShaderParameter.
const (
	// GLFragmentShader is passed to CreateShader to define a fragment shader.
	GLFragmentShader GLEnum = gl.FRAGMENT_SHADER
	// GLVertexShader is passed to CreateShader to define a vertex shader.
	GLVertexShader GLEnum = gl.VERTEX_SHADER
	// GLCompileStatus is passed to GetShaderParameter to get the status of the
	// compilation. Returns false if the shader was not compiled. You can then
	// query GetShaderInfoLog to find the exact error.
	GLCompileStatus GLEnum = gl.COMPILE_STATUS
	// GLDeleteStatus is passed to GetShaderParameter to determine if a shader
	// was deleted via DeleteShader. Returns true if it was, false otherwise.
	GLDeleteStatus GLEnum = gl.DELETE_STATUS
	// GLLinkStatus is passed to GetProgramParameter after calling LinkProgram
	// to determine if a program was linked correctly. Returns false if there
	// were erors. Use GetProgramInfoLog to find the exact error.
	GLLinkStatus GLEnum = gl.LINK_STATUS
	// GLValidateStatus is passed to GetProgramParameter after calling
	// ValidateProgram to determine if it is valid. Returns false if errors were
	// found.
	GLValidateStatus GLEnum = gl.VALIDATE_STATUS
	// GLAttachedShaders is passed to GetProgramParameter after calling
	// AttachShader to determine if the shader was attached correctly. Returns
	// false if errors occurred.
	GLAttachedShaders GLEnum = gl.ATTACHED_SHADERS
	// GLActiveAttributes is passed to GetProgramParameter to get the number of
	// attributes active in a program.
	GLActiveAttributes GLEnum = gl.ACTIVE_ATTRIBUTES
	// GLActiveUniforms is passed to GetProgramParameter to get the number of
	// uniforms active in a program.
	GLActiveUniforms GLEnum = gl.ACTIVE_UNIFORMS
	// GLMaxVertexAttribs is the maximum number of entries possible in the
	// vertex attribute list.
	GLMaxVertexAttribs             GLEnum = gl.MAX_VERTEX_ATTRIBS
	GLMaxVertexUniformVectors      GLEnum = gl.MAX_VERTEX_UNIFORM_VECTORS
	GLMaxVaryingVectors            GLEnum = gl.MAX_VARYING_VECTORS
	GLMaxCombinedTextureImageUnits GLEnum = gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS
	GLMaxVertexTextureImageUnits   GLEnum = gl.MAX_VERTEX_TEXTURE_IMAGE_UNITS
	// GLMaxTextureImageUnits is the implementation dependent number of maximum
	// texture units. At least 8.
	GLMaxTextureImageUnits      GLEnum = gl.MAX_TEXTURE_IMAGE_UNITS
	GLMaxFragmentUniformVectors GLEnum = gl.MAX_FRAGMENT_UNIFORM_VECTORS
	GLShaderType                GLEnum = gl.SHADER_TYPE
	GLShadingLanguageVersion    GLEnum = gl.SHADING_LANGUAGE_VERSION
	GLCurrentProgram            GLEnum = gl.CURRENT_PROGRAM
)

// Depth or stencil tests
//
// Constants passed to DepthFunc or StencilFunc.
const (
	// GLNever is passed to DepthFunc or StencilFunc to specify depth or stencil
	// tests will never pass, i.e. nothing will be drawn.
	GLNever GLEnum = gl.NEVER
	// GLLess is passed to DepthFunc or StencilFunc to specify depth or stencil
	// tests will pass if the new depth value is less than the stored value.
	GLLess GLEnum = gl.LESS
	// GLEqual is passed to DepthFunc or StencilFunc to specify depth or stencil
	// tests will pass if the new depth value is equal to the stored value.
	GLEqual GLEnum = gl.EQUAL
	// GLLEqual is passed to DepthFunc or StencilFunc to specify depth or
	// stencil tests will pass if the new depth value is less than or equal to
	// the stored value.
	GLLEqual GLEnum = gl.LEQUAL
	// GLGreater is passed to DepthFunc or StencilFunc to specify depth or
	// stencil tests will pass if the new depth value is greater than the stored
	// value.
	GLGreater GLEnum = gl.GREATER
	// GLNotEqual is passed to DepthFunc or StencilFunc to specify depth or
	// stencil tests will pass if the new depth value is not equal to the stored
	// value.
	GLNotEqual GLEnum = gl.NOTEQUAL
	// GLGEqual is passed to DepthFunc or StencilFunc to specify depth or
	// stencil tests will pass if the new depth value is greater than or equal
	// to the stored value.
	GLGEqual GLEnum = gl.GEQUAL
	// GLAlways is passed to DepthFunc or StencilFunc to specify depth or
	// stencil tests will always pass, i.e. pixels will be drawnin the order
	// they are drawn.
	GLAlways GLEnum = gl.ALWAYS
)

// Stencil actions
//
// Constants passed to StencilOp.
const (
	GLKeep     GLEnum = gl.KEEP
	GLReplace  GLEnum = gl.REPLACE
	GLIncr     GLEnum = gl.INCR
	GLDecr     GLEnum = gl.DECR
	GLInvert   GLEnum = gl.INVERT
	GLIncrWrap GLEnum = gl.INCR_WRAP
	GLDecrWrap GLEnum = gl.DECR_WRAP
)

// Textures
//
// Constants passed to TexParameteri, TexParameterf, BindTexture, TexImage2D,
// and others.
const (
	GLNearest                 GLEnum = gl.NEAREST
	GLLinear                  GLEnum = gl.LINEAR
	GLNearestMipmapNearest    GLEnum = gl.NEAREST_MIPMAP_NEAREST
	GLLinearMipmapNearest     GLEnum = gl.LINEAR_MIPMAP_NEAREST
	GLNearestMipmapLinear     GLEnum = gl.NEAREST_MIPMAP_LINEAR
	GLLinearMipmapLinear      GLEnum = gl.LINEAR_MIPMAP_LINEAR
	GLTextureMagFilter        GLEnum = gl.TEXTURE_MAG_FILTER
	GLTextureMinFilter        GLEnum = gl.TEXTURE_MIN_FILTER
	GLTextureWrapS            GLEnum = gl.TEXTURE_WRAP_S
	GLTextureWrapT            GLEnum = gl.TEXTURE_WRAP_T
	GLTexture2D               GLEnum = gl.TEXTURE_2D
	GLTexture                 GLEnum = gl.TEXTURE
	GLTextureCubeMap          GLEnum = gl.TEXTURE_CUBE_MAP
	GLTextureBindingCubeMap   GLEnum = gl.TEXTURE_BINDING_CUBE_MAP
	GLTextureCubeMapPositiveX GLEnum = gl.TEXTURE_CUBE_MAP_POSITIVE_X
	GLTextureCubeMapNegativeX GLEnum = gl.TEXTURE_CUBE_MAP_NEGATIVE_X
	GLTextureCubeMapPositiveY GLEnum = gl.TEXTURE_CUBE_MAP_POSITIVE_Y
	GLTextureCubeMapNegativeY GLEnum = gl.TEXTURE_CUBE_MAP_NEGATIVE_Y
	GLTextureCubeMapPositiveZ GLEnum = gl.TEXTURE_CUBE_MAP_POSITIVE_Z
	GLTextureCubeMapNegativeZ GLEnum = gl.TEXTURE_CUBE_MAP_NEGATIVE_Z
	GLMaxCubeMapTextureSize   GLEnum = gl.MAX_CUBE_MAP_TEXTURE_SIZE
	// GLTexture0 is a texture unit.
	GLTexture0 GLEnum = gl.TEXTURE0
	// GLTexture1 is a texture unit.
	GLTexture1 GLEnum = gl.TEXTURE1
	// GLTexture2 is a texture unit.
	GLTexture2 GLEnum = gl.TEXTURE2
	// GLTexture3 is a texture unit.
	GLTexture3 GLEnum = gl.TEXTURE3
	// GLTexture4 is a texture unit.
	GLTexture4 GLEnum = gl.TEXTURE4
	// GLTexture5 is a texture unit.
	GLTexture5 GLEnum = gl.TEXTURE5
	// GLTexture6 is a texture unit.
	GLTexture6 GLEnum = gl.TEXTURE6
	// GLTexture7 is a texture unit.
	GLTexture7 GLEnum = gl.TEXTURE7
	// GLTexture8 is a texture unit.
	GLTexture8 GLEnum = gl.TEXTURE8
	// GLTexture9 is a texture unit.
	GLTexture9 GLEnum = gl.TEXTURE9
	// GLTexture10 is a texture unit.
	GLTexture10 GLEnum = gl.TEXTURE10
	// GLTexture11 is a texture unit.
	GLTexture11 GLEnum = gl.TEXTURE11
	// GLTexture12 is a texture unit.
	GLTexture12 GLEnum = gl.TEXTURE12
	// GLTexture13 is a texture unit.
	GLTexture13 GLEnum = gl.TEXTURE13
	// GLTexture14 is a texture unit.
	GLTexture14 GLEnum = gl.TEXTURE14
	// GLTexture15 is a texture unit.
	GLTexture15 GLEnum = gl.TEXTURE15
	// GLTexture16 is a texture unit.
	GLTexture16 GLEnum = gl.TEXTURE16
	// GLTexture17 is a texture unit.
	GLTexture17 GLEnum = gl.TEXTURE17
	// GLTexture18 is a texture unit.
	GLTexture18 GLEnum = gl.TEXTURE18
	// GLTexture19 is a texture unit.
	GLTexture19 GLEnum = gl.TEXTURE19
	// GLTexture20 is a texture unit.
	GLTexture20 GLEnum = gl.TEXTURE20
	// GLTexture21 is a texture unit.
	GLTexture21 GLEnum = gl.TEXTURE21
	// GLTexture22 is a texture unit.
	GLTexture22 GLEnum = gl.TEXTURE22
	// GLTexture23 is a texture unit.
	GLTexture23 GLEnum = gl.TEXTURE23
	// GLTexture24 is a texture unit.
	GLTexture24 GLEnum = gl.TEXTURE24
	// GLTexture25 is a texture unit.
	GLTexture25 GLEnum = gl.TEXTURE25
	// GLTexture26 is a texture unit.
	GLTexture26 GLEnum = gl.TEXTURE26
	// GLTexture27 is a texture unit.
	GLTexture27 GLEnum = gl.TEXTURE27
	// GLTexture28 is a texture unit.
	GLTexture28 GLEnum = gl.TEXTURE28
	// GLTexture29 is a texture unit.
	GLTexture29 GLEnum = gl.TEXTURE29
	// GLTexture30 is a texture unit.
	GLTexture30 GLEnum = gl.TEXTURE30
	// GLTexture31 is a texture unit.
	GLTexture31 GLEnum = gl.TEXTURE31
	// GLActiveTexture is the current active texture unit.
	GLActiveTexture  GLEnum = gl.ACTIVE_TEXTURE
	GLRepeat         GLEnum = gl.REPEAT
	GLClampToEdge    GLEnum = gl.CLAMP_TO_EDGE
	GLMirroredRepeat GLEnum = gl.MIRRORED_REPEAT
)

// Uniform types
const (
	GLFloatVec2   GLEnum = gl.FLOAT_VEC2
	GLFloatVec3   GLEnum = gl.FLOAT_VEC3
	GLFloatVec4   GLEnum = gl.FLOAT_VEC4
	GLIntVec2     GLEnum = gl.INT_VEC2
	GLIntVec3     GLEnum = gl.INT_VEC3
	GLIntVec4     GLEnum = gl.INT_VEC4
	GLBool        GLEnum = gl.BOOL
	GLBoolVec2    GLEnum = gl.BOOL_VEC2
	GLBoolVec3    GLEnum = gl.BOOL_VEC3
	GLBoolVec4    GLEnum = gl.BOOL_VEC4
	GLFloatMat2   GLEnum = gl.FLOAT_MAT2
	GLFloatMat3   GLEnum = gl.FLOAT_MAT3
	GLFloatMat4   GLEnum = gl.FLOAT_MAT4
	GLSampler2D   GLEnum = gl.SAMPLER_2D
	GLSamplerCube GLEnum = gl.SAMPLER_CUBE
)

// Shader precision-specified types
const (
	GLLowFloat    GLEnum = gl.LOW_FLOAT
	GLMediumFloat GLEnum = gl.MEDIUM_FLOAT
	GLHighFloat   GLEnum = gl.HIGH_FLOAT
	GLLowInt      GLEnum = gl.LOW_INT
	GLMediumInt   GLEnum = gl.MEDIUM_INT
	GLHighInt     GLEnum = gl.HIGH_INT
)

// Framebuffers and renderbuffers
const (
	GLFramebuffer                             GLEnum = gl.FRAMEBUFFER
	GLRenderbuffer                            GLEnum = gl.RENDERBUFFER
	GLRGBA4                                   GLEnum = gl.RGBA4
	GLRGB5A1                                  GLEnum = gl.RGB5_A1
	GLRGB565                                  GLEnum = gl.RGB565
	GLDepthComponent16                        GLEnum = gl.DEPTH_COMPONENT16
	GLStencilIndex8                           GLEnum = gl.STENCIL_INDEX8
	GLDepthStencil                            GLEnum = gl.DEPTH_STENCIL
	GLRenderbufferWidth                       GLEnum = gl.RENDERBUFFER_WIDTH
	GLRenderbufferHeight                      GLEnum = gl.RENDERBUFFER_HEIGHT
	GLRenderbufferInternalFormat              GLEnum = gl.RENDERBUFFER_INTERNAL_FORMAT
	GLRenderbufferRedSize                     GLEnum = gl.RENDERBUFFER_RED_SIZE
	GLRenderbufferGreenSize                   GLEnum = gl.RENDERBUFFER_GREEN_SIZE
	GLRenderbufferBlueSize                    GLEnum = gl.RENDERBUFFER_BLUE_SIZE
	GLRenderbufferAlphaSize                   GLEnum = gl.RENDERBUFFER_ALPHA_SIZE
	GLRenderbufferDepthSize                   GLEnum = gl.RENDERBUFFER_DEPTH_SIZE
	GLRenderbufferStencilSize                 GLEnum = gl.RENDERBUFFER_STENCIL_SIZE
	GLFramebufferAttachmentObjectType         GLEnum = gl.FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE
	GLFramebufferAttachmentObjectName         GLEnum = gl.FRAMEBUFFER_ATTACHMENT_OBJECT_NAME
	GLFramebufferAttachmentTextureLevel       GLEnum = gl.FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL
	GLFramebufferAttachmentTextureCubeMapFace GLEnum = gl.FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE
	GLColorAttachment0                        GLEnum = gl.COLOR_ATTACHMENT0
	GLDepthAttachment                         GLEnum = gl.DEPTH_ATTACHMENT
	GLStencilAttachment                       GLEnum = gl.STENCIL_ATTACHMENT
	GLDepthStencilAttachment                  GLEnum = gl.DEPTH_STENCIL_ATTACHMENT
	GLNone                                    GLEnum = gl.NONE
	GLFramebufferComplete                     GLEnum = gl.FRAMEBUFFER_COMPLETE
	GLFramebufferIncompleteAttachment         GLEnum = gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT
	GLFramebufferIncompleteMissingAttachment  GLEnum = gl.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT
	GLFramebufferUnsupported                  GLEnum = gl.FRAMEBUFFER_UNSUPPORTED
	GLFramebufferBinding                      GLEnum = gl.FRAMEBUFFER_BINDING
	GLRenderbufferBinding                     GLEnum = gl.RENDERBUFFER_BINDING
	GLMaxRenderbufferSize                     GLEnum = gl.MAX_RENDERBUFFER_SIZE
	GLInvalidFramebufferOperation             GLEnum = gl.INVALID_FRAMEBUFFER_OPERATION
)

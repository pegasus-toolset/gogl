package gogl

// Buffer represents an opaque buffer object storing data such as vertices or
// colors.
type Buffer uint32

// Framebuffer represents a collection of buffers that serve as a rendering
// destination.
type Framebuffer uint32

// Renderbuffer represents a buffer that can contain an image, or can be source
// or target of an rendering operation.
type Renderbuffer uint32

// Texture represents an opaque texture object providing storage and state for
// texturing operations.
type Texture uint32

// Program is a combination of two compiled Shaders consisting of a vertex
// shader and a fragment shader (both written in GLSL). To create a Program,
// call the CreateProgram function. After attaching the shader programs using
// AttachShader, you link them into a usable program.
type Program uint32

// Shader can either be a vertex or a fragment shader. A Program requires both
// types of shaders.
type Shader uint32

// UniformLocation represents the location of a uniform variable in a shader
// program.
type UniformLocation int32

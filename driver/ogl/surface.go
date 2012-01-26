package ogl

import (
	"image"

	"github.com/Agon/baukasten"

	gl "github.com/chsc/gogl/gl33"
)

var SurfaceFragmentShader = "#version 120\nvarying vec2 f_texcoord;\nuniform sampler2D texture;\nvoid main(void) {\ngl_FragColor = texture2D(texture, f_texcoord);\n}"
var SurfaceVertexShader = "#version 120\nattribute vec2 coord2d;\nattribute vec2 texcoord;\nvarying vec2 f_texcoord;\nuniform mat4 mvp;\nvoid main(void) {\ngl_Position = mvp * vec4(coord2d, 0.0, 1.0);\nf_texcoord = texcoord;\n}"

var SurfaceTriangles = []float32{
	0, 0, 0,
	1, 0, 0,
	1, 1, 0,
	1, 1, 0,
	0, 1, 0,
	0, 0, 0,
}

type Surface struct {
	// Scale
	scaleX float32
	scaleY float32
	// Rotate
	angle float32
	// Size
	width, height float32
	// OpenGL
	texture *Texture
	vbo     *VertexBufferObject
	program *Program
	vs      *Shader
	fs      *Shader

	// Locations
	coord2d  *AttributeLocation
	texcoord *AttributeLocation
	mvp      *UniformLocation
	texId    *UniformLocation
}

var ScreenHeight = 480
var ScreenWidth = 640

func (d *Driver) OpenSurface(name string) (baukasten.Surface, error) {
	img, err := baukasten.OpenImage(name)
	if err != nil {
		return nil, err
	}
	return d.LoadSurface(img)
}

func (d *Driver) LoadSurface(img image.Image) (baukasten.Surface, error) {
	var err error
	s := &Surface{}
	// Texture
	s.texture, err = d.LoadTexture(img)
	if err != nil {
		return nil, err
	}
	// Generate triangles
	width := float32(img.Bounds().Dx())
	height := float32(img.Bounds().Dx())

	x := width / 2
	y := height / 2
	triangles := []float32{
		-x, -y, 0, 0,
		x, -y, 1, 0,
		x, y, 1, 1,
		x, y, 1, 1,
		-x, y, 0, 1,
		-x, -y, 0, 0,
	} // TODO We need to adjust or scale the vertices to width/height

	// VertexBufferObject
	s.vbo = NewVBO()
	s.vbo.BufferData(triangles)

	// Shaderprogram
	s.program = NewProgram()
	s.vs, err = LoadShader(SurfaceVertexShader, VertexShaderType)
	if err != nil {
		return nil, err
	}
	s.fs, err = LoadShader(SurfaceFragmentShader, FragmentShaderType)
	if err != nil {
		return nil, err
	}
	s.program.AttachShaders(s.vs, s.fs)
	err = s.program.Link()
	if err != nil {
		return nil, err
	}

	// AttributeLocations
	s.coord2d, err = s.program.GetAttributeLocation("coord2d")
	if err != nil {
		return nil, err
	}
	s.texcoord, err = s.program.GetAttributeLocation("texcoord")
	if err != nil {
		return nil, err
	}
	// UniformLocations
	s.mvp, err = s.program.GetUniformLocation("mvp")
	if err != nil {
		return nil, err
	}
	s.texId, err = s.program.GetUniformLocation("texture")
	if err != nil {
		return nil, err
	}
	s.width = width
	s.height = height
	s.scaleX = 1
	s.scaleY = 1
	return s, nil
}

func (s *Surface) Delete() {
	s.texture.Delete()
	s.vbo.Delete()
	s.program.Delete()
	s.vs.Delete()
	s.fs.Delete()
}

func (s *Surface) Scale(x, y float32) {
	s.scaleX = x
	s.scaleY = y
}

func (s *Surface) Rotate(angle float32) {
	s.angle = angle
}

func (s *Surface) Draw(x, y float32) {
	s.program.Use()
	// projection->view->model(translation->rotation)
	model := baukasten.TranslationMatrix(x, y, 0)
	model = model.Multiplied(baukasten.ScaleMatrix(s.scaleX, s.scaleY, 1))
	if s.angle != 0 {
		model = model.Multiplied(baukasten.RotationMatrix(s.angle, baukasten.Vector3{0, 0, 1}))
	}
	projection := baukasten.Ortho2D(0, float32(ScreenWidth), float32(ScreenHeight), 0)
	matrix := projection.Multiplied(model)
	s.mvp.UniformMatrix4fv(1, false, matrix.Transposed())

	s.texId.Uniform1i(0)

	s.coord2d.Enable()
	defer s.coord2d.Disable()

	s.vbo.Bind()
	defer s.vbo.Unbind()
	s.coord2d.AttribPointer(2, gl.FLOAT, false, 16, gl.Offset(nil, 0))

	s.texture.Bind()
	defer s.texture.Unbind()

	s.texcoord.Enable()
	defer s.texcoord.Disable()
	s.texcoord.AttribPointer(2, gl.FLOAT, false, 16, gl.Offset(nil, 8))

	s.vbo.DrawArrays(0, 6)
}

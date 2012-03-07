package ogl

import (
	"image"
	"image/color"

	gl "github.com/chsc/gogl/gl33"

	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/geometry"
)

const (
	SurfaceVertexShader   = "#version 120\nattribute vec2 coord2d;\nattribute vec2 texcoord;\nvarying vec2 f_texcoord;\nvarying vec4 f_color;\nuniform mat4 mvp;\nuniform vec4 v_color;\nvoid main(void) {\ngl_Position = mvp * vec4(coord2d, 0.0, 1.0);\nf_texcoord = texcoord;\nf_color = v_color;\n}"
	SurfaceFragmentShader = "#version 120\nvarying vec2 f_texcoord;\nvarying vec4 f_color;\nuniform sampler2D texture;\nvoid main(void) {\n" +
		"gl_FragColor = f_color * texture2D(texture, f_texcoord);\n}"
)

type Surface struct {
	// Scale
	scaleX float32
	scaleY float32
	// Rotate
	angle float32
	// Size
	width, height float32
	// Color
	r, g, b, a float32
	// OpenGL
	texture *Texture
	vbo     *VertexBufferObject
	program *Program
	vs      *Shader
	fs      *Shader

	// Locations
	coord2d    *AttributeLocation
	texcoord   *AttributeLocation
	colorCoord *UniformLocation
	mvp        *UniformLocation
	texId      *UniformLocation

	// driver
	driver *Driver
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
	height := float32(img.Bounds().Dy())

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
	s.colorCoord, err = s.program.GetUniformLocation("v_color")
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
	s.driver = d
	s.r, s.g, s.b, s.a = 1, 1, 1, 1
	return s, nil
}

func (s *Surface) Delete() {
	s.texture.Delete()
	s.vbo.Delete()
	s.program.Delete()
	s.vs.Delete()
	s.fs.Delete()
}

func (s *Surface) Color() color.Color {
	return geometry.ConvertFColor(s.r, s.g, s.b, s.a)
}

func (s *Surface) SetColor(c color.Color) {
	s.r, s.g, s.b, s.a = geometry.ConvertColorF(c)
}

func (s *Surface) Scale(x, y float32) {
	s.scaleX = x
	s.scaleY = y
}

func (s *Surface) GetScale() (float32, float32) {
	return s.scaleX, s.scaleY
}

func (s *Surface) Rotate(angle float32) {
	s.angle = angle
}

func (s *Surface) Rotation() float32 {
	return s.angle
}

func (s *Surface) Draw(x, y float32) {
	s.program.Use()
	// projection->view->model(translation->rotation)
	model := geometry.TranslationMatrix(x, y, 0)
	model = model.Mul(geometry.ScaleMatrix(s.scaleX, s.scaleY, 1))
	if s.angle != 0 {
		model = model.Mul(geometry.RotationMatrix(s.angle, geometry.Vector3{0, 0, 1}))
	}
	matrix := s.driver.Camera().Get().Mul(model)
	s.mvp.UniformMatrix4fv(1, false, matrix.Transposed())

	s.texId.Uniform1i(0)
	s.colorCoord.Uniform4f(s.r, s.g, s.b, s.a)

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

func (s *Surface) DrawRec(rec geometry.Rectanglef) {
	s.program.Use()
	// projection->view->model(translation->rotation)
	pos := rec.Center()
	model := geometry.TranslationMatrix(pos.X, pos.Y, 0)
	model = model.Mul(geometry.ScaleMatrix(s.scaleX, s.scaleY, 1))
	if s.angle != 0 {
		model = model.Mul(geometry.RotationMatrix(s.angle, geometry.Vector3{0, 0, 1}))
	}
	matrix := s.driver.Camera().Get().Mul(model)
	s.mvp.UniformMatrix4fv(1, false, matrix.Transposed())

	s.texId.Uniform1i(0)
	s.colorCoord.Uniform4f(s.r, s.g, s.b, s.a)

	s.coord2d.Enable()
	defer s.coord2d.Disable()

	vertices := []float32{
		rec.Min.X, rec.Min.Y,
		rec.Max.X, rec.Min.Y,
		rec.Max.X, rec.Max.Y,
		rec.Max.X, rec.Max.Y,
		rec.Min.X, rec.Max.Y,
		rec.Min.X, rec.Min.Y,
	}

	s.coord2d.AttribPointer(2, gl.FLOAT, false, 0, gl.Pointer(&vertices[0]))

	s.vbo.Bind()
	defer s.vbo.Unbind()

	s.texture.Bind()
	defer s.texture.Unbind()

	s.texcoord.Enable()
	defer s.texcoord.Disable()
	s.texcoord.AttribPointer(2, gl.FLOAT, false, 16, gl.Offset(nil, 8))

	s.vbo.DrawArrays(0, 6)
}

func (s *Surface) DrawRegionRec(src geometry.Rectanglef, dest geometry.Rectanglef) {
	s.program.Use()
	// projection->view->model(translation->rotation)
	pos := dest.Center()
	model := geometry.TranslationMatrix(pos.X, pos.Y, 0)
	model = model.Mul(geometry.ScaleMatrix(s.scaleX, s.scaleY, 1))
	if s.angle != 0 {
		model = model.Mul(geometry.RotationMatrix(s.angle, geometry.Vector3{0, 0, 1}))
	}
	matrix := s.driver.Camera().Get().Mul(model)
	s.mvp.UniformMatrix4fv(1, false, matrix.Transposed())

	s.texId.Uniform1i(0)
	s.colorCoord.Uniform4f(s.r, s.g, s.b, s.a)

	s.coord2d.Enable()
	defer s.coord2d.Disable()

	vertices := []float32{
		dest.Min.X, dest.Min.Y,
		dest.Max.X, dest.Min.Y,
		dest.Max.X, dest.Max.Y,
		dest.Max.X, dest.Max.Y,
		dest.Min.X, dest.Max.Y,
		dest.Min.X, dest.Min.Y,
	}

	s.coord2d.AttribPointer(2, gl.FLOAT, false, 0, gl.Pointer(&vertices[0]))

	texCoords := []float32{
		src.Min.X, src.Min.Y,
		src.Max.X, src.Min.Y,
		src.Max.X, src.Max.Y,
		src.Max.X, src.Max.Y,
		src.Min.X, src.Max.Y,
		src.Min.X, src.Min.Y,
	}

	s.texture.Bind()
	defer s.texture.Unbind()

	s.texcoord.Enable()
	defer s.texcoord.Disable()
	s.texcoord.AttribPointer(2, gl.FLOAT, false, 0, gl.Pointer(&texCoords[0]))

	s.vbo.DrawArrays(0, 6)
}

package gl

import (
	"errors"
	"fmt"
	"io/ioutil"

	gl "github.com/chsc/gogl/gl33"
)

type ShaderType byte

const (
	VertexShader ShaderType = iota
	FragmentShader
)

var (
	ErrUnknownShader = errors.New("Unknown shader type.")
)

type Shader struct {
	id gl.Uint
}

func OpenShader(name string, shaderType ShaderType) (*Shader, error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return LoadShader(string(b), shaderType)
}

func LoadShader(data string, shaderType ShaderType) (*Shader, error) {
	var id gl.Uint
	switch shaderType {
	case VertexShader:
		id = gl.CreateShader(gl.VERTEX_SHADER)
	case FragmentShader:
		id = gl.CreateShader(gl.FRAGMENT_SHADER)
	default:
		return nil, ErrUnknownShader
	}
	shader := &Shader{id: id}
	src := gl.GLStringArray(data)
	length := gl.Int(-1)
	gl.ShaderSource(shader.id, gl.Sizei(1), &src[0], &length)
	gl.GLStringArrayFree(src)
	return shader, nil
}

func (s *Shader) Delete() {
	gl.DeleteShader(s.id)
}

func (s *Shader) Id() gl.Uint {
	return s.id
}

func (s *Shader) Compile() error {
	var compileOk gl.Int
	gl.CompileShader(s.id)
	gl.GetShaderiv(s.id, gl.COMPILE_STATUS, &compileOk)
	if compileOk == 0 {
		errNum := gl.GetError()
		return errors.New(fmt.Sprintf("Error in vertex shader: %d\n", errNum))
	}
	return nil
}

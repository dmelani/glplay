package scene

import (
	"github.com/dmelani/glplay/models"
	"github.com/go-gl/gl"
	"fmt"
)

type Scene struct {
	width           int
	height          int
	rectangle       *models.Rectangle
	shader_program   gl.Program
}

func Create(width int, height int) (scene *Scene) {
	scene = &Scene{
		width:           width,
		height:          height,
		rectangle:       models.NewRectangle(-0.5, -0.5, 1.0, 1.0),
		shader_program: gl.CreateProgram(),
	}

	return scene
}

func (s *Scene) Init() {
	fragment_shader_source := `
			#version 330 core

			out vec4 finalColor;

			void main() {
				//set every drawn pixel to white
				finalColor = vec4(1.0, 1.0, 1.0, 1.0);
			}`

	vertex_shader_source := `
			#version 150

			in vec3 vert;

			void main() {
				// does not alter the vertices at all
				gl_Position = vec4(vert, 1);
			}`

	fs := gl.CreateShader(gl.FRAGMENT_SHADER)
	fs.Source(fragment_shader_source)
	fs.Compile()
	if fs.Get(gl.COMPILE_STATUS) == 0 {
		fmt.Println("Fragment shader compilation failed:")
	}
	fmt.Println(fs.GetInfoLog())

	vs := gl.CreateShader(gl.VERTEX_SHADER)
	vs.Source(vertex_shader_source)
	vs.Compile()

	gl.ClearColor(0.0, 0.0, 0.0, 0.0)

	gl.Viewport(0, 0, s.width, s.height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-1, 1, -1, 1, -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func (s *Scene) Render() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	s.rectangle.Render()
}

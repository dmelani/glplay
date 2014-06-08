package scene

import (
	"github.com/go-gl/gl"
	"github.com/dmelani/glplay/models"
)

type Scene struct {
	width int
	height int
	rectangle *models.Rectangle
	vertex_shader string
	fragment_shader string
}

func Create(width int, height int) (scene *Scene) {
	scene = &Scene{
		width: width,
		height: height,
		rectangle: models.NewRectangle(-0.5, -0.5, 1.0, 1.0),
		vertex_shader: `
			#version 150

			in vec3 vert;

			void main() {
				// does not alter the vertices at all
				gl_Position = vec4(vert, 1);
			}
		`,
		fragment_shader: `
			#version 150

			out vec4 finalColor;

			void main() {
				//set every drawn pixel to white
				finalColor = vec4(1.0, 1.0, 1.0, 1.0);
			}
		`,

	}

	return scene
}

func (s *Scene) Init() {
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

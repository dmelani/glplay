package main

import (
	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
	"github.com/dmelani/glplay/scene"
)

const (
	Title = "Derp"
	Width = 800
	Height = 600
)

func main() {
	glfw.Init()
	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.WindowNoResize, 1)

	glfw.OpenWindow(Width, Height, 0, 0, 0, 0, 16, 0, glfw.Windowed)
	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(Title)

	gl.Init()
	s := scene.Create(Width, Height)
	s.Init()

	for glfw.WindowParam(glfw.Opened) == 1 {
		s.Render()
		glfw.SwapBuffers()
	}
}

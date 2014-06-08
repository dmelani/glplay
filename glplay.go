package main

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
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

	glfw.WindowHint(glfw.DepthBits, 16)

	window, err := glfw.CreateWindow(Width, Height, Title, nil, nil);
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	gl.Init()
	s := scene.Create(Width, Height)
	s.Init()

	for !window.ShouldClose() {
		s.Render()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

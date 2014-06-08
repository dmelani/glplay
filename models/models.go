package models

import (
	"github.com/go-gl/gl"
)

/*
type Model interface {
	Render()
}
*/

type Rectangle struct {
	x float32
	y float32
	w float32
	h float32
}

func NewRectangle(x float32, y float32, w float32, h float32) (rectangle *Rectangle) {
	return &Rectangle{x, y, w, h,}
}

func (r *Rectangle) Render() {
	gl.PushMatrix()
	gl.Translatef(r.x, r.y, 0.0)

	gl.Begin(gl.QUADS)
	gl.Color3f(1.0, 1.0, 1.0)
	gl.Vertex2f(0.0, 0.0)
	gl.Vertex2f(0.0, r.h)
	gl.Vertex2f(r.w, r.h)
	gl.Vertex2f(r.w, 0.0)
	gl.End()

	gl.PopMatrix()
}


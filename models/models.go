package models

import (
	gl "github.com/chsc/gogl/gl21"
)

/*
type Model interface {
	Render()
}
*/

type Rectangle struct {
	x gl.Float
	y gl.Float
	w gl.Float
	h gl.Float
}

func CreateRectangle(x float32, y float32, w float32, h float32) (rectangle *Rectangle) {
	return &Rectangle{gl.Float(x), gl.Float(y), gl.Float(w), gl.Float(h),}
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


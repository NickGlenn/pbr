package surface

import (
	"github.com/hunterloftis/pbr/geom"
	"github.com/hunterloftis/pbr/material"
)

// Surface is an interface to all surface types (sphere, cube, etc).
// A surface is anything that can be intersected by a Ray.
// TODO: rename? solid.Surface? solid.Box?
type Surface interface {
	Intersect(*geom.Ray3) Hit
	At(point geom.Vector3) (normal geom.Direction, material *material.Sample)
	Box() *Box
	Center() geom.Vector3
	Emits() bool
}

package collision

import (
	"math"
)

type CollisionAreaType int

const (
	TYPE_RECTANGLE CollisionAreaType = iota
	TYPE_CIRCLE
)

type CollisionArea struct {
	Type                        CollisionAreaType
	X, Y, Width, Height, Radius float64
}

type Collision interface {
	GetCollisionArea() CollisionArea
}

func Rectangle(x, y, width, height float64) CollisionArea {
	c := CollisionArea{}
	c.Type = TYPE_RECTANGLE
	c.X = x
	c.Y = y
	c.Width = width
	c.Height = height
	return c
}

func Circle(x, y, radius float64) CollisionArea {
	c := CollisionArea{}
	c.Type = TYPE_CIRCLE
	c.X = x
	c.Y = y
	c.Radius = radius
	return c
}

func Collides(c1, c2 CollisionArea) bool {
	switch c1.Type {
	case TYPE_CIRCLE:
		if c2.Type == TYPE_CIRCLE {
			return c2c(c1, c2)
		}

		return r2c(c2, c1)

	case TYPE_RECTANGLE:
		if c2.Type == TYPE_RECTANGLE {
			return r2r(c1, c2)
		}
	}

	return r2c(c1, c2)
}

func r2r(c1, c2 CollisionArea) bool {
	return (c1.X < c2.X+c2.Width &&
		c1.X+c1.Width > c2.X &&
		c1.Y < c2.Y+c2.Height &&
		c1.Y+c1.Height > c2.Y)
}

func r2c(c1, c2 CollisionArea) bool {
	c2.X -= c2.Radius
	c2.Y -= c2.Radius
	c2.Height = c2.Radius * 2
	c2.Width = c2.Radius * 2
	return r2r(c1, c2)
}

func c2c(c1, c2 CollisionArea) bool {
	distance := math.Sqrt(math.Pow(c2.X-c1.X, 2) + math.Pow(c2.Y-c1.Y, 2))

	return distance < c1.Radius+c2.Radius
}

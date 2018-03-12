// Copyright (c) 2018, Randall C. O'Reilly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"math/rand"
	"testing"
)

func TestRect(t *testing.T) {
	parent := NewViewport2D(400, 400)
	parent.SetThisName(parent, "vp1")
	vpfill := parent.AddNewChildNamed(KiT_Viewport2DFill, "vpfill").(*Viewport2DFill)
	vpfill.SetProp("fill", "#FFF")

	rect1 := vpfill.AddNewChildNamed(KiT_Rect, "rect1").(*Rect)
	rect1.SetProp("fill", "#008800")
	rect1.SetProp("stroke", "#0000FF")
	rect1.SetProp("stroke-width", 5.0)
	// rect1.SetProp("stroke-linejoin", "round")
	rect1.Pos = Point2D{10, 10}
	rect1.Size = Size2D{100, 100}
	parent.Render2DRoot()
	// parent.SavePNG("test_rect.png")
}

func TestShapesAll(t *testing.T) {
	width := 400
	height := 400
	parent := NewViewport2D(width, height)
	parent.SetThisName(parent, "vp1")
	vpfill := parent.AddNewChildNamed(KiT_Viewport2DFill, "vpfill").(*Viewport2DFill)
	vpfill.SetProp("fill", "#FFF")

	rect1 := vpfill.AddNewChildNamed(KiT_Rect, "rect1").(*Rect)
	rect1.SetProp("fill", "#008800")
	rect1.SetProp("stroke", "#0000FF")
	rect1.SetProp("stroke-width", 5.0)
	rect1.Pos = Point2D{10, 10}
	rect1.Size = Size2D{100, 100}

	circle1 := vpfill.AddNewChildNamed(KiT_Circle, "circle1").(*Circle)
	circle1.SetProp("fill", "none") // todo: need to process
	circle1.SetProp("stroke", "#CC0000")
	circle1.SetProp("stroke-width", 2.0)
	circle1.Pos = Point2D{200, 100}
	circle1.Radius = 40

	ellipse1 := circle1.AddNewChildNamed(KiT_Ellipse, "ellipse1").(*Ellipse)
	ellipse1.SetProp("fill", "#55000055")
	ellipse1.SetProp("stroke", "#880000")
	ellipse1.SetProp("stroke-width", 2.0)
	ellipse1.Pos = Point2D{100, 100}
	ellipse1.Radii = Size2D{80, 20}

	line1 := vpfill.AddNewChildNamed(KiT_Line, "line1").(*Line)
	line1.SetProp("stroke", "#888800")
	line1.SetProp("stroke-width", 5.0)
	line1.Start = Point2D{100, 100}
	line1.End = Point2D{150, 200}

	polyline1 := vpfill.AddNewChildNamed(KiT_Polyline, "polyline1").(*Polyline)
	polyline1.SetProp("stroke", "#888800")
	polyline1.SetProp("stroke-width", 4.0)

	for i := 0; i < 10; i++ {
		x1 := rand.Float64() * float64(width)
		y1 := rand.Float64() * float64(height)
		polyline1.Points = append(polyline1.Points, Point2D{x1, y1})
	}

	polygon1 := vpfill.AddNewChildNamed(KiT_Polygon, "polygon1").(*Polygon)
	polygon1.SetProp("fill", "#55005555")
	polygon1.SetProp("stroke", "#888800")
	polygon1.SetProp("stroke-width", 4.0)

	for i := 0; i < 10; i++ {
		x1 := rand.Float64() * float64(width)
		y1 := rand.Float64() * float64(height)
		polygon1.Points = append(polygon1.Points, Point2D{x1, y1})
	}

	parent.Render2DRoot()
	parent.SavePNG("test_shape_all.png")
}
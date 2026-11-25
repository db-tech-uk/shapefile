package shp

import "testing"

func polygonsAreSame(a, b *Polygon) bool {
	return bboxesAreSame(&a.BBox, &b.BBox) &&
		a.NumberOfParts == b.NumberOfParts &&
		a.NumberOfPoints == b.NumberOfPoints &&
		allInt32sAreSame(a.Parts, b.Parts) &&
		allPointsAreSame(a.Points, b.Points)
}

func TestPolygon(t *testing.T) {
	expectIn(t, "../test_files/polygon", &Expected{
		Header: &Header{
			FileLength: 220,
			ShapeType:  TypePolygon,
			BBox: BBox{
				MinX: 0,
				MinY: 0,
				MaxX: 5,
				MaxY: 5,
			},
		},
		Shapes: []Shape{
			&Polygon{
				BBox: BBox{
					MinX: 0,
					MinY: 0,
					MaxX: 5,
					MaxY: 5,
				},
				NumberOfParts:  2,
				NumberOfPoints: 10,
				Parts:          []int32{0, 5},
				Points: []Point{
					{0, 0},
					{0, 5},
					{5, 5},
					{5, 0},
					{0, 0},
					{1, 1},
					{1, 4},
					{4, 4},
					{4, 1},
					{1, 1},
				},
			},
			&Polygon{
				BBox: BBox{
					MinX: 0,
					MinY: 0,
					MaxX: 5,
					MaxY: 5,
				},
				NumberOfParts:  1,
				NumberOfPoints: 4,
				Parts:          []int32{0},
				Points: []Point{
					{0, 0},
					{5, 5},
					{5, 0},
					{0, 0},
				},
			},
		},
	})
}

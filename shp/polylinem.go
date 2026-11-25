package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PolylineM is a Polyline with optional measure (M) data. Missing M data is
// specified as NaN.
type PolylineM struct {
	BBox           BBox
	NumberOfParts  int32
	NumberOfPoints int32
	Parts          []int32
	Points         []Point
	MData
}

// Type ...
func (p *PolylineM) Type() ShapeType {
	return TypePolylineM
}

func readPolylineM(r io.Reader, cl int32) (*PolylineM, error) {
	var p PolylineM

	// BBox
	if err := binary.Read(r, binary.LittleEndian, &p.BBox); err != nil {
		return nil, err
	}

	// NumberOfParts
	if err := binary.Read(r, binary.LittleEndian, &p.NumberOfParts); err != nil {
		return nil, err
	}

	// NumberOfPoints
	if err := binary.Read(r, binary.LittleEndian, &p.NumberOfPoints); err != nil {
		return nil, err
	}

	minLen := 22 + 2*p.NumberOfParts + 8*p.NumberOfPoints
	maxLen := minLen + 8 + 4*p.NumberOfPoints
	if cl != minLen && cl != maxLen {
		return nil, fmt.Errorf("invalid content length for PolylineM: %d", cl)
	}

	// Parts
	p.Parts = make([]int32, p.NumberOfParts)
	if err := binary.Read(r, binary.LittleEndian, &p.Parts); err != nil {
		return nil, err
	}

	// Points
	p.Points = make([]Point, p.NumberOfPoints)
	if err := binary.Read(r, binary.LittleEndian, &p.Points); err != nil {
		return nil, err
	}

	// MData
	if cl == minLen {
		p.empty(p.NumberOfPoints)
	} else if err := p.read(r, p.NumberOfPoints); err != nil {
		return nil, err
	}

	return &p, nil
}

package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// MultiPointM represents a set of PointMs. Note that the M data is optional
// and missing values are specified with NaN.
type MultiPointM struct {
	BBox           BBox
	NumberOfPoints int32
	Points         []Point
	MData
}

// Type ...
func (m *MultiPointM) Type() ShapeType {
	return TypeMultiPointM
}

func readMultiPointM(r io.Reader, cl int32) (*MultiPointM, error) {
	var mp MultiPointM

	// BBox
	if err := binary.Read(r, binary.LittleEndian, &mp.BBox); err != nil {
		return nil, err
	}

	// NumberOfPoints
	if err := binary.Read(r, binary.LittleEndian, &mp.NumberOfPoints); err != nil {
		return nil, err
	}

	minLen := 20 + mp.NumberOfPoints*8
	maxLen := minLen + 8 + mp.NumberOfPoints*4

	if cl != minLen && cl != maxLen {
		return nil, fmt.Errorf("invalid content length for MultiPointM: %d", cl)
	}

	// Points
	mp.Points = make([]Point, mp.NumberOfPoints)
	if err := binary.Read(r, binary.LittleEndian, &mp.Points); err != nil {
		return nil, err
	}

	// MData
	if cl == minLen {
		mp.empty(mp.NumberOfPoints)
	} else if err := mp.read(r, mp.NumberOfPoints); err != nil {
		return nil, err
	}

	return &mp, nil
}

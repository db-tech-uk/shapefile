package shapefile

import (
	"github.com/db-tech-uk/shapefile/dbf"
	"github.com/db-tech-uk/shapefile/shp"
)

// Record is record that combines the record taken from the
// specified shp file and also, optionally, adds metadata
// attributes if a dbf file was specified.
type Record struct {
	Shape shp.Shape
	*dbf.Record
}

// Attrs returns all the attributes for this record as a
// slice of strings.
func (r *Record) Attrs() []string {
	if r.Record == nil {
		return nil
	}
	return r.Record.Attrs()
}

package area

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"image"
	// "github.com/paulmach/go.geo"
)

var (
	// Unfortunately we can't use arrays as constants.
	MYSQL_PREFIX = []byte{0, 0, 0, 0}
	// Selecting polygons as text begins with POLYGON((
	POLYGON_PREFIX = []byte{80, 79, 76, 89, 71, 79, 78, 40, 40}
	POLYGON_SUFFIX = []byte{41, 41}
)

// This represents a zero-rectangle Area in the same vein as image.ZR
var ZR = Area{
	Rectangle: image.ZR,
	Valid:     false,
}

// This creates a new Area struct in the same vein as image.NewRect()
func NewArea(x1, y1, x2, y2 int) Area {
	return Area{
		Rectangle: image.Rect(x1, y1, x2, y2),
		Valid:     true,
	}
}

// This area struct represents a monitored area within a page. It inherits from
// an image.Rectangle with the added benefits of scanning directly from SQL and marshalling
// nicely.
type Area struct {
	image.Rectangle
	Valid bool
}

func (a *Area) String() string {
	return a.ToWKT()
}

// Return the monitored area rectangle as a WKT-encoded geom polygon.
// This uses a shitty sprintf hack to make things work easily, cuz fuck geos.
// @TODO: Implement GEOS polygon WKT creation instead of using a poor mans sprintf
func (a *Area) ToWKT() string {
	if a.Rectangle == image.ZR {
		return ""
	}

	return fmt.Sprintf(
		"POLYGON((%d %d, %d %d, %d %d, %d %d, %d %d))",
		a.Rectangle.Min.X, a.Rectangle.Min.Y,
		a.Rectangle.Min.X, a.Rectangle.Max.Y,
		a.Rectangle.Max.X, a.Rectangle.Max.Y,
		a.Rectangle.Max.X, a.Rectangle.Min.Y,
		a.Rectangle.Min.X, a.Rectangle.Min.Y,
	)
}

//
// JSON methods
//

// This should be marshalled into lowercase min/max objects each containing an X and Y coordinate
func (a Area) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"min": map[string]interface{}{
			"x": a.Rectangle.Min.X,
			"y": a.Rectangle.Min.Y,
		},
		"max": map[string]interface{}{
			"x": a.Rectangle.Max.X,
			"y": a.Rectangle.Max.Y,
		},
	})
}

// This mimics the structure of an image.Rectangle using floats instead of ints.
// This allows us to pass floats to the API as monitored areas and marshal from JSON without
// errors; we convert the floats into ints to save.
type FloatRect struct {
	Min struct {
		X float64
		Y float64
	}
	Max struct {
		X float64
		Y float64
	}
}

// Convert the float rectangle into a standard image.Rectangle
func (fr FloatRect) ToRect() image.Rectangle {
	return image.Rect(int(fr.Min.X), int(fr.Min.Y), int(fr.Max.X), int(fr.Max.Y))
}

// This overrides unmarshalling an image.Rectangle to provide support for floats, which are rounded
// to ints to make a rectangle
func (a *Area) UnmarshalJSON(data []byte) (err error) {
	rect := &FloatRect{}
	if err = json.Unmarshal(data, rect); err != nil {
		return
	}
	a.Rectangle = rect.ToRect()
	if a.Rectangle != image.ZR {
		a.Valid = true
	}
	return
}

//
// Database methods
//

// This implements the database driver.Valuer interface which converts itself
// into a string for inserting into a sequel database.
func (a Area) Value() (driver.Value, error) {
	return a.ToWKT(), nil
}

// Implements the sql.Scan interface to decode MySQL polygon data into an Area automatically
func (a *Area) Scan(value interface{}) (err error) {
	var data []byte

	// If the first four bytes of this are 0000
	switch value.(type) {
	// Same as []byte
	case []uint8:
		data = value.([]byte)
	case nil:
		a.Rectangle = image.ZR
		a.Valid = false
		return
	default:
		return fmt.Errorf("Invalid format: can't convert %T into Area", value)
	}

	// MySQL prefixes opengis wkb definitions with four zero-bytes
	if bytes.Equal(MYSQL_PREFIX, data[0:4]) {
		wkb := data[4:len(data)]
		if a.Rectangle, err = decodeWkb(wkb); err != nil {
			return
		}
		a.Valid = true
		return
	}

	if bytes.Equal(POLYGON_PREFIX, data[0:len(POLYGON_PREFIX)]) {
		a.Rectangle = decodeWkt(data)
		a.Valid = true
	}

	return nil
}

// Decode a GEOS WKB byte series into an image.Rectangle
func decodeWkb(wkb []byte) (r image.Rectangle, err error) {
	/*
		var geom, shell, point *geos.Geometry
		var points int

		// Parse our binary data
		if geom, err = geos.FromWKB(wkb); err != nil {
			return
		}

		// Get the LinearRing from our geos polygon
		if shell, err = geom.Shell(); err != nil {
			return
		}

		// Get the total number of points that are in our polygon
		if points, err = shell.NPoint(); err != nil {
			return
		}

		if points != 5 {
			return image.ZR, fmt.Errorf("Unexpected polygon shape; expected 5 points")
		}

		// Iterate through and make our rectangle.
		// For rectangles we save five points in a GEOS polygon: x1y1, x1y2, x2y2, x2y1, and x1y1 again.
		// We only need indexes 0 and 2 to make our rect.
		//
		// When we implement complex monitored areas - not rectangles - we may have to iterate over
		// the geometries using geom.NGeometry(), getting the geometry i via geom.Geometry(i) and then
		// getting the LinearRing via g.Shell()
		//
		// @see http://stackoverflow.com/questions/24017295 for more info
		for i := 0; i <= 2; i += 2 {
			if point, err = shell.Point(i); err != nil {
				return
			}

			x, _ := point.X()
			y, _ := point.Y()

			switch i {
			case 0:
				r.Min.X = int(x)
				r.Min.Y = int(y)
			case 2:
				r.Max.X = int(x)
				r.Max.Y = int(y)
			}
		}
	*/

	return r, err
}

// Decodes a rectangular polygon from WKT format into a rectangle
// where WKT represents a string such as:
// POLYGON((21 22,21 102,101 102,101 22,21 22))
// This is by far the lazy mans way of doing things.
func decodeWkt(wkt []byte) image.Rectangle {
	var x1, x2, y1, y2 int
	fmt.Sscanf(string(wkt), "POLYGON((%d %d,%d %d,%d %d", &x1, &y1, &x1, &y2, &x2, &y2)
	return image.Rect(x1, y1, x2, y2)
}

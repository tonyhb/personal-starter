package types

type Scan struct {
	UUID      string
	Name      string
	Page      string
	CreatedAt time.Time
	Type      ScanType

	HTMLHash        string
	HTMLSize        int
	ImageHash       string
	ImageSize       int
	ImageMime       string
	ImageDifference float64

	// TODO: save changed areas in postgis format

	Tags []Tag

	IsValid bool // scans can be marked as invalid
}

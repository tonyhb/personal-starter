package types

type Page struct {
	UUID      string
	AccountID int

	Name           string
	URL            string
	ScanFrequency  int64 // in minutes
	MonitoredAreas []area.Area
	IgnoredAreas   []area.Area
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

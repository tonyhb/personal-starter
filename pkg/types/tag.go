package types

import (
	"time"
)

// Tag represents a tag associated with a scan
type Tag struct {
	ID        int
	AccountID int
	Name      int
	CreatedAt time.Time
}

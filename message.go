package mod_experiment

import (
	"fmt"
	"time"
)

// LogLine .
type LogLine struct {
	Date time.Time
	Data string
	ID   int64
}

// String .
func (l LogLine) String() string {
	return fmt.Sprintln(l.Date, l.Data, l.ID)
}

// GetData .
func (l LogLine) GetData() string {
	return l.Data
}

// GetID .
func (l LogLine) GetID() int64 {
	return l.ID
}

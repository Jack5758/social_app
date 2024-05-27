package timeparser

import (
	"fmt"
	"time"
)

func ParseDuration(duration string) (time.Duration, error) {
	d, err := time.ParseDuration(duration)
	if err != nil {
		return 0, fmt.Errorf("failed to parse duration %w", err)
	}
	return d, nil
}

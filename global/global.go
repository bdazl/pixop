package global

import (
	"github.com/faiface/pixel"
	"time"
)

var (
	// Program start time
	StartTime = time.Now()

	// Time in seconds since program start
	Time = 0.0

	// Window bounds
	Bounds pixel.Rect

	// Window mid point
	Midpoint pixel.Vec
)

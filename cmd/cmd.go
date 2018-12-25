package cmd

import (
	"github.com/faiface/pixel/pixelgl"
)

type Scene interface {
	Setup() error
	Draw(*pixelgl.Window)
}

package main

import (
	"log"
	"math"
	"os"
	"time"

	"github.com/hexhacks/pixop/cmd/lissajous"
	"github.com/hexhacks/pixop/global"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"github.com/urfave/cli"

	"golang.org/x/image/colornames"
)

const (
	twoPi = math.Pi * 2.0
)

var (
	drawFunc func(*pixelgl.Window)

	width  uint
	height uint
)

func main() {
	app := cli.NewApp()
	app.Name = "pixop"

	app.Flags = []cli.Flag{
		cli.UintFlag{
			Name:        "width, x",
			Value:       1024,
			Usage:       "hidth of the window",
			Destination: &width,
		},
		cli.UintFlag{
			Name:        "height, y",
			Value:       768,
			Usage:       "height of the window",
			Destination: &height,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "lissajous",
			Usage: "draw a sinusoidal pattern",
			Action: func(c *cli.Context) error {
				drawFunc = lissajous.Draw
				pixelgl.Run(run)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run() {
	global.Bounds = pixel.R(0, 0, float64(width), float64(height))

	cfg := pixelgl.WindowConfig{
		Title:  "Jacob <3 Ulrika",
		Bounds: global.Bounds,
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	for !win.Closed() {
		calcTime()

		win.Clear(colornames.Skyblue)

		drawFunc(win)

		win.Update()
	}
}

func calcTime() {
	start := global.StartTime
	global.Time = time.Now().Sub(start).Seconds()
}

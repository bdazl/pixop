package main

import (
	"log"
	"math"
	"os"
	"time"

	"github.com/hexhacks/pixop/cmd"
	"github.com/hexhacks/pixop/cmd/epicycles"
	"github.com/hexhacks/pixop/cmd/gem"
	"github.com/hexhacks/pixop/cmd/koch"
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
	scene cmd.Scene

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
				scene = lissajous.New()
				pixelgl.Run(run)
				return nil
			},
		},
		{
			Name:  "koch",
			Usage: "draw a koch snowflake",
			Action: func(c *cli.Context) error {
				scene = koch.New()
				pixelgl.Run(run)
				return nil
			},
		},
		{
			Name:  "epicycles",
			Usage: "computes and draws epicycles",
			Action: func(c *cli.Context) error {
				scene = epicycles.New()
				pixelgl.Run(run)
				return nil
			},
		},
		{
			Name:  "gem",
			Usage: "gejm time!",
			Action: func(c *cli.Context) error {
				scene = gem.New()
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
		Title:  "Pixop",
		Bounds: global.Bounds,
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = scene.Setup()
	if err != nil {
		log.Fatal(err)
	}

	for !win.Closed() {
		calcTime()

		win.Clear(colornames.Skyblue)

		scene.Draw(win)

		win.Update()
	}
}

func calcTime() {
	start := global.StartTime
	global.Time = time.Now().Sub(start).Seconds()
}

package main

import (
	"bytes"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure-island/resources"
)

var a = 1

var objectImage *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(resources.Objects_png))

	if err != nil {
		log.Fatal(err)
	}
	objectImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

type EntityType int

type Entity struct {
	Coord
	Type EntityType
}

const (
	Coin EntityType = iota + 1
	Score
	Empty
)

func subObject(typ EntityType, frame int) *ebiten.Image {
	var width, height, offsetX, offsetY, sx, sy int
	switch typ {
	case Coin:
		width, height = 16, 16
		offsetY = 62
		frame = frame % 4
	}
	sx = offsetX + frame*width
	sy = offsetY
	return objectImage.SubImage(image.Rect(sx, sy, sx+width, sy+height)).(*ebiten.Image)
}

func (o Entity) Draw(screen *ebiten.Image) {
	switch o.Type {
	case Score:
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(o.X)*tileSize, float64(o.Y)*tileSize)
	t := time.Now().Nanosecond() / 1000 / 1000 / 100 // 10th of 2nd

	screen.DrawImage(subObject(o.Type, t), op)
}

func (o Entity) Destory() Entity {
	switch o.Type {
	case Coin:
		return Entity{
			Coord: Coord{-1, -1},
			Type:  Score,
		}
	}

	return Entity{
		Coord: Coord{-1, -1},
		Type:  Empty,
	}
}